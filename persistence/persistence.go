package persistence

import (
	"database/sql"
	"github.com/carvalhorr/go-urlshortener/model"
)

type ShortUrlDAO interface {
	Add(url model.ShortURL) error
	Get(shortURLCode string) (*model.ShortURL, error)
}

type Database interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
