package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Db *sql.DB
	dburl = os.Getenv("DB_URL")
)

func New() *sql.DB {
	db, err := sql.Open("sqlite3", dburl)

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	return db
}

func NewTest() *sql.DB {
	db, err := sql.Open("sqlite3", "../storage/test.db")

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	return db
}
