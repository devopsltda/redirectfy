package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"redirectify/internal/models"
	"redirectify/internal/services/database"
	"redirectify/internal/utils"
)

// HistoricoPlanoDeAssinaturaReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a planos de assinatura no sistema
// @Tags    Histórico
// @Accept  json
// @Produce json
// @Success 200 {object} []models.HistoricoPlanoDeAssinatura
// @Failure 400 {object} Erro
// @Failure 500 {object} Erro
// @Router  /v1/api/planos_de_assinatura/historico [get]
func HistoricoPlanoDeAssinaturaReadAll(c echo.Context) error {
	historico, err := models.HistoricoPlanoDeAssinaturaReadAll(database.Db)

	if err != nil {
		log.Printf("HistoricoPlanoDeAssinaturaReadAll: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, historico)
}

// HistoricoUsuarioReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a usuários no sistema
// @Tags    Histórico
// @Accept  json
// @Produce json
// @Success 200 {object} []models.HistoricoUsuario
// @Failure 400 {object} Erro
// @Failure 500 {object} Erro
// @Router  /v1/api/usuarios/historico [get]
func HistoricoUsuarioReadAll(c echo.Context) error {
	historico, err := models.HistoricoUsuarioReadAll(database.Db)

	if err != nil {
		log.Printf("HistoricoUsuarioReadAll: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, historico)
}

// HistoricoRedirecionadorReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a redirecionadores no sistema
// @Tags    Histórico
// @Accept  json
// @Produce json
// @Success 200 {object} []models.HistoricoRedirecionador
// @Failure 400 {object} Erro
// @Failure 500 {object} Erro
// @Router  /v1/api/redirecionadores/historico [get]
func HistoricoRedirecionadorReadAll(c echo.Context) error {
	historico, err := models.HistoricoRedirecionadorReadAll(database.Db)

	if err != nil {
		log.Printf("HistoricoRedirecionadorReadAll: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, historico)
}

// HistoricoLinkReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a links no sistema
// @Tags    Histórico
// @Accept  json
// @Produce json
// @Success 200 {object} []models.HistoricoLink
// @Failure 400 {object} Erro
// @Failure 500 {object} Erro
// @Router  /v1/api/links/historico [get]
func HistoricoLinkReadAll(c echo.Context) error {
	historico, err := models.HistoricoLinkReadAll(database.Db)

	if err != nil {
		log.Printf("HistoricoLinkReadAll: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, historico)
}
