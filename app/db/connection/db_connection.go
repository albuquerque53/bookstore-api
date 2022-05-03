package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConectToDatabase() {
	dsn := BuildDsn()

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("unexpected error on database connection: %v", err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("unexpected error on database ping: %v", err)
	}

	fmt.Println("Connected to Database!")
}
