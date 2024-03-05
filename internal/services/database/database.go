package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Db    *sql.DB
	dbUrl = os.Getenv("DB_URL")
)

func New() {
	var err error

	switch os.Getenv("APP_ENV") {
	case "test":
		Db, err = sql.Open("sqlite3", "file::memory:?cache=shared")
		seed(os.Getenv("DB_SOURCE_PATH"))
	default:
		Db, err = sql.Open("sqlite3", dbUrl)
	}

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
}

func seed(dbSeedPath string) error {
	query, err := os.ReadFile(dbSeedPath)

	if err != nil {
		return err
	}

	_, err = Db.Exec(string(query))

	return err
}
