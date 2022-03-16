package util

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "admin2:12345@tcp(localhost:3306)/golang_example?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func Close(DB *sql.DB) {
	err := DB.Close()
	if err != nil {
		panic(err)
	}
}
