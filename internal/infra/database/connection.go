package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConectToDatabase() *sql.DB {
	dsn := BuildDsn()

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Panicf("unexpected error on database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Panicf("unexpected during database ping: %v", err)
	}

	return db
}
