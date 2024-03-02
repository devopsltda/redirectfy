package server

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"redirectify/cmd/resources"
	"redirectify/internal/handlers/v1/api"
	"redirectify/internal/handlers/web"
	"redirectify/internal/views"

	_ "github.com/joho/godotenv/autoload"
)

// @title API do Redirectify
// @version 1.0.0
// @description API para interagir com o Redirectify

// @contact.name Equipe da DevOps (Pablo Eduardo, Guilherme Bernardo e Eduardo Henrique)
// @contact.email comercialdevops@gmail.com
func RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/api/docs")
		},
	}))

	// WEB

	// Telas
	// - Landpage
	// - Tela de Pricing
	// - Tela de Contato
	// - Tela de Login
	// - Tela de Signup
	// - Tela de Esqueci Minha Senha
	// - Tela de Criar Nova Senha
	// - Tela de Confirmação de Mudança de Senha
	// - Tela de Editar Usuário
	// - Tela de Contratar Plano
	// - Tela de Cancelar Plano
	// - Tela de Gerenciamento de Links
	e.GET("/hey", echo.WrapHandler(templ.Handler(views.Landpage())))
	e.GET("/main", echo.WrapHandler(http.HandlerFunc(web.MainWebHandler)))
	e.POST("/link_create", echo.WrapHandler(http.HandlerFunc(web.LinkCreateWebHandler)))
	e.GET("/link_create_form", echo.WrapHandler(http.HandlerFunc(web.LinkCreateFormWebHandler)))
	e.GET("/link_create_button", echo.WrapHandler(http.HandlerFunc(web.LinkCreateButtonWebHandler)))

	// Roteamento Dinâmico
	e.GET("/to/:codigo_hash", echo.WrapHandler(http.HandlerFunc(web.LinkAccessWebHandler)))

	// Arquivos estáticos (CSS, JS, imagens, etc)
	StaticFileServer := http.FileServer(http.FS(resources.StaticFiles))
	e.GET("/static/*", echo.WrapHandler(StaticFileServer))

	// API - v1
	v1 := e.Group("/v1")
	api.RegisterRoutesV1(v1)

	return e
}
