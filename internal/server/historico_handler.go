package server

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirectify/internal/model"
	"github.com/TheDevOpsCorp/redirectify/internal/util"
	"github.com/labstack/echo/v4"
)

// HistoricoReadAll godoc
//
// @Summary Retorna o histórico de ações no sistema
// @Tags    Histórico
// @Accept  json
// @Produce json
// @Success 200 {object} []model.Historico
// @Failure 400 {object} Erro
// @Failure 500 {object} Erro
// @Router  /api/historico [get]
func (s *Server) HistoricoReadAll(c echo.Context) error {
	historico, err := model.HistoricoReadAll(s.db)

	if err != nil {
		log.Printf("HistoricoReadAll: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, historico)
}
