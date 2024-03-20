package server

import (
	"net/http"
	"strings"

	"redirectify/cmd/resources"
	"redirectify/internal/handlers/v1/api"
	"redirectify/internal/handlers/web"
	"redirectify/internal/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.GET("/account_validated", echo.WrapHandler(templ.Handler(views.AccountValidated())))
	e.GET("/account_not_validated", echo.WrapHandler(templ.Handler(views.AccountNotValidated())))
	e.GET("/forgot_password", echo.WrapHandler(templ.Handler(views.ForgotPassword())))
	e.GET("/change_password", echo.WrapHandler(templ.Handler(views.ChangePassword())))
	e.GET("/password_forgotten", echo.WrapHandler(templ.Handler(views.PasswordForgotten())))
	e.GET("/changed_password", echo.WrapHandler(templ.Handler(views.PasswordChanged())))
	e.GET("/choose_plan", echo.WrapHandler(templ.Handler(views.SelectPlan())))
	e.GET("/base_test", echo.WrapHandler(templ.Handler(views.BaseTest())))
	e.GET("/novo_link", echo.WrapHandler(templ.Handler(views.NovoLink())))
	e.GET("/novo_link_whatsapp", echo.WrapHandler(templ.Handler(views.AdicionarLinkWhatsapp())))
	e.GET("/novo_link_telegram", echo.WrapHandler(templ.Handler(views.AdicionarLinkTelegram())))
	e.GET("/nome_redirect", echo.WrapHandler(templ.Handler(views.NameRedirect())))
	e.GET("/not_found", echo.WrapHandler(templ.Handler(views.Error404())))
	e.GET("/erro_5xx", echo.WrapHandler(templ.Handler(views.Error5xx())))

	// Roteamento Dinâmico
	e.GET("/to/:codigo_hash", web.GoToLinkWebHandler)

	// API - v1
	v1 := e.Group("/v1")
	api.RegisterRoutesV1(v1)

	return e
}
