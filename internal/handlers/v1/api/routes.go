package api

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"redirectify/internal/auth"

	_ "github.com/joho/godotenv/autoload"
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
	a.POST("/usuarios/login", UsuarioLogin)
	a.PATCH("/usuarios/troca_de_senha/:valor", UsuarioTrocaDeSenha)
	a.POST("/usuarios/:nome_de_usuario/troca_de_senha", UsuarioTrocaDeSenhaExigir)

	// API - Autenticação
	a.PATCH("/autenticacao/:valor", UsuarioAutenticado)

	// API - Plano de Assinatura
	a.GET("/planos_de_assinatura", PlanoDeAssinaturaReadAll)
	a.GET("/planos_de_assinatura/historico", HistoricoPlanoDeAssinaturaReadAll)
	a.GET("/planos_de_assinatura/:nome", PlanoDeAssinaturaReadByNome)
	a.POST("/planos_de_assinatura", PlanoDeAssinaturaCreate)
	a.PATCH("/planos_de_assinatura/:nome", PlanoDeAssinaturaUpdate)
	a.DELETE("/planos_de_assinatura/:nome", PlanoDeAssinaturaRemove)

	// API - Link
	a.GET("/links", LinkReadAll)
	a.GET("/links/historico", HistoricoLinkReadAll)
	a.GET("/links/:codigo_hash", LinkReadByCodigoHash)
	a.POST("/links", LinkCreate)
	a.PATCH("/links/:codigo_hash/rehash", LinkRehash)
	a.PATCH("/links/:codigo_hash", LinkUpdate)
	a.DELETE("/links/:codigo_hash", LinkRemove)
}
