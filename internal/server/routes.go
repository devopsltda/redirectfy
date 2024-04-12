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
		AllowCredentials: true,
		ExposeHeaders: []string{"Set-Cookie"},
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
	e.GET("/u/:username", s.UsuarioReadByNomeDeUsuario)
	e.PATCH("/u/:username", s.UsuarioUpdate)

	// API - Kirvano
	e.POST("/kirvano/to_user/:hash", s.KirvanoToUser)
	e.POST("/kirvano", s.KirvanoCreate)

	// API - Autenticação
	e.POST("/u/login", s.UsuarioLogin)
	e.POST("/u/logout", s.UsuarioLogout)
	e.PATCH("/u/change_password/:hash", s.UsuarioTrocaDeSenha)
	e.PATCH("/u/:username/change_password", s.UsuarioTrocaDeSenhaExigir)

	// API - Plano de Assinatura
	e.GET("/pricing", s.PlanoDeAssinaturaReadAll)
	e.GET("/pricing/:name", s.PlanoDeAssinaturaReadByNome)
	e.POST("/pricing", s.PlanoDeAssinaturaCreate)
	e.PATCH("/pricing/:name", s.PlanoDeAssinaturaUpdate)
	e.DELETE("/pricing/:name", s.PlanoDeAssinaturaRemove)

	// API - Redirecionador
	e.GET("/r", s.RedirecionadorReadAll)
	e.GET("/r/:hash", s.RedirecionadorReadByCodigoHash)
	e.POST("/r", s.RedirecionadorCreate)
	e.PATCH("/r/:hash/refresh", s.RedirecionadorRehash)
	e.PATCH("/r/:hash", s.RedirecionadorUpdate)
	e.DELETE("/r/:hash", s.RedirecionadorRemove)

	// API - Link
	e.GET("/r/:hash/links", s.LinkReadByCodigoHash)
	e.GET("/r/:hash/links/:id", s.LinkReadById)
	e.POST("/r/:hash/links", s.LinkCreate)
	e.PATCH("/r/:hash/links/:id", s.LinkUpdate)
	e.DELETE("/r/:hash/links/:id", s.LinkRemove)

	// API - Admin
	e.GET("/admin/user_history", s.UserHistory)
	e.GET("/admin/redirect_history", s.RedirectHistory)
	e.GET("/admin/kirvano_history", s.KirvanoHistory)
	e.GET("/admin/pricing_history", s.PricingHistory)
	e.GET("/admin/link_history", s.LinkHistory)

	return e
}
