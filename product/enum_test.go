package product

import (
	"testing"
)

func Test_convertSortDirection(t *testing.T) {
	type args struct {
		sortDirection string
	}
	tests := []struct {
		name   string
		args   args
		wantSd SortDirection
	}{
		{
			name: "with ascending input",
			args: args{
				sortDirection: "asc",
			},
			wantSd: sortAscending,
		},
		{
			name: "with descending input",
			args: args{
				sortDirection: "desc",
			},
			wantSd: sortDescending,
		},
		{
			name: "with invalid input",
			args: args{
				sortDirection: "invalid",
			},
			wantSd: sortDescending,
		},
		{
			name: "with empty input",
			args: args{
				sortDirection: "",
			},
			wantSd: sortDescending,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSd := convertSortDirection(tt.args.sortDirection); gotSd != tt.wantSd {
				t.Errorf("convertSortDirection() = %v, want %v", gotSd, tt.wantSd)
			}
		})
	}
}

func Test_convertSortColumn(t *testing.T) {
	type args struct {
		sortBy string
	}
	tests := []struct {
		name   string
		args   args
		wantSc SortColumn
	}{
		{
			name: "with created_at input",
			args: args{
				sortBy: "created_at",
			},
			wantSc: sortByDate,
		},
		{
			name: "with rating input",
			args: args{
				sortBy: "rating",
			},
			wantSc: sortByRating,
		},
		{
			name: "with invalid input",
			args: args{
				sortBy: "invalid",
			},
			wantSc: sortByDate,
		},
		{
			name: "with empty input",
			args: args{
				sortBy: "empty",
			},
			wantSc: sortByDate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSc := convertSortColumn(tt.args.sortBy); gotSc != tt.wantSc {
				t.Errorf("convertSortColumn() = %v, want %v", gotSc, tt.wantSc)
			}
		})
	}
}
