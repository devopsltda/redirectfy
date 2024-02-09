package server

import (
	"net/http"
	"strings"

	"github.com/TheDevOpsCorp/redirect-max/cmd/web"
	"github.com/TheDevOpsCorp/redirect-max/internal/auth"
	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/TheDevOpsCorp/redirect-max/docs"
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
			return strings.Contains(c.Path(), "/web")
		},
	}))

	// Arquivos estáticos (CSS, JS, imagens, etc)
	JSFileServer := http.FileServer(http.FS(web.JSFiles))
	CSSFileServer := http.FileServer(http.FS(web.CSSFiles))
	e.GET("/static/js/*", echo.WrapHandler(JSFileServer))
	e.GET("/static/css/*", echo.WrapHandler(CSSFileServer))

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
	a.GET("/swagger/*", echoSwagger.WrapHandler)

	// API - Usuário
	a.GET("/usuario", s.UsuarioReadAll)
	a.GET("/usuario/:nome_de_usuario", s.UsuarioReadByNomeDeUsuario)
	a.POST("/usuario", s.UsuarioCreate)
	a.PATCH("/usuario/:nome_de_usuario", s.UsuarioUpdate)
	a.DELETE("/usuario/:nome_de_usuario", s.UsuarioRemove)
	a.POST("/usuario/login", s.UsuarioLogin)

	// API - Plano de Assinatura
	a.GET("/plano_de_assinatura", s.PlanoDeAssinaturaReadAll)
	a.GET("/plano_de_assinatura/:nome", s.PlanoDeAssinaturaReadByNome)
	a.POST("/plano_de_assinatura", s.PlanoDeAssinaturaCreate)
	a.PATCH("/plano_de_assinatura/:nome", s.PlanoDeAssinaturaUpdate)
	a.DELETE("/plano_de_assinatura/:nome", s.PlanoDeAssinaturaRemove)

	// API - Link
	a.GET("/link", s.LinkReadAll)
	a.GET("/link/:codigo_hash", s.LinkReadByCodigoHash)
	a.POST("/link", s.LinkCreate)
	a.PATCH("/link/:codigo_hash", s.LinkUpdate)
	a.DELETE("/link/:codigo_hash", s.LinkRemove)

	// API - Histórico
	a.GET("/historico", s.HistoricoReadAll)

	// WEB
	w := e.Group("/web")
	w.GET("", echo.WrapHandler(templ.Handler(web.HelloForm())))
	w.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	return e
}
