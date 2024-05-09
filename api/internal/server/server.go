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

// Essa estrutura do servidor armazena informações essenciais
// para seu uso, como a porta do servidor, a pool de conexões
// com o banco de dados e os modelos para cada tabela específica.
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
	// Mostra o nível de log DEBUG caso a variável de ambiente APP_ENV
	// seja "debug" (banco em nuvem) ou "debugtest" (banco em memória)
	if utils.AppEnv == "debug" || utils.AppEnv == "debugtest" {
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

	// Inicialiação de um servidor HTTP baseado no servidor criado
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
