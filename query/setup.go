package main

import "github.com/carvalhorr/go-urlshortener/persistence"

type shortUrlQuery struct {
	shortUrlDao persistence.ShortUrlDAO
}

func createShortUrlQuery() shortUrlQuery {
	return shortUrlQuery{
		shortUrlDao: &persistence.MysqlShortUrlDAO{DB: persistence.CreateDatabaseConnection()},
	}
}
