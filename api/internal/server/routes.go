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

// @title API do Redirectfy
//
// @version 1.0.0
//
// @description API para interagir com o Redirectfy
//
// @contact.name Equipe da DevOps (Pablo Eduardo, Guilherme Bernardo e Eduardo Henrique)
//
// @contact.email comercialdevops@gmail.com
func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	// Middleware de logs
	e.Use(middleware.Logger())

	// Middleware que adiciona headers que visam aumentar
	// a segurança da API
	e.Use(middleware.Secure())

	// Middleware que reinicia o servidor em caso de panic
	e.Use(middleware.Recover())

	// Middleware de compressão de dados
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "/docs/index.html")
		},
	}))

	// Essa é a única rota que não passa pelo CORS, porque não sabemos qual
	// o servidor da Kirvano e se ele vai ser persistente. A autenticidade
	// da requisição deve ser conferido pelo token enviado por eles,
	// e essa verificação é realizada pelo middleware abaixo.
	e.POST("/api/kirvano", s.KirvanoCreate, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			val, ok := c.Request().Header["Security-Token"]

			if !ok || val[0] != utils.KirvanoToken {
				utils.DebugLog("KirvanoMiddleware", "Token enviado pela Kirvano não corresponde ao KIRVANO_TOKEN", nil)
				return utils.Erro(http.StatusUnauthorized, "Você não tem permissão para enviar requisições para esse endpoint.")
			}

			return next(c)
		}
	})

	api := e.Group("/api")

	// Middleware de CORS
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{
			"http://localhost", // Ambiente de desenvolvimento do Angular
			"http://localhost:80", // Ambiente de desenvolvimento do Angular
			"http://localhost:4200", // Ambiente de desenvolvimento do Angular
			"https://redirectfy.fly.dev", // Ambiente de homologação da aplicação
		},
		AllowCredentials: true,
	}))

	// Middleware de autenticação (mais especificamente, do que diz respeito
	// ao token de acesso)
	api.Use(echojwt.WithConfig(echojwt.Config{
		ContextKey:  "usuario",
		SigningKey:  []byte(auth.ChaveDeAcesso),
		TokenLookup: "cookie:access-token",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.Claims)
		},
		Skipper: auth.PathWithNoAuthRequired,
		ErrorHandler: func(c echo.Context, err error) error {
			switch err.Error() {
			case echojwt.ErrJWTInvalid.Message.(string):
				utils.DebugLog("EchoJwtErrorHandler", "Erro na validação do token JWT", err)
				return utils.Erro(http.StatusUnauthorized, "Por favor, forneça um token JWT válido.")
			case echojwt.ErrJWTMissing.Message.(string):
				utils.DebugLog("EchoJwtErrorHandler", "Erro na busca pelo token JWT", err)
				return utils.Erro(http.StatusUnauthorized, "Por favor, forneça um token JWT válido.")
			case "missing value in cookies":
				utils.DebugLog("EchoJwtErrorHandler", "Erro de cookies sem valores válidos", err)
				utils.DebugLog("EchoJwtErrorHandler", fmt.Sprintf("%+v", c.Cookies()), err)
				return utils.Erro(http.StatusUnauthorized, "Por favor, forneça os cookies com tokens JWT válidos.")
			default:
				return nil
			}
		},
	}))

	// Middleware que gera mais tokens de acesso baseado na validade do token de
	// atualização
	api.Use(auth.TokenRefreshMiddleware)

	// API - Documentação Swagger
	api.GET("/docs/*", echoSwagger.WrapHandler)

	// API - Obter Links do Redirecionador
	api.GET("/to/:hash", s.RedirecionadorLinksToGoTo)

	// API - Usuário
	api.GET("/u", s.UsuarioReadByNomeDeUsuario)

	// API - Kirvano
	api.POST("/kirvano/to_user/:hash", s.KirvanoToUser)

	// API - Autenticação
	api.POST("/u/login", s.UsuarioLogin)
	api.POST("/u/logout", s.UsuarioLogout)
	api.PATCH("/u/change_password/:hash", s.UsuarioTrocaDeSenha)
	api.PATCH("/u/:username/change_password", s.UsuarioSolicitarTrocaDeSenha)

	// API - Plano de Assinatura
	api.GET("/pricing", s.PlanoDeAssinaturaReadAll)
	api.GET("/pricing/:name", s.PlanoDeAssinaturaReadByNome)

	// API - Redirecionador
	api.GET("/r", s.RedirecionadorReadAll)
	api.GET("/r/:hash", s.RedirecionadorReadByCodigoHash)
	api.POST("/r", s.RedirecionadorCreate)

	// Rota de rehash tem um middleware específico para verificar se o usuário é
	// de um plano Pro que tenha acesso a essa funcionalidade.
	api.PATCH("/r/:hash/refresh", s.RedirecionadorRefresh, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Get("usuario") == nil {
				utils.DebugLog("PricingMiddleware", "Erro ao ler o contexto 'usuario'", nil)
				return utils.Erro(http.StatusBadRequest, "Você não contém um ou mais dos cookies necessários para autenticação.")
			}

			if !strings.HasPrefix(c.Get("usuario").(*jwt.Token).Claims.(*auth.Claims).PlanoDeAssinatura, "Pro") {
				utils.DebugLog("PricingMiddleware", "O usuário não tem o plano de assinatura apropriado para usar o rehash", nil)
				return utils.Erro(http.StatusPaymentRequired, "O seu plano de assinatura não oferece o recurso de rehash.")
			}

			return next(c)
		}
	})

	api.PATCH("/r/:hash", s.RedirecionadorUpdate)
	api.DELETE("/r/:hash", s.RedirecionadorRemove)

	// API - Link
	api.GET("/r/:hash/links", s.LinkReadByCodigoHash)
	api.GET("/r/:hash/links/:id", s.LinkReadById)
	api.POST("/r/:hash/links", s.LinkCreate)
	api.PATCH("/r/:hash/links/:id", s.LinkUpdate)
	api.PATCH("/r/:hash/links/:id/enable", s.LinkEnable)
	api.PATCH("/r/:hash/links/:id/disable", s.LinkDisable)
	api.DELETE("/r/:hash/links/:id", s.LinkRemove)

	return e
}
