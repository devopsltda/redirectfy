package api

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"redirectify/internal/auth"

	_ "redirectify/docs"
)

// @title API do Redirectify
// @version 1.0.0
// @description API para interagir com o Redirectify

// @contact.name Equipe da DevOps (Pablo Eduardo, Guilherme Bernardo e Eduardo Henrique)
// @contact.email comercialdevops@gmail.com
func RegisterRoutesV1(e *echo.Group) {
	// API
	a := e.Group("/api")
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
	a.GET("/usuarios", UsuarioReadAll)
	a.GET("/usuarios/historico", HistoricoUsuarioReadAll)
	a.GET("/usuarios/:nome_de_usuario", UsuarioReadByNomeDeUsuario)
	a.POST("/usuarios", UsuarioCreate)
	a.PATCH("/usuarios/:nome_de_usuario", UsuarioUpdate)
	a.DELETE("/usuarios/:nome_de_usuario", UsuarioRemove)

	// API - Autenticação
	a.POST("/usuarios/login", UsuarioLogin)
	a.PATCH("/usuarios/troca_de_senha/:valor", UsuarioTrocaDeSenha)
	a.POST("/usuarios/:nome_de_usuario/troca_de_senha", UsuarioTrocaDeSenhaExigir)
	a.PATCH("/autenticacao/:valor", UsuarioAutenticado)

	// API - Plano de Assinatura
	a.GET("/planos_de_assinatura", PlanoDeAssinaturaReadAll)
	a.GET("/planos_de_assinatura/historico", HistoricoPlanoDeAssinaturaReadAll)
	a.GET("/planos_de_assinatura/:nome", PlanoDeAssinaturaReadByNome)
	a.POST("/planos_de_assinatura", PlanoDeAssinaturaCreate)
	a.PATCH("/planos_de_assinatura/:nome", PlanoDeAssinaturaUpdate)
	a.DELETE("/planos_de_assinatura/:nome", PlanoDeAssinaturaRemove)

	// API - Redirecionador
	a.GET("/redirecionadores", RedirecionadorReadAll)
	a.GET("/redirecionadores/historico", HistoricoRedirecionadorReadAll)
	a.GET("/redirecionadores/:codigo_hash", RedirecionadorReadByCodigoHash)
	a.POST("/redirecionadores", RedirecionadorCreate)
	a.PATCH("/redirecionadores/:codigo_hash/rehash", RedirecionadorRehash)
	a.PATCH("/redirecionadores/:codigo_hash", RedirecionadorUpdate)
	a.DELETE("/redirecionadores/:codigo_hash", RedirecionadorRemove)

	// API - Link
	a.GET("/redirecionadores/:codigo_hash/links", LinkReadAll)
	a.GET("/redirecionadores/:codigo_hash/links/historico", HistoricoLinkReadAll)
	a.GET("/redirecionadores/:codigo_hash/links/:id", LinkReadById)
	a.POST("/redirecionadores/:codigo_hash/links", LinkCreate)
	a.PATCH("/redirecionadores/:codigo_hash/links/:id", LinkUpdate)
	a.DELETE("/redirecionadores/:codigo_hash/links/:id", LinkRemove)
}
