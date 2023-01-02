package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // It is important
	"log"
)

var (
	db *sql.DB
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/social?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
