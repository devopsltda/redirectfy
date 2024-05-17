package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"redirectfy/internal/utils"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var (
	dbUrl        = os.Getenv("DB_URL")
	dbToken      = os.Getenv("DB_TOKEN")
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

	db, err := sql.Open("libsql", fmt.Sprintf("%s?authToken=%s", dbUrl, dbToken))

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
	// err = seed(db, "./internal/services/database/source/")
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
