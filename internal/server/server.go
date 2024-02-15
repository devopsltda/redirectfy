package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/TheDevOpsCorp/redirectify/internal/database"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int
	db   *sql.DB
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	database.Db = database.New()

	NewServer := &Server{
		port: port,
		db:   database.Db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func NewTestServer() *Server {
	return &Server{
		db: database.NewTest(),
	}
}
