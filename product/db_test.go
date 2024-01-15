package product_test

import (
	"goodies/product"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestSetDB(t *testing.T) {
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	type args struct {
		newdb *sqlx.DB
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "empty db",
			args:    args{},
			wantErr: product.ErrEmptyDB,
		},
		{
			name: "with mock db",
			args: args{
				newdb: sqlxDB,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := product.SetDB(tt.args.newdb); err != tt.wantErr {
				t.Errorf("SetDB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
