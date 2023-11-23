package product

import (
	"context"
	"math"
	"os"
)

func getProductDetails(ctx context.Context, id uint64) (productDetail ProductDetail, err error) {
	product, err := getProduct(ctx, id)
	if err != nil {
		return
	}

	images, err := getProductImages(ctx, id)
	if err != nil {
		return
	}
	setImageUrl(&images)

	reviews, err := getProductReviews(ctx, id)
	if err != nil {
		return
	}
	rating := calculateRatings(reviews)

	productDetail.Product = product
	productDetail.Rating = rating
	if len(images) > 0 {
		productDetail.Images = &images
	}
	if len(reviews) > 0 {
		productDetail.Reviews = &reviews
	}
	return
}

func setImageUrl(images *[]ProductImage) {
	imgPath := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT") + "/static/img/"
	for i := range *images {
		(*images)[i].ImageUrl = imgPath + (*images)[i].ImageUrl
	}
}

func calculateRatings(reviews []ProductReview) (rating float64) {
	var ratingTotal, ratingCount uint
	for _, review := range reviews {
		ratingTotal += uint(review.Rating)
		ratingCount += 1
	}
	if ratingCount > 0 {
		rating = math.Round((float64(ratingTotal)/float64(ratingCount))*10) / 10
	}
	return
}

func submitProduct(ctx context.Context, product *Product) (err error) {
	err = saveProduct(ctx, product)
	return
}

func submitProductImage(ctx context.Context, imgData ProductImageUpload) (err error) {
	_, err = getProduct(ctx, imgData.ProductID)
	if err != nil {
		return
	}

	filename, err := saveProductImageFile(imgData.File, imgData.Extension)
	if err != nil {
		return
	}

	productImage := ProductImage{
		ProductID:        imgData.ProductID,
		ImageDescription: imgData.Description,
		ImageUrl:         filename,
	}
	err = saveProductImage(ctx, &productImage)
	return
}

func submitProductReview(ctx context.Context, productReview ProductReview) (err error) {
	_, err = getProduct(ctx, productReview.ProductID)
	if err != nil {
		return
	}

	err = saveProductReview(ctx, &productReview)
	return
}

func listProducts(ctx context.Context, searchQuery ProductSearchQuery) (products []ProductDetail, err error) {
	products, err = getProducts(ctx, searchQuery)
	return
}
