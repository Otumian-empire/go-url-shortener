package repository

import (
	"database/sql"
	"fmt"

	"github.com/otumian-empire/go-url-shortener/util"
)

// the root of the tables
type Store struct {
	*UrlStore
}

// database connection
func NewStore(driverName, dataSourceName string) (*Store, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if util.IsNotNil(err) {
		return nil, fmt.Errorf("%v: %w", DATABASE_OPENING_ERROR, err)
	}

	if err := db.Ping(); util.IsNotNil(err) {
		return nil, fmt.Errorf("%v: %w", DATABASE_CONNECTING_ERROR, err)
	}

	return &Store{UrlStore: &UrlStore{DB: db}}, nil
}
