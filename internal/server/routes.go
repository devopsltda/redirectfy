package server

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"redirectify/internal/auth"

	_ "redirectify/docs"
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
			return strings.Contains(c.Path(), "/api/docs")
		},
	}))

	// API - v1
	v1 := e.Group("/v1")

	// API
	a := v1.Group("/api")
	a.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(auth.ChaveDeAcesso),
		TokenLookup: "cookie:access-token",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.Claims)
		},
		Skipper: auth.PathWithNoAuthRequired,
	}))

	// API - Documentação Swagger
	a.GET("/docs/*", echoSwagger.WrapHandler)

	// API - Usuário
	a.GET("/usuarios", s.UsuarioReadAll)
	a.GET("/usuarios/historico", s.HistoricoUsuarioReadAll)
	a.GET("/usuarios/:nome_de_usuario", s.UsuarioReadByNomeDeUsuario)
	a.POST("/usuarios_temporarios", s.UsuarioTemporarioCreate)
	a.POST("/usuarios_temporarios/historico", s.HistoricoUsuarioTemporarioReadAll)
	a.PATCH("/usuarios/:nome_de_usuario", s.UsuarioUpdate)
	a.DELETE("/usuarios/:nome_de_usuario", s.UsuarioRemove)

	// API - Autenticação
	a.POST("/usuarios/login", s.UsuarioLogin)
	a.PATCH("/usuarios/troca_de_senha/:valor", s.UsuarioTrocaDeSenha)
	a.POST("/usuarios/:nome_de_usuario/troca_de_senha", s.UsuarioTrocaDeSenhaExigir)
	a.PATCH("/autenticacao/:valor", s.UsuarioAutenticado)

	// API - Plano de Assinatura
	a.GET("/planos_de_assinatura", s.PlanoDeAssinaturaReadAll)
	a.GET("/planos_de_assinatura/historico", s.HistoricoPlanoDeAssinaturaReadAll)
	a.GET("/planos_de_assinatura/:nome", s.PlanoDeAssinaturaReadByNome)
	a.POST("/planos_de_assinatura", s.PlanoDeAssinaturaCreate)
	a.PATCH("/planos_de_assinatura/:nome", s.PlanoDeAssinaturaUpdate)
	a.DELETE("/planos_de_assinatura/:nome", s.PlanoDeAssinaturaRemove)

	// API - Redirecionador
	a.GET("/redirecionadores", s.RedirecionadorReadAll)
	a.GET("/redirecionadores/historico", s.HistoricoRedirecionadorReadAll)
	a.GET("/redirecionadores/:codigo_hash", s.RedirecionadorReadByCodigoHash)
	a.POST("/redirecionadores", s.RedirecionadorCreate)
	a.PATCH("/redirecionadores/:codigo_hash/rehash", s.RedirecionadorRehash)
	a.PATCH("/redirecionadores/:codigo_hash", s.RedirecionadorUpdate)
	a.DELETE("/redirecionadores/:codigo_hash", s.RedirecionadorRemove)

	// API - Link
	a.GET("/redirecionadores/:codigo_hash/links", s.LinkReadByCodigoHash)
	a.GET("/redirecionadores/:codigo_hash/links/historico", s.HistoricoLinkReadAll)
	a.GET("/redirecionadores/:codigo_hash/links/:id", s.LinkReadById)
	a.POST("/redirecionadores/:codigo_hash/links", s.LinkCreate)
	a.PATCH("/redirecionadores/:codigo_hash/links/:id", s.LinkUpdate)
	a.DELETE("/redirecionadores/:codigo_hash/links/:id", s.LinkRemove)

	return e
}
