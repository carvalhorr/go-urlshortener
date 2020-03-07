package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-urlshortener/model"
	"time"
)

type MysqlDatabase struct {
	Mysql *sql.DB
}

func CreateMySQLDatabase(user string, password string, server string, port uint32, dbName string) (*sql.DB, error) {
	dbConnectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, dbName)
	db, err := sql.Open("mysql", dbConnectionStr)
	db.SetConnMaxLifetime(10 * time.Minute)
	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)
	return db, err
}

func (db *MysqlDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Mysql.Query(query, args...)
}

type MysqlShortUrlDAO struct {
	DB Database
}

func (dao *MysqlShortUrlDAO) Add(url model.ShortURL) error {
	sql := fmt.Sprintf("INSERT INTO short_url(short_key, long_url) VALUES('%s', '%s')",
		url.ShortURLCode,
		url.LongURL)

	insert, err := dao.DB.Query(sql)
	if err != nil {
		// TODO Log underlying reason
		return fmt.Errorf("failed to persist short URL")
	}

	defer insert.Close()
	return nil
}

func (dao *MysqlShortUrlDAO) Get(shortURLCode string) (*model.ShortURL, error) {
	// TODO Implement
	return nil, fmt.Errorf("not implemented")
}
