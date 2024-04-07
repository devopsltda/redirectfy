package server

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"redirectfy/internal/auth"

	_ "redirectfy/docs"
)

// @title API do Redirectify
//
// @version 1.0.0
//
// @description API para interagir com o Redirectify
//
// @contact.name Equipe da DevOps (Pablo Eduardo, Guilherme Bernardo e Eduardo Henrique)
//
// @contact.email comercialdevops@gmail.com
func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/docs")
		},
	}))

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(auth.ChaveDeAcesso),
		TokenLookup: "cookie:access-token",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.Claims)
		},
		Skipper: auth.PathWithNoAuthRequired,
	}))

	// API - Documentação Swagger
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// API - Usuário
	e.GET("/usuarios", s.UsuarioReadAll)
	e.GET("/usuarios/historico", s.HistoricoUsuarioReadAll)
	e.GET("/usuarios/:nome_de_usuario", s.UsuarioReadByNomeDeUsuario)
	e.POST("/usuarios_temporarios", s.UsuarioTemporarioCreate)
	e.POST("/usuarios_temporarios/historico", s.HistoricoUsuarioTemporarioReadAll)
	e.PATCH("/usuarios/:nome_de_usuario", s.UsuarioUpdate)
	e.DELETE("/usuarios/:nome_de_usuario", s.UsuarioRemove)

	// API - Autenticação
	e.POST("/usuarios/login", s.UsuarioLogin)
	e.PATCH("/usuarios/troca_de_senha/:valor", s.UsuarioTrocaDeSenha)
	e.POST("/usuarios/:nome_de_usuario/troca_de_senha", s.UsuarioTrocaDeSenhaExigir)
	e.PATCH("/autenticacao/:valor", s.UsuarioAutenticado)

	// API - Plano de Assinatura
	e.GET("/planos_de_assinatura", s.PlanoDeAssinaturaReadAll)
	e.GET("/planos_de_assinatura/historico", s.HistoricoPlanoDeAssinaturaReadAll)
	e.GET("/planos_de_assinatura/:nome", s.PlanoDeAssinaturaReadByNome)
	e.POST("/planos_de_assinatura", s.PlanoDeAssinaturaCreate)
	e.PATCH("/planos_de_assinatura/:nome", s.PlanoDeAssinaturaUpdate)
	e.DELETE("/planos_de_assinatura/:nome", s.PlanoDeAssinaturaRemove)

	// API - Redirecionador
	e.GET("/redirecionadores", s.RedirecionadorReadAll)
	e.GET("/redirecionadores/historico", s.HistoricoRedirecionadorReadAll)
	e.GET("/redirecionadores/:codigo_hash", s.RedirecionadorReadByCodigoHash)
	e.POST("/redirecionadores", s.RedirecionadorCreate)
	e.PATCH("/redirecionadores/:codigo_hash/rehash", s.RedirecionadorRehash)
	e.PATCH("/redirecionadores/:codigo_hash", s.RedirecionadorUpdate)
	e.DELETE("/redirecionadores/:codigo_hash", s.RedirecionadorRemove)

	// API - Link
	e.GET("/redirecionadores/:codigo_hash/links", s.LinkReadByCodigoHash)
	e.GET("/redirecionadores/:codigo_hash/links/historico", s.HistoricoLinkReadAll)
	e.GET("/redirecionadores/:codigo_hash/links/:id", s.LinkReadById)
	e.POST("/redirecionadores/:codigo_hash/links", s.LinkCreate)
	e.PATCH("/redirecionadores/:codigo_hash/links/:id", s.LinkUpdate)
	e.DELETE("/redirecionadores/:codigo_hash/links/:id", s.LinkRemove)

	return e
}
