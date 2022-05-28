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
		log.Fatalf("unexpected error on database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("unexpected error on database ping: %v", err)
	}

	fmt.Println("Connected to Database!")

	return db
}
