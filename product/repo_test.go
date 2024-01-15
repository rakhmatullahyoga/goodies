package product

import (
	"context"
	"mime/multipart"
	"reflect"
	"testing"
)

func Test_getProducts(t *testing.T) {
	type args struct {
		ctx         context.Context
		searchQuery ProductSearchQuery
	}
	tests := []struct {
		name         string
		args         args
		wantProducts []ProductDetail
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProducts, err := getProducts(tt.args.ctx, tt.args.searchQuery)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProducts, tt.wantProducts) {
				t.Errorf("getProducts() = %v, want %v", gotProducts, tt.wantProducts)
			}
		})
	}
}

func Test_saveProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		product *Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveProduct(tt.args.ctx, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("saveProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint64
	}
	tests := []struct {
		name        string
		args        args
		wantProduct Product
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProduct, err := getProduct(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("getProduct() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func Test_saveProductImage(t *testing.T) {
	type args struct {
		ctx          context.Context
		productImage *ProductImage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveProductImage(tt.args.ctx, tt.args.productImage); (err != nil) != tt.wantErr {
				t.Errorf("saveProductImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_saveProductImageFile(t *testing.T) {
	type args struct {
		file      multipart.File
		extension string
	}
	tests := []struct {
		name         string
		args         args
		wantFilename string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFilename, err := saveProductImageFile(tt.args.file, tt.args.extension)
			if (err != nil) != tt.wantErr {
				t.Errorf("saveProductImageFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFilename != tt.wantFilename {
				t.Errorf("saveProductImageFile() = %v, want %v", gotFilename, tt.wantFilename)
			}
		})
	}
}

func Test_saveProductReview(t *testing.T) {
	type args struct {
		ctx           context.Context
		productReview *ProductReview
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveProductReview(tt.args.ctx, tt.args.productReview); (err != nil) != tt.wantErr {
				t.Errorf("saveProductReview() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getProductImages(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID uint64
	}
	tests := []struct {
		name       string
		args       args
		wantImages []ProductImage
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotImages, err := getProductImages(tt.args.ctx, tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProductImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotImages, tt.wantImages) {
				t.Errorf("getProductImages() = %v, want %v", gotImages, tt.wantImages)
			}
		})
	}
}

func Test_getProductReviews(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID uint64
	}
	tests := []struct {
		name        string
		args        args
		wantReviews []ProductReview
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReviews, err := getProductReviews(tt.args.ctx, tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProductReviews() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReviews, tt.wantReviews) {
				t.Errorf("getProductReviews() = %v, want %v", gotReviews, tt.wantReviews)
			}
		})
	}
}
