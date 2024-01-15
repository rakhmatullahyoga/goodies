package product

import "mime/multipart"

type ProductSearchQuery struct {
	SKU      string
	Title    string
	Category string
	SortBy   SortColumn
	Sort     SortDirection
	Offset   uint64
	Limit    uint64
}

type ProductImageUpload struct {
	ProductID   uint64
	File        multipart.File
	Extension   string
	Description string
}
