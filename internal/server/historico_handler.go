package server

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirect-max/internal/model"
	"github.com/labstack/echo/v4"
)

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
