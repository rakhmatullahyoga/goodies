package product

import "errors"

var (
	ErrTitleMissing     = errors.New("product's title required")
	ErrSkuMissing       = errors.New("product's sku required")
	ErrPriceMissing     = errors.New("product's price required")
	ErrRatingMissing    = errors.New("product's rating required")
	ErrRatingInvalid    = errors.New("product's rating must be in 1-5")
	ErrProductIdMissing = errors.New("product's id required")
)

func validateProduct(product Product) (err error) {
	switch {
	case product.Title == "":
		err = ErrTitleMissing
	case product.Sku == "":
		err = ErrSkuMissing
	case product.Price == 0:
		err = ErrPriceMissing
	}
	return
}

func validateProductReview(productReview ProductReview) (err error) {
	switch {
	case productReview.Rating == 0:
		err = ErrRatingMissing
	case productReview.Rating > 5:
		err = ErrRatingInvalid
	case productReview.ProductID == 0:
		err = ErrProductIdMissing
	}
	return
}
