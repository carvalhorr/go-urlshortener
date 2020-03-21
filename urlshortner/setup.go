package main

import (
	"github.com/carvalhorr/go-urlshortener/persistence"
)

type urlShortener struct {
	shortUrlDao persistence.ShortUrlDAO
}

func createUrlShortener() urlShortener {
	return urlShortener{
		shortUrlDao: &persistence.MysqlShortUrlDAO{DB: persistence.CreateDatabaseConnection()},
	}
}
