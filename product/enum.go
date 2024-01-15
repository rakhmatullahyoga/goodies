package product

type SortColumn string

type SortDirection string

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
