package server

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirect-max/internal/model"
	"github.com/TheDevOpsCorp/redirect-max/internal/util"
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
// @Failure 400  {object} util.Erro
// @Failure 500  {object} util.Erro
// @Router  /api/plano_de_assinatura/:nome [get]
func (s *Server) PlanoDeAssinaturaReadByNome(c echo.Context) error {
	nome := c.Param("nome")

	if err := util.Validate.Var(nome, "required,min=3,max=120"); err != nil {
		return util.ErroValidacaoNome
	}

	planoDeAssinatura, err := model.PlanoDeAssinaturaReadByNome(s.db, nome)

	if err != nil {
		log.Printf("PlanoDeAssinaturaReadByNome: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, planoDeAssinatura)
}

// PlanoDeAssinaturaReadAll godoc
//
// @Summary Retorna os planos de assinatura
// @Tags    Plano de Assinatura
// @Accept  json
// @Produce json
// @Success 200  {object} []model.PlanoDeAssinatura
// @Failure 400  {object} util.Erro
// @Failure 500  {object} util.Erro
// @Router  /api/plano_de_assinatura [get]
func (s *Server) PlanoDeAssinaturaReadAll(c echo.Context) error {
	planosDeAssinatura, err := model.PlanoDeAssinaturaReadAll(s.db)

	if err != nil {
		log.Printf("PlanoDeAssinaturaReadAll: %v", err)
		return util.ErroBancoDados
	}

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
// @Failure 400          {object} util.Erro
// @Failure 500          {object} util.Erro
// @Router  /api/plano_de_assinatura [post]
func (s *Server) PlanoDeAssinaturaCreate(c echo.Context) error {
	parametros := struct {
		Nome        string `json:"nome"`
		ValorMensal int64  `json:"valor_mensal"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome e o valor mensal do plano de assinatura nos parâmetros 'nome' e 'valor_mensal', respectivamente.")
	}

	if err := util.Validate.Var(parametros.Nome, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido no parâmetro 'nome'.")
	}

	if err := util.Validate.Var(parametros.ValorMensal, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um valor mensal válido no parâmetro 'valor_mensal'.")
	}

	if len(erros) > 0 {
		return util.ErroValidacaoParametro(erros)
	}

	err := model.PlanoDeAssinaturaCreate(s.db, parametros.Nome, parametros.ValorMensal)

	if err != nil {
		log.Printf("PlanoDeAssinaturaCreate: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O plano de assinatura foi adicionado com sucesso.",
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
// @Failure 400          {object} util.Erro
// @Failure 500          {object} util.Erro
// @Router  /api/plano_de_assinatura/:nome [patch]
func (s *Server) PlanoDeAssinaturaUpdate(c echo.Context) error {
	type parametrosUpdate struct {
		Nome        string `json:"nome"`
		ValorMensal int64  `json:"valor_mensal"`
	}

	parametros := parametrosUpdate{}

	nome := c.Param("nome")

	if err := util.Validate.Var(nome, "required,min=3,max=120"); err != nil {
		return util.ErroValidacaoNome
	}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome do plano de assinatura no parâmetro 'nome' e o valor mensal no parâmetro 'valor_mensal'.")
	}

	if err := util.Validate.Var(parametros.Nome, "min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido para o parâmetro 'nome'.")
	}

	if err := util.Validate.Var(parametros.ValorMensal, "gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um valor mensal válido no parâmetro 'valor_mensal'.")
	}

	if (parametrosUpdate{}) == parametros {
		erros = append(erros, "Por favor, forneça algum valor válido para a atualização.")
	}

	if len(erros) > 0 {
		return util.ErroValidacaoParametro(erros)
	}

	err := model.PlanoDeAssinaturaUpdate(s.db, nome, parametros.Nome, parametros.ValorMensal)

	if err != nil {
		log.Printf("PlanoDeAssinaturaUpdate: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O plano de assinatura foi atualizado com sucesso.",
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
// @Failure 400  {object} util.Erro
// @Failure 500  {object} util.Erro
// @Router  /api/plano_de_assinatura/:nome [delete]
func (s *Server) PlanoDeAssinaturaRemove(c echo.Context) error {
	nome := c.Param("nome")

	if err := util.Validate.Var(nome, "required,min=3,max=120"); err != nil {
		return util.ErroValidacaoNome
	}

	err := model.PlanoDeAssinaturaRemove(s.db, nome)

	if err != nil {
		log.Printf("PlanoDeAssinaturaRemove: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O plano de assinatura foi removido com sucesso.",
	})
}
