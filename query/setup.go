package main

import "go-urlshortener/persistence"

type shortUrlQuery struct {
	shortUrlDao persistence.ShortUrlDAO
}

func createShortUrlQuery() shortUrlQuery {
	return shortUrlQuery{
		shortUrlDao: &persistence.MysqlShortUrlDAO{DB: persistence.CreateDatabaseConnection()},
	}
}
