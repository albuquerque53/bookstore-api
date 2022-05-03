package database

import (
	"fmt"
	"os"
)

var (
	username = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	hostname = os.Getenv("MYSQL_HOST")
	dbName   = os.Getenv("MYSQL_DATABASE")
)

const DsnTemplate string = "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"

func BuildDsn() string {
	return fmt.Sprintf(DsnTemplate, username, password, hostname, dbName)
}
