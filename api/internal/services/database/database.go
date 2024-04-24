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

	TempDir     string
	dbConnector *libsql.Connector
)

func New() *sql.DB {
	var err error

	// Caso a variaǘel APP_ENV seja de "test" ou "debugtest", o banco de
	// dados utilizado é inicializado em memória, fazendo com que ele não
	// seja persistente.
	if utils.AppEnv == "test" || utils.AppEnv == "debugtest" {
		db, err := sql.Open("sqlite", "file::memory:?cache=shared")

		if err != nil {
			slog.Error("BancoDeDadosTeste", slog.Any("error", err))
			os.Exit(1)
		}

		err = seed(db, "../internal/services/database/source/")

		if err != nil {
			slog.Error("BancoDeDadosTeste", slog.Any("error", err))
			os.Exit(1)
		}

		return db
	}

	// É necessário criar um diretório temporário para as réplicas de banco
	// de dados criadas pelo Turso
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

	err = db.Ping()

	if err != nil {
		slog.Error("BancoDeDados", slog.Any("error", err))
		os.Exit(1)
	}

	// Essas linhas apenas devem ser descomentadas na primeira execução, logo
	// depois de criar o banco de dados. Ela é responsável por rodar os scripts
	// de criação de tabelas, índices e triggers, além de popular o banco com
	// os planos de assinatura.
	// Isso é necessário porque o Turso não consegue ler os triggers diretamente
	// pela CLI, como descrito na issue:
	//
	// https://github.com/tursodatabase/libsql-shell-go/issues/124
	//
	// err = seed(db, "../internal/services/database/source/")
	// 
	// if err != nil {
	// 	slog.Error("BancoDeDados", slog.Any("error", err))
	// 	os.Exit(1)
	// }

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
