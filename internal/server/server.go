package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"redirectify/internal/services/database"
	"redirectify/internal/utils"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   *sql.DB
}

func NewServer() *http.Server {
	if utils.AppEnv == "debug" {
		appLevel := new(slog.LevelVar)
		h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: appLevel})
		slog.SetDefault(slog.New(h))
		appLevel.Set(slog.LevelDebug)
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	database.New()

	newServer := &Server{
		port: port,
		db:   database.Db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
