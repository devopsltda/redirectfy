package database

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"
	"redirectify/internal/utils"
	"time"

	"github.com/tursodatabase/go-libsql"

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

	if utils.AppEnv == "test" {
		Db, err = sql.Open("sqlite", "file::memory:?cache=shared")
		seed(os.Getenv("DB_SOURCE_PATH"))

		if err != nil {
			slog.Error("BancoDeDados", slog.Any("error", err))
			os.Exit(1)
		}

		return
	}

	TempDir, err := os.MkdirTemp("", "libsql-*")

	if err != nil {
		slog.Error("BancoDeDados", slog.Any("error", err))
		os.Exit(1)
	}

	dbPath := filepath.Join(TempDir, dbName)
	DbConnector, err = libsql.NewEmbeddedReplicaConnector(dbPath, dbUrl, libsql.WithAuthToken(dbToken), libsql.WithSyncInterval(15*time.Minute))

	if err != nil {
		slog.Error("BancoDeDados", slog.Any("error", err))
		os.Exit(1)
	}

	Db = sql.OpenDB(DbConnector)

	if err != nil {
		slog.Error("BancoDeDados", slog.Any("error", err))
		os.Exit(1)
	}
}

func seed(dbSourcePath string) error {
	queryDDL, err := os.ReadFile(dbSourcePath + "ddl.sql")

	if err != nil {
		return err
	}

	_, err = Db.Exec(string(queryDDL))

	if err != nil {
		return err
	}

	queryTriggers, err := os.ReadFile(dbSourcePath + "triggers.sql")

	if err != nil {
		return err
	}

	_, err = Db.Exec(string(queryTriggers))

	return err
}
