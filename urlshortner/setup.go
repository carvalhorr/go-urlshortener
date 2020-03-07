package main

import (
	"github.com/spf13/viper"
	"go-urlshortener/persistence"
)

func createUrlShortener() urlShortener {
	return urlShortener{
		shortUrlDao: &persistence.MysqlShortUrlDAO{DB: createDatabaseConnection()},
	}
}

func createDatabaseConnection() persistence.Database {
	db, _ := persistence.CreateMySQLDatabase(
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetUint32("mysql.port"),
		viper.GetString("mysql.db"),
	)
	return &persistence.MysqlDatabase{Mysql: db}
}
