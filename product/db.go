package product

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB

	ErrEmptyDB = errors.New("cannot assign with empty db")
)

func SetDB(newdb *sqlx.DB) (err error) {
	if newdb == nil {
		err = ErrEmptyDB
		return
	}

	db = newdb
	return
}
