package server

import (
	"net/http"
	"strings"

	"github.com/TheDevOpsCorp/redirect-max/cmd/web"
	"github.com/a-h/templ"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Validate = validator.New()

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "web")
		},
	}))

	// API
	a := e.Group("/api")

	// API - Usuário
	a.GET("/usuario", s.UsuarioReadAll)
	a.GET("/usuario/:nome_de_usuario", s.UsuarioReadByNomeDeUsuario)
	a.POST("/usuario", s.UsuarioCreate)
	a.PATCH("/usuario/:nome_de_usuario", s.UsuarioUpdate)
	a.DELETE("/usuario/:nome_de_usuario", s.UsuarioRemove)

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
	fileServer := http.FileServer(http.FS(web.Files))
	w.GET("/js/*", echo.WrapHandler(fileServer))
	w.GET("/", echo.WrapHandler(templ.Handler(web.HelloForm())))
	w.POST("/hello", echo.WrapHandler(http.HandlerFunc(web.HelloWebHandler)))

	return e
}
