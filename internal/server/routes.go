package server

import (
	"net/http"
	"strings"

	"redirectify/internal/handlers/v1/api"

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

	// API - v1
	v1 := e.Group("/v1")
	api.RegisterRoutesV1(v1)

	return e
}
