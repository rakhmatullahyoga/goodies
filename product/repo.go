package product

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	getProductImagesQuery    = "SELECT id, product_id, image_url, image_description FROM product_images WHERE product_id = ?"
	getProductQuery          = "SELECT id, title, sku, description, category, weight, price, created_at, updated_at FROM products WHERE id = ?"
	getProductReviewsQuery   = "SELECT id, product_id, rating, comment FROM product_reviews WHERE product_id = ?"
	insertProductQuery       = "INSERT INTO products (title, sku, description, category, weight, price) VALUES (?, ?, ?, ?, ?, ?)"
	insertProductImageQuery  = "INSERT INTO product_images (product_id, image_url, image_description) VALUES (?, ?, ?)"
	insertProductReviewQuery = "INSERT INTO product_reviews (product_id, rating, comment) VALUES (?, ?, ?)"
	updateProductQuery       = "UPDATE products SET title = ?, sku = ?, description = ?, category = ?, weight = ?, price = ?, updated_at = ? WHERE id = ?"

	duplicateDbConstraint = "Duplicate entry"
)

var (
	ErrNoRows       = errors.New("no rows affected")
	ErrNotFound     = errors.New("record not found")
	ErrSkuDuplicate = errors.New("sku must be unique")
)

func imageDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b) + "/../static/img/"
}

func getProducts(ctx context.Context, searchQuery ProductSearchQuery) (products []ProductDetail, err error) {
	var whereClause string
	var whereClauses []string
	var whereParams []interface{}

	if searchQuery.Title != "" {
		whereClauses = append(whereClauses, `title = ?`)
		whereParams = append(whereParams, searchQuery.Title)
	}
	if searchQuery.SKU != "" {
		whereClauses = append(whereClauses, `sku = ?`)
		whereParams = append(whereParams, searchQuery.SKU)
	}
	if searchQuery.Category != "" {
		whereClauses = append(whereClauses, `category = ?`)
		whereParams = append(whereParams, searchQuery.Category)
	}

	if len(whereClauses) > 0 {
		whereClause = `WHERE ` + strings.Join(whereClauses, " AND ")
	}

	query := `SELECT p.*, COALESCE(pr.rating,0) AS rating FROM products p LEFT JOIN ` +
		`(SELECT product_id, COALESCE(ROUND(SUM(rating)/COUNT(rating),1),0) AS rating FROM product_reviews GROUP BY product_id) AS pr ` +
		`ON pr.product_id = p.id ` + whereClause +
		` ORDER BY ` + string(searchQuery.SortBy) + ` ` + string(searchQuery.Sort) + ` LIMIT ? OFFSET ?`
	whereParams = append(whereParams, searchQuery.Limit, searchQuery.Offset)

	err = db.SelectContext(ctx, &products, query, whereParams...)
	return
}

func saveProduct(ctx context.Context, product *Product) (err error) {
	if product.ID == 0 {
		// insert new data
		res, err := db.ExecContext(ctx, insertProductQuery, product.Title, product.Sku, product.Description, product.Category, product.Weight, product.Price)
		if err != nil {
			if strings.Contains(err.Error(), duplicateDbConstraint) {
				err = ErrSkuDuplicate
			}

			return err
		}

		id, _ := res.LastInsertId()
		product.ID = uint64(id)
	} else {
		// do update
		res, err := db.ExecContext(ctx, updateProductQuery, product.Title, product.Sku, product.Description, product.Category, product.Weight, product.Price, time.Now(), product.ID)
		if err != nil {
			if strings.Contains(err.Error(), duplicateDbConstraint) {
				err = ErrSkuDuplicate
			}

			return err
		}

		count, _ := res.RowsAffected()
		if count == 0 {
			return ErrNoRows
		}
	}
	return
}

func getProduct(ctx context.Context, id uint64) (product Product, err error) {
	err = db.GetContext(ctx, &product, getProductQuery, id)
	if err != nil {
		if err == sql.ErrNoRows {
			err = ErrNotFound
		}
	}

	return
}

func saveProductImage(ctx context.Context, productImage *ProductImage) (err error) {
	res, err := db.ExecContext(ctx, insertProductImageQuery, productImage.ProductID, productImage.ImageUrl, productImage.ImageDescription)
	if err != nil {
		return
	}

	id, _ := res.LastInsertId()
	productImage.ID = uint64(id)
	return
}

func saveProductImageFile(file multipart.File, extension string) (filename string, err error) {
	tempFile, err := os.CreateTemp(imageDirectory(), "upload-*."+extension)
	if err != nil {
		return
	}
	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return
	}

	if _, err = tempFile.Write(fileBytes); err != nil {
		return
	}

	filePath := tempFile.Name()
	dirs := strings.Split(filePath, "/")
	filename = dirs[len(dirs)-1]
	return
}

func saveProductReview(ctx context.Context, productReview *ProductReview) (err error) {
	res, err := db.ExecContext(ctx, insertProductReviewQuery, productReview.ProductID, productReview.Rating, productReview.Comment)
	if err != nil {
		return
	}

	id, _ := res.LastInsertId()
	productReview.ID = uint64(id)
	return
}

func getProductImages(ctx context.Context, productID uint64) (images []ProductImage, err error) {
	err = db.SelectContext(ctx, &images, getProductImagesQuery, productID)
	return
}

func getProductReviews(ctx context.Context, productID uint64) (reviews []ProductReview, err error) {
	err = db.SelectContext(ctx, &reviews, getProductReviewsQuery, productID)
	return
}
