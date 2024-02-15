package server

import (
	"net/http"
	"strings"

	"github.com/TheDevOpsCorp/redirectify/cmd/resources"
	"github.com/TheDevOpsCorp/redirectify/internal/handlers/api"
	"github.com/TheDevOpsCorp/redirectify/internal/handlers/web"
	"github.com/TheDevOpsCorp/redirectify/internal/auth"
	"github.com/TheDevOpsCorp/redirectify/internal/views"
	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/TheDevOpsCorp/redirectify/docs"
	_ "github.com/joho/godotenv/autoload"
)

// @title API do Redirect Max
// @version 0.0.0-alpha
// @description API para interagir com o Redirect Max

// @contact.name Equipe da DevOps (Pablo, Guilherme e Eduardo)
// @contact.email test@test.com
func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/web") || strings.Contains(c.Path(), "/docs")
		},
	}))

	// Roteamento Dinâmico
	e.GET("/:codigo_hash", echo.WrapHandler(http.HandlerFunc(web.LinkAccessWebHandler)))

	// Arquivos estáticos (CSS, JS, imagens, etc)
	StaticFileServer := http.FileServer(http.FS(resources.StaticFiles))
	e.GET("/static/*", echo.WrapHandler(StaticFileServer))

	// API - Documentação Swagger
	e.GET("/docs/*", echoSwagger.WrapHandler)

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

	// API - Usuário
	a.GET("/usuario", api.UsuarioReadAll)
	a.GET("/usuario/:nome_de_usuario", api.UsuarioReadByNomeDeUsuario)
	a.POST("/usuario", api.UsuarioCreate)
	a.PATCH("/usuario/:nome_de_usuario", api.UsuarioUpdate)
	a.DELETE("/usuario/:nome_de_usuario", api.UsuarioRemove)
	a.POST("/usuario/login", api.UsuarioLogin)

	// API - Plano de Assinatura
	a.GET("/plano_de_assinatura", api.PlanoDeAssinaturaReadAll)
	a.GET("/plano_de_assinatura/:nome", api.PlanoDeAssinaturaReadByNome)
	a.POST("/plano_de_assinatura", api.PlanoDeAssinaturaCreate)
	a.PATCH("/plano_de_assinatura/:nome", api.PlanoDeAssinaturaUpdate)
	a.DELETE("/plano_de_assinatura/:nome", api.PlanoDeAssinaturaRemove)

	// API - Link
	a.GET("/link", api.LinkReadAll)
	a.GET("/link/:codigo_hash", api.LinkReadByCodigoHash)
	a.POST("/link", api.LinkCreate)
	a.PATCH("/link/:codigo_hash", api.LinkUpdate)
	a.DELETE("/link/:codigo_hash", api.LinkRemove)

	// API - Histórico
	a.GET("/historico", api.HistoricoReadAll)

	// WEB
	w := e.Group("/web")

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
	w.GET("/hey", echo.WrapHandler(templ.Handler(views.Landpage())))
	w.GET("/main", echo.WrapHandler(http.HandlerFunc(web.MainWebHandler)))
	w.POST("/link_create", echo.WrapHandler(http.HandlerFunc(web.LinkCreateWebHandler)))
	w.GET("/link_create_form", echo.WrapHandler(http.HandlerFunc(web.LinkCreateFormWebHandler)))
	w.GET("/link_create_button", echo.WrapHandler(http.HandlerFunc(web.LinkCreateButtonWebHandler)))

	return e
}
