package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirect-max/internal/model"
	"github.com/labstack/echo/v4"
)

// PlanoDeAssinaturaReadByNome godoc
//
// @Summary Retorna o plano de assinatura com o nome fornecido
// @Tags    Plano de Assinatura
// @Accept  json
// @Produce json
// @Param   nome path     string  true  "Nome"
// @Success 200  {object} model.PlanoDeAssinatura
// @Failure 400  {object} Erro
// @Failure 500  {object} Erro
// @Router  /api/plano_de_assinatura/:nome [get]
func (s *Server) PlanoDeAssinaturaReadByNome(c echo.Context) error {
	/*** Parâmetros ***/
	parametros := struct {
		Nome string `param:"nome"`
	}{}
	/*** Parâmetros ***/

	/*** Validação ***/
	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome do plano de assinatura no parâmetro 'nome'.")
	}

	if len(erros) > 0 {
		return ErroValidacaoParametro(erros)
	}
	/*** Validação ***/

	/*** Banco de Dados ***/
	var planoDeAssinatura model.PlanoDeAssinatura

	row := s.db.QueryRow(
		"SELECT * FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL AND NOME = $1",
		parametros.Nome,
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
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, planoDeAssinatura)
}

// PlanoDeAssinaturaReadAll godoc
//
// @Summary Retorna os planos de assinatura
// @Tags    Plano de Assinatura
// @Accept  json
// @Produce json
// @Success 200  {object} []model.PlanoDeAssinatura
// @Failure 400  {object} Erro
// @Failure 500  {object} Erro
// @Router  /api/plano_de_assinatura [get]
func (s *Server) PlanoDeAssinaturaReadAll(c echo.Context) error {
	/*** Parâmetros ***/
	/*** Parâmetros ***/

	/*** Validação ***/
	/*** Validação ***/

	/*** Banco de Dados ***/
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
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, planosDeAssinatura)
}

// PlanoDeAssinaturaCreate godoc
//
// @Summary Cria um plano de assinatura
// @Tags    Plano de Assinatura
// @Accept  json
// @Produce json
// @Param   nome         body     string true "Nome"
// @Param   valor_mensal body     int    true "Valor Mensal"
// @Success 200          {object} map[string]string
// @Failure 400          {object} Erro
// @Failure 500          {object} Erro
// @Router  /api/plano_de_assinatura [post]
func (s *Server) PlanoDeAssinaturaCreate(c echo.Context) error {
	/*** Parâmetros ***/
	parametros := struct {
		Nome        string `json:"nome"`
		ValorMensal int64  `json:"valor_mensal"`
	}{}
	/*** Parâmetros ***/

	/*** Validação ***/
	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome do plano de assinatura no parâmetro 'nome'.")
	}

	if err := Validate.Var(parametros.ValorMensal, "gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um valor mensal válido no parâmetro 'valor_mensal'.")
	}

	if parametros.Nome == "" {
		erros = append(erros, "Por favor, forneça um nome válido para o parâmetro 'nome'.")
	}

	if len(erros) > 0 {
		return ErroValidacaoParametro(erros)
	}
	/*** Validação ***/

	/*** Banco de Dados ***/
	_, err := s.db.Exec(
		"INSERT INTO PLANO_DE_ASSINATURA (NOME, VALOR_MENSAL, REMOVIDO_EM) VALUES ($1, $2, $3)",
		parametros.Nome,
		parametros.ValorMensal,
		nil,
	)

	if err != nil {
		log.Printf("PlanoDeAssinaturaCreate: %v", err)
		return err
	}
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Plano de assinatura adicionado com sucesso",
	})
}

// PlanoDeAssinaturaUpdate godoc
//
// @Summary Atualiza um plano de assinatura
// @Tags    Plano de Assinatura
// @Accept  json
// @Produce json
// @Param   nome         path     string true  "Nome"
// @Param   nome         body     string false "Nome"
// @Param   valor_mensal body     int    false "Valor Mensal"
// @Success 200          {object} map[string]string
// @Failure 400          {object} Erro
// @Failure 500          {object} Erro
// @Router  /api/plano_de_assinatura/:nome [patch]
func (s *Server) PlanoDeAssinaturaUpdate(c echo.Context) error {
	/*** Parâmetros ***/
	parametros := struct {
		Nome        string `json:"nome"`
		NomeParam   string `param:"nome"`
		ValorMensal int64  `json:"valor_mensal"`
	}{}
	/*** Parâmetros ***/

	/*** Validação ***/
	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome do plano de assinatura no parâmetro 'nome' e o valor mensal no parâmetro 'valor_mensal'.")
	}

	if err := Validate.Var(parametros.ValorMensal, "gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um valor mensal válido no parâmetro 'valor_mensal'.")
	}

	if parametros.Nome == "" {
		erros = append(erros, "Por favor, forneça um nome válido para o parâmetro 'nome'.")
	}

	if len(erros) > 0 {
		return ErroValidacaoParametro(erros)
	}
	/*** Validação ***/

	/*** Banco de Dados ***/
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
		parametros.NomeParam,
	)

	if err != nil {
		log.Printf("PlanoDeAssinaturaUpdate: %v", err)
		return err
	}
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Plano de assinatura atualizado com sucesso",
	})
}

// PlanoDeAssinaturaRemove godoc
//
// @Summary Remove um plano de assinatura
// @Tags    Plano de Assinatura
// @Accept  json
// @Produce json
// @Param   nome path     string true "Nome"
// @Success 200  {object} map[string]string
// @Failure 400  {object} Erro
// @Failure 500  {object} Erro
// @Router  /api/plano_de_assinatura/:nome [delete]
func (s *Server) PlanoDeAssinaturaRemove(c echo.Context) error {
	/*** Parâmetros ***/
	parametros := struct {
		Nome string `param:"nome"`
	}{}
	/*** Parâmetros ***/

	/*** Validação ***/
	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome do plano de assinatura no parâmetro 'nome'.")
	}

	if len(erros) > 0 {
		return ErroValidacaoParametro(erros)
	}
	/*** Validação ***/

	/*** Banco de Dados ***/
	_, err := s.db.Exec(
		"UPDATE PLANO_DE_ASSINATURA SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE NOME = $1",
		parametros.Nome,
	)

	if err != nil {
		log.Printf("PlanoDeAssinaturaRemove: %v", err)
		return err
	}
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Plano de assinatura removido com sucesso",
	})
}
