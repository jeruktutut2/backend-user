package util

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeruktutut2/backend-user/configuration"
)

func NewDatabaseConnection(databaseConfiguration configuration.Database) *sql.DB {
	db, err := sql.Open("mysql", databaseConfiguration.Username+":"+databaseConfiguration.Password+"@tcp("+databaseConfiguration.Host+":"+strconv.Itoa(databaseConfiguration.Port)+")/"+databaseConfiguration.Database+"?parseTime=true")
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
