package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirect-max/internal/model"
	"github.com/labstack/echo/v4"
)

func (s *Server) PlanoDeAssinaturaReadByNome(c echo.Context) error {
	var planoDeAssinatura model.PlanoDeAssinatura

	row := s.db.QueryRow(
		"SELECT * FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL AND NOME = $1",
		c.Param("nome"),
	)

	if err := row.Scan(
		&planoDeAssinatura.Id,
		&planoDeAssinatura.Nome,
		&planoDeAssinatura.ValorMensal,
		&planoDeAssinatura.CriadoEm,
		&planoDeAssinatura.AtualizadoEm,
		&planoDeAssinatura.RemovidoEm,
	); err != nil {
		log.Printf("PlanoDeAssinaturaReadByNome: %v", err)
		return err
	}

	if err := row.Err(); err != nil {
		log.Printf("PlanoDeAssinaturaReadByNome: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, planoDeAssinatura)
}

func (s *Server) PlanoDeAssinaturaReadAll(c echo.Context) error {
	var planosDeAssinatura []model.PlanoDeAssinatura

	rows, err := s.db.Query("SELECT * FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL")

	if err != nil {
		log.Printf("PlanoDeAssinaturaReadAll: %v", err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var planoDeAssinatura model.PlanoDeAssinatura

		if err := rows.Scan(
			&planoDeAssinatura.Id,
			&planoDeAssinatura.Nome,
			&planoDeAssinatura.ValorMensal,
			&planoDeAssinatura.CriadoEm,
			&planoDeAssinatura.AtualizadoEm,
			&planoDeAssinatura.RemovidoEm,
		); err != nil {
			log.Printf("PlanoDeAssinaturaReadAll: %v", err)
			return err
		}

		planosDeAssinatura = append(planosDeAssinatura, planoDeAssinatura)
	}

	if err := rows.Err(); err != nil {
		log.Printf("PlanoDeAssinaturaReadAll: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, planosDeAssinatura)
}

func (s *Server) PlanoDeAssinaturaCreate(c echo.Context) error {
	var planoDeAssinatura model.PlanoDeAssinatura

	if err := c.Bind(&planoDeAssinatura); err != nil {
		log.Printf("PlanoDeAssinaturaCreate: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Mensagem": "Requisição teve algum erro",
		})
	}

	_, err := s.db.Exec(
		"INSERT INTO PLANO_DE_ASSINATURA (NOME, VALOR_MENSAL, REMOVIDO_EM) VALUES ($1, $2, $3)",
		planoDeAssinatura.Nome,
		planoDeAssinatura.ValorMensal,
		nil,
	)

	if err != nil {
		log.Printf("PlanoDeAssinaturaCreate: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Plano de assinatura adicionado com sucesso",
	})
}

func (s *Server) PlanoDeAssinaturaUpdate(c echo.Context) error {
	parametros := struct {
		Nome        string `json:"nome"`
		ValorMensal int64  `json:"valor_mensal"`
	}{}

	if err := c.Bind(&parametros); err != nil {
		return err
	}

	sqlQuery := "UPDATE PLANO_DE_ASSINATURA SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if parametros.Nome != "" {
		sqlQuery += ", SET NOME = '" + parametros.Nome + "'"
	}

	if parametros.ValorMensal != 0 {
		sqlQuery += ", SET VALOR_MENSAL = " + fmt.Sprint(parametros.ValorMensal)
	}

	sqlQuery += " WHERE NOME = $1"

	_, err := s.db.Exec(
		sqlQuery,
		c.Param("nome"),
	)

	if err != nil {
		log.Printf("PlanoDeAssinaturaUpdate: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Plano de assinatura atualizado com sucesso",
	})
}

func (s *Server) PlanoDeAssinaturaRemove(c echo.Context) error {
	_, err := s.db.Exec(
		"UPDATE PLANO_DE_ASSINATURA SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE NOME = $1",
		c.Param("nome"),
	)

	if err != nil {
		log.Printf("PlanoDeAssinaturaRemove: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Plano de assinatura removido com sucesso",
	})
}
