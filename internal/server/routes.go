package server

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"redirectify/cmd/resources"
	"redirectify/internal/handlers/v1/api"
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

	// Arquivos estáticos (CSS, JS, imagens, etc)
	StaticFileServer := http.FileServer(http.FS(resources.StaticFiles))
	e.GET("/static/*", echo.WrapHandler(StaticFileServer))

	// WEB
	e.GET("/login", echo.WrapHandler(templ.Handler(views.Login())))
	e.GET("/signup", echo.WrapHandler(templ.Handler(views.Signup())))
	e.GET("/account_created", echo.WrapHandler(templ.Handler(views.AccountCreated())))

	// Roteamento Dinâmico
	// e.GET("/to/:codigo_hash", echo.WrapHandler(http.HandlerFunc(web.LinkAccessWebHandler)))

	// API - v1
	v1 := e.Group("/v1")
	api.RegisterRoutesV1(v1)

	return e
}
