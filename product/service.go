package product

import (
	"context"
	"math"
	"os"
)

var (
	getProductsRepo          = getProducts
	getProductRepo           = getProduct
	getProductImagesRepo     = getProductImages
	getProductReviewsRepo    = getProductReviews
	saveProductRepo          = saveProduct
	saveProductImageFileRepo = saveProductImageFile
	saveProductImageRepo     = saveProductImage
	saveProductReviewRepo    = saveProductReview
)

func getProductDetails(ctx context.Context, id uint64) (productDetail ProductDetail, err error) {
	product, err := getProductRepo(ctx, id)
	if err != nil {
		return
	}

	images, err := getProductImagesRepo(ctx, id)
	if err != nil {
		return
	}
	setImageUrl(&images)

	reviews, err := getProductReviewsRepo(ctx, id)
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
	err = saveProductRepo(ctx, product)
	return
}

func submitProductImage(ctx context.Context, imgData ProductImageUpload) (err error) {
	_, err = getProductRepo(ctx, imgData.ProductID)
	if err != nil {
		return
	}

	filename, err := saveProductImageFileRepo(imgData.File, imgData.Extension)
	if err != nil {
		return
	}

	productImage := ProductImage{
		ProductID:        imgData.ProductID,
		ImageDescription: imgData.Description,
		ImageUrl:         filename,
	}
	err = saveProductImageRepo(ctx, &productImage)
	return
}

func submitProductReview(ctx context.Context, productReview ProductReview) (err error) {
	_, err = getProductRepo(ctx, productReview.ProductID)
	if err != nil {
		return
	}

	err = saveProductReviewRepo(ctx, &productReview)
	return
}

func listProducts(ctx context.Context, searchQuery ProductSearchQuery) (products []ProductDetail, err error) {
	products, err = getProductsRepo(ctx, searchQuery)
	return
}
