package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/libsql/go-libsql"

	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"
)

var (
	dbUrl   = os.Getenv("DB_URL")
	dbName  = os.Getenv("DB_NAME")
	dbToken = os.Getenv("DB_TOKEN")

	Db          *sql.DB
	TempDir     string
	DbConnector *libsql.Connector
)

func New() {
	var err error

	switch os.Getenv("APP_ENV") {
	case "test":
		Db, err = sql.Open("sqlite", "file::memory:?cache=shared")
		seed(os.Getenv("DB_SOURCE_PATH"))
	default:
		TempDir, err := os.MkdirTemp("", "libsql-*")

		if err != nil {
			log.Fatal("Erro ao criar um diretório temporário:", err)
		}

		dbPath := filepath.Join(TempDir, dbName)
		DbConnector, err = libsql.NewEmbeddedReplicaConnector(dbPath, dbUrl, libsql.WithAuthToken(dbToken), libsql.WithAutoSync(15*time.Minute))

		if err != nil {
			log.Fatal("Erro ao criar um conector:", err)
		}

		Db = sql.OpenDB(DbConnector)
	}

	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
}

func seed(dbSourcePath string) error {
	queryDDL, err := os.ReadFile(dbSourcePath+"ddl.sql")

	if err != nil {
		return err
	}

	_, err = Db.Exec(string(queryDDL))

	if err != nil {
		return err
	}

	queryTriggers, err := os.ReadFile(dbSourcePath+"triggers.sql")

	if err != nil {
		return err
	}

	_, err = Db.Exec(string(queryTriggers))

	return err
}
