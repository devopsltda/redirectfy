package api

import (
	"github.com/TheDevOpsCorp/redirectify/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/TheDevOpsCorp/redirectify/docs"
	_ "github.com/joho/godotenv/autoload"
)

// @title API do Redirect Max
// @version 1.0.0
// @description API para interagir com o Redirect Max

// @contact.name Equipe da DevOps (Pablo, Guilherme e Eduardo)
// @contact.email test@test.com
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
	a.GET("/usuarios/:nome_de_usuario", UsuarioReadByNomeDeUsuario)
	a.POST("/usuarios", UsuarioCreate)
	a.PATCH("/usuarios/:nome_de_usuario/autentica", UsuarioAutenticado)
	a.PATCH("/usuarios/:nome_de_usuario", UsuarioUpdate)
	a.DELETE("/usuarios/:nome_de_usuario", UsuarioRemove)
	a.POST("/usuarios/login", UsuarioLogin)

	// API - Plano de Assinatura
	a.GET("/planos_de_assinatura", PlanoDeAssinaturaReadAll)
	a.GET("/planos_de_assinatura/:nome", PlanoDeAssinaturaReadByNome)
	a.POST("/planos_de_assinatura", PlanoDeAssinaturaCreate)
	a.PATCH("/planos_de_assinatura/:nome", PlanoDeAssinaturaUpdate)
	a.DELETE("/planos_de_assinatura/:nome", PlanoDeAssinaturaRemove)

	// API - Link
	a.GET("/links", LinkReadAll)
	a.GET("/links/:codigo_hash", LinkReadByCodigoHash)
	a.POST("/links", LinkCreate)
	a.PATCH("/links/:codigo_hash/rehash", LinkRehash)
	a.PATCH("/links/:codigo_hash", LinkUpdate)
	a.DELETE("/links/:codigo_hash", LinkRemove)

	// API - Histórico
	a.GET("/historico", HistoricoReadAll)
}
