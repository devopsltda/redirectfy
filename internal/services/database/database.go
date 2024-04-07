package database

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"
	"redirectfy/internal/utils"
	"time"

	"github.com/tursodatabase/go-libsql"

	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"
)

var (
	dbUrl        = os.Getenv("DB_URL")
	dbName       = os.Getenv("DB_NAME")
	dbToken      = os.Getenv("DB_TOKEN")
	dbSourcePath = os.Getenv("DB_SOURCE_PATH")

	TempDir     string
	dbConnector *libsql.Connector
)

func New() *sql.DB {
	var err error

	if utils.AppEnv == "test" {
		db, err := sql.Open("sqlite", "file::memory:?cache=shared")

		if err != nil {
			slog.Error("BancoDeDadosTeste", slog.Any("error", err))
			os.Exit(1)
		}

		err = seed(db, dbSourcePath)

		if err != nil {
			slog.Error("BancoDeDadosTeste", slog.Any("error", err))
			os.Exit(1)
		}

		return db
	}

	TempDir, err := os.MkdirTemp("", "libsql-*")

	if err != nil {
		slog.Error("BancoDeDados", slog.Any("error", err))
		os.Exit(1)
	}

	dbPath := filepath.Join(TempDir, dbName)
	dbConnector, err = libsql.NewEmbeddedReplicaConnector(dbPath, dbUrl, libsql.WithAuthToken(dbToken), libsql.WithSyncInterval(15*time.Minute))

	if err != nil {
		slog.Error("BancoDeDados", slog.Any("error", err))
		os.Exit(1)
	}

	db := sql.OpenDB(dbConnector)

	if err != nil {
		slog.Error("BancoDeDados", slog.Any("error", err))
		os.Exit(1)
	}

	err = seed(db, dbSourcePath[1:])

	if err != nil {
		slog.Error("BancoDeDados", slog.Any("error", err))
		os.Exit(1)
	}

	return db
}

func seed(db *sql.DB, dbSourcePath string) error {
	queryDDL, err := os.ReadFile(dbSourcePath + "ddl.sql")

	if err != nil {
		return err
	}

	_, err = db.Exec(string(queryDDL))

	if err != nil {
		return err
	}

	queryTriggers, err := os.ReadFile(dbSourcePath + "triggers.sql")

	if err != nil {
		return err
	}

	_, err = db.Exec(string(queryTriggers))

	return err
}
