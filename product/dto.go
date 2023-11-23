package product

import "mime/multipart"

type SortColumn string

type SortDirection string

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

const (
	sortByDate   SortColumn = "created_at"
	sortByRating SortColumn = "rating"
)

const (
	sortAscending  SortDirection = "asc"
	sortDescending SortDirection = "desc"
)

func convertSortDirection(sortDirection string) (sd SortDirection) {
	switch sortDirection {
	case string(sortAscending):
		sd = sortAscending
	case string(sortDescending):
		sd = sortDescending
	default:
		sd = sortDescending
	}
	return
}

func convertSortColumn(sortBy string) (sc SortColumn) {
	switch sortBy {
	case string(sortByDate):
		sc = sortByDate
	case string(sortByRating):
		sc = sortByRating
	default:
		sc = sortByDate
	}
	return
}
