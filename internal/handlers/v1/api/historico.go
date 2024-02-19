package api

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirectify/internal/database"
	"github.com/TheDevOpsCorp/redirectify/internal/models"
	"github.com/TheDevOpsCorp/redirectify/internal/utils"
	"github.com/labstack/echo/v4"
)

// HistoricoReadAll godoc
//
// @Summary Retorna o histórico de ações no sistema
// @Tags    Histórico
// @Accept  json
// @Produce json
// @Success 200 {object} []models.Historico
// @Failure 400 {object} Erro
// @Failure 500 {object} Erro
// @Router  /api/historico [get]
func HistoricoReadAll(c echo.Context) error {
	historico, err := models.HistoricoReadAll(database.Db)

	if err != nil {
		log.Printf("HistoricoReadAll: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, historico)
}
