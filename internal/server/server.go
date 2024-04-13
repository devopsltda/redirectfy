package server

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"redirectfy/internal/models"
	"redirectfy/internal/services/database"
	"redirectfy/internal/services/email"
	"redirectfy/internal/utils"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port                   int
	db                     *sql.DB
	email                  *email.Email
	EmailAutenticacaoModel *models.EmailAutenticacaoModel
	LinkModel              *models.LinkModel
	PlanoDeAssinaturaModel *models.PlanoDeAssinaturaModel
	RedirecionadorModel    *models.RedirecionadorModel
	UsuarioModel           *models.UsuarioModel
	UsuarioKirvanoModel    *models.UsuarioKirvanoModel
}

func NewServer() *http.Server {
	if utils.AppEnv == "debug" {
		appLevel := new(slog.LevelVar)
		h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: appLevel})
		slog.SetDefault(slog.New(h))
		appLevel.Set(slog.LevelDebug)
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	db := database.New()

	newServer := &Server{
		port:                   port,
		db:                     database.New(),
		email:                  email.New(),
		EmailAutenticacaoModel: &models.EmailAutenticacaoModel{DB: db},
		LinkModel:              &models.LinkModel{DB: db},
		PlanoDeAssinaturaModel: &models.PlanoDeAssinaturaModel{DB: db},
		RedirecionadorModel:    &models.RedirecionadorModel{DB: db},
		UsuarioModel:           &models.UsuarioModel{DB: db},
		UsuarioKirvanoModel:    &models.UsuarioKirvanoModel{DB: db},
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
