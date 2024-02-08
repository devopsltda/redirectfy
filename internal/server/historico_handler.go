package server

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirect-max/internal/model"
	"github.com/labstack/echo/v4"
)

// HistoricoReadAll godoc
//
// @Summary Retorna o histórico de ações no sistema
// @Tags    Histórico
// @Accept  json
// @Produce json
// @Success 200 {object} []model.Historico
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router  /api/historico [get]
func (s *Server) HistoricoReadAll(c echo.Context) error {
	var historico []model.Historico

	rows, err := s.db.Query("SELECT * FROM HISTORICO")

	if err != nil {
		log.Printf("HistoricoReadAll: %v", err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var registro model.Historico

		if err := rows.Scan(
			&registro.Id,
			&registro.Usuario,
			&registro.ValorOriginal,
			&registro.ValorNovo,
			&registro.TabelaModificada,
			&registro.CriadoEm,
		); err != nil {
			log.Printf("HistoricoReadAll: %v", err)
			return err
		}

		historico = append(historico, registro)
	}

	if err := rows.Err(); err != nil {
		log.Printf("HistoricoReadAll: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, historico)
}
