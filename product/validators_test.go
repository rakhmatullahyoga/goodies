package product

import (
	"testing"
)

func Test_validateProduct(t *testing.T) {
	type args struct {
		product Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "missing product's title",
			args: args{
				product: Product{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateProduct(tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("validateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateProductReview(t *testing.T) {
	type args struct {
		productReview ProductReview
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
			if err := validateProductReview(tt.args.productReview); (err != nil) != tt.wantErr {
				t.Errorf("validateProductReview() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
