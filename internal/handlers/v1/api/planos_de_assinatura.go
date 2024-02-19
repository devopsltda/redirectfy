package api

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirectify/internal/database"
	"github.com/TheDevOpsCorp/redirectify/internal/models"
	"github.com/TheDevOpsCorp/redirectify/internal/utils"
	"github.com/labstack/echo/v4"
)

// PlanoDeAssinaturaReadByNome godoc
//
// @Summary Retorna o plano de assinatura com o nome fornecido
// @Tags    Planos de Assinatura
// @Accept  json
// @Produce json
// @Param   nome path     string  true  "Nome"
// @Success 200  {object} models.PlanoDeAssinatura
// @Failure 400  {object} utils.Erro
// @Failure 500  {object} utils.Erro
// @Router  /v1/api/planos_de_assinatura/:nome [get]
func PlanoDeAssinaturaReadByNome(c echo.Context) error {
	nome := c.Param("nome")

	if err := utils.Validate.Var(nome, "required,min=3,max=120"); err != nil {
		return utils.ErroValidacaoNome
	}

	planoDeAssinatura, err := models.PlanoDeAssinaturaReadByNome(database.Db, nome)

	if err != nil {
		log.Printf("PlanoDeAssinaturaReadByNome: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, planoDeAssinatura)
}

// PlanoDeAssinaturaReadAll godoc
//
// @Summary Retorna os planos de assinatura
// @Tags    Planos de Assinatura
// @Accept  json
// @Produce json
// @Success 200  {object} []models.PlanoDeAssinatura
// @Failure 400  {object} utils.Erro
// @Failure 500  {object} utils.Erro
// @Router  /v1/api/planos_de_assinatura [get]
func PlanoDeAssinaturaReadAll(c echo.Context) error {
	planosDeAssinatura, err := models.PlanoDeAssinaturaReadAll(database.Db)

	if err != nil {
		log.Printf("PlanoDeAssinaturaReadAll: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, planosDeAssinatura)
}

// PlanoDeAssinaturaCreate godoc
//
// @Summary Cria um plano de assinatura
// @Tags    Planos de Assinatura
// @Accept  json
// @Produce json
// @Param   nome         body     string true "Nome"
// @Param   valor_mensal body     int    true "Valor Mensal"
// @Success 200          {object} map[string]string
// @Failure 400          {object} utils.Erro
// @Failure 500          {object} utils.Erro
// @Router  /v1/api/planos_de_assinatura [post]
func PlanoDeAssinaturaCreate(c echo.Context) error {
	parametros := struct {
		Nome          string `json:"nome"`
		ValorMensal   int64  `json:"valor_mensal"`
		Limite        int64  `json:"limite"`
		PeriodoLimite string `json:"periodo_limite"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome e o valor mensal do plano de assinatura nos parâmetros 'nome' e 'valor_mensal', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Nome, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido no parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.ValorMensal, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um valor mensal válido no parâmetro 'valor_mensal'.")
	}

	if err := utils.Validate.Var(parametros.Limite, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um limite válido no parâmetro 'limite'.")
	}

	if err := utils.Validate.Var(parametros.PeriodoLimite, "required,oneof=s m h d M"); err != nil {
		erros = append(erros, "Por favor, forneça um período válido para o limite no parâmetro 'periodo_limite'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := models.PlanoDeAssinaturaCreate(database.Db, parametros.Nome, parametros.ValorMensal, parametros.Limite, parametros.PeriodoLimite)

	if err != nil {
		log.Printf("PlanoDeAssinaturaCreate: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemPlanoDeAssinaturaCriadoComSucesso)
}

// PlanoDeAssinaturaUpdate godoc
//
// @Summary Atualiza um plano de assinatura
// @Tags    Planos de Assinatura
// @Accept  json
// @Produce json
// @Param   nome         path     string true  "Nome"
// @Param   nome         body     string false "Nome"
// @Param   valor_mensal body     int    false "Valor Mensal"
// @Success 200          {object} map[string]string
// @Failure 400          {object} utils.Erro
// @Failure 500          {object} utils.Erro
// @Router  /v1/api/planos_de_assinatura/:nome [patch]
func PlanoDeAssinaturaUpdate(c echo.Context) error {
	type parametrosUpdate struct {
		Nome          string `json:"nome"`
		ValorMensal   int64  `json:"valor_mensal"`
		Limite        int64  `json:"limite"`
		PeriodoLimite string `json:"periodo_limite"`
	}

	parametros := parametrosUpdate{}

	nome := c.Param("nome")

	if err := utils.Validate.Var(nome, "required,min=3,max=120"); err != nil {
		return utils.ErroValidacaoNome
	}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome do plano de assinatura no parâmetro 'nome' e o valor mensal no parâmetro 'valor_mensal'.")
	}

	if err := utils.Validate.Var(parametros.Nome, "min=3,max=120"); parametros.Nome != "" && err != nil {
		erros = append(erros, "Por favor, forneça um nome válido para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.ValorMensal, "gte=0"); parametros.ValorMensal != 0 && err != nil {
		erros = append(erros, "Por favor, forneça um valor mensal válido no parâmetro 'valor_mensal'.")
	}

	if err := utils.Validate.Var(parametros.Limite, "gte=0"); parametros.Limite != 0 && err != nil {
		erros = append(erros, "Por favor, forneça um limite válido no parâmetro 'limite'.")
	}

	if err := utils.Validate.Var(parametros.PeriodoLimite, "oneof=s m h d M"); parametros.PeriodoLimite != "" && err != nil {
		erros = append(erros, "Por favor, forneça um período válido para o limite no parâmetro 'periodo_limite'.")
	}

	if (parametrosUpdate{}) == parametros {
		erros = append(erros, "Por favor, forneça algum valor válido para a atualização.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := models.PlanoDeAssinaturaUpdate(database.Db, nome, parametros.Nome, parametros.ValorMensal, parametros.Limite, parametros.PeriodoLimite)

	if err != nil {
		log.Printf("PlanoDeAssinaturaUpdate: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemPlanoDeAssinaturaAtualizadoComSucesso)
}

// PlanoDeAssinaturaRemove godoc
//
// @Summary Remove um plano de assinatura
// @Tags    Planos de Assinatura
// @Accept  json
// @Produce json
// @Param   nome path     string true "Nome"
// @Success 200  {object} map[string]string
// @Failure 400  {object} utils.Erro
// @Failure 500  {object} utils.Erro
// @Router  /v1/api/planos_de_assinatura/:nome [delete]
func PlanoDeAssinaturaRemove(c echo.Context) error {
	nome := c.Param("nome")

	if err := utils.Validate.Var(nome, "required,min=3,max=120"); err != nil {
		return utils.ErroValidacaoNome
	}

	err := models.PlanoDeAssinaturaRemove(database.Db, nome)

	if err != nil {
		log.Printf("PlanoDeAssinaturaRemove: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, utils.MensagemPlanoDeAssinaturaRemovidoComSucesso)
}
