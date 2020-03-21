package persistence

import (
	"database/sql"
	"fmt"
	"github.com/carvalhorr/go-urlshortener/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"time"
)

type MysqlDatabase struct {
	Mysql *sql.DB
}

func CreateDatabaseConnection() Database {
	db, _ := CreateMySQLDatabase(
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetUint32("mysql.port"),
		viper.GetString("mysql.db"),
	)
	return &MysqlDatabase{Mysql: db}
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
	sql := fmt.Sprintf("SELECT id, long_url FROM short_url WHERE short_key = '%s'", shortURLCode)
	result, err := dao.DB.Query(sql)
	if err != nil {
		// TODO Log error
		return nil, fmt.Errorf("Could not retrieve short URL.")
	}
	defer result.Close()
	var longUrl string
	var id int64
	if result.Next() {
		err := result.Scan(&id, &longUrl)
		fmt.Println(err)
	}
	return &model.ShortURL{
		ID:           id,
		ShortURLCode: shortURLCode,
		LongURL:      longUrl,
	}, nil
}
