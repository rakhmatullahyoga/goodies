package product

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func Test_getProductDetails(t *testing.T) {
	type args struct {
		ctx context.Context
		id  uint64
	}
	tests := []struct {
		name              string
		args              args
		wantProductDetail ProductDetail
		wantErr           bool
		init              func()
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotProductDetail, err := getProductDetails(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProductDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProductDetail, tt.wantProductDetail) {
				t.Errorf("getProductDetails() = %v, want %v", gotProductDetail, tt.wantProductDetail)
			}
		})
	}
}

func Test_submitProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		product *Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		init    func()
	}{
		{
			name: "submitProduct() success",
			args: args{
				ctx:     context.Background(),
				product: &Product{},
			},
			wantErr: false,
			init: func() {
				saveProductRepo = func(ctx context.Context, product *Product) (err error) {
					return
				}
			},
		},
		{
			name: "submitProduct() error",
			args: args{
				ctx:     context.Background(),
				product: &Product{},
			},
			wantErr: true,
			init: func() {
				saveProductRepo = func(ctx context.Context, product *Product) (err error) {
					err = errors.New("unexpected error")
					return
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.init()
			if err := submitProduct(tt.args.ctx, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("submitProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
