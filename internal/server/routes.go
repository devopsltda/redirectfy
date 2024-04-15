package server

import (
	"fmt"
	"net/http"
	"strings"

	"redirectfy/internal/auth"
	"redirectfy/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

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
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Set-Cookie"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/docs")
		},
	}))

	e.Use(auth.TokenRefreshMiddleware)
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(auth.ChaveDeAcesso),
		TokenLookup: "cookie:access-token",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.Claims)
		},
		Skipper: auth.PathWithNoAuthRequired,
		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Println(err.Error())
			switch err.Error() {
			case echojwt.ErrJWTInvalid.Message.(string):
				utils.ErroLog("EchoJwtErrorHandler", "Erro na validação do token JWT", err)
				return utils.Erro(http.StatusUnauthorized, "Por favor, forneça um token JWT válido.")
			case echojwt.ErrJWTMissing.Message.(string):
				utils.ErroLog("EchoJwtErrorHandler", "Erro na busca pelo token JWT", err)
				return utils.Erro(http.StatusUnauthorized, "Por favor, forneça um token JWT válido.")
			case "missing value in cookies":
				utils.ErroLog("EchoJwtErrorHandler", "Erro de cookies sem valores válidos", err)
				return utils.Erro(http.StatusUnauthorized, "Por favor, forneça os cookies com tokens JWT válidos.")
			default:
				return nil
			}
		},
	}))
	e.Use(auth.IsUserTheSameMiddleware)

	// API - Documentação Swagger
	e.GET("/docs/*", echoSwagger.WrapHandler)

	// API - Obter Links do Redirecionador
	e.GET("/to/:hash", s.RedirecionadorLinksToGoTo)

	// API - Usuário
	e.GET("/u/:username", s.UsuarioReadByNomeDeUsuario)

	// API - Kirvano
	e.POST("/kirvano/to_user/:hash", s.KirvanoToUser)
	e.POST("/kirvano", s.KirvanoCreate)

	// API - Autenticação
	e.POST("/u/login", s.UsuarioLogin)
	e.POST("/u/logout", s.UsuarioLogout)
	e.PATCH("/u/change_password/:hash", s.UsuarioTrocaDeSenha)
	e.PATCH("/u/:username/change_password", s.UsuarioSolicitarTrocaDeSenha)

	// API - Plano de Assinatura
	e.GET("/pricing", s.PlanoDeAssinaturaReadAll)
	e.GET("/pricing/:name", s.PlanoDeAssinaturaReadByNome)

	// API - Redirecionador
	e.GET("/r", s.RedirecionadorReadAll)
	e.GET("/r/:hash", s.RedirecionadorReadByCodigoHash)
	e.POST("/r", s.RedirecionadorCreate)
	e.PATCH("/r/:hash/refresh", auth.PricingMiddleware(s.RedirecionadorRefresh))
	e.PATCH("/r/:hash", s.RedirecionadorUpdate)
	e.DELETE("/r/:hash", s.RedirecionadorRemove)

	// API - Link
	e.GET("/r/:hash/links", s.LinkReadByCodigoHash)
	e.GET("/r/:hash/links/:id", s.LinkReadById)
	e.POST("/r/:hash/links", s.LinkCreate)
	e.PATCH("/r/:hash/links/:id", s.LinkUpdate)
	e.DELETE("/r/:hash/links/:id", s.LinkRemove)

	// API - Admin
	// e.GET("/admin/user_history", s.UserHistory)
	// e.GET("/admin/redirect_history", s.RedirectHistory)
	// e.GET("/admin/kirvano_history", s.KirvanoHistory)
	// e.GET("/admin/pricing_history", s.PricingHistory)
	// e.GET("/admin/link_history", s.LinkHistory)

	return e
}
