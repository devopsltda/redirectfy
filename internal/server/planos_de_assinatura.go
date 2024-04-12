package server

import (
	"log/slog"
	"net/http"

	"redirectfy/internal/utils"

	"github.com/labstack/echo/v4"

	_ "redirectfy/internal/models"
)

// PlanoDeAssinaturaReadByNome godoc
//
// @Summary Retorna o plano de assinatura com o nome fornecido
//
// @Tags    Planos de Assinatura
//
// @Accept  json
//
// @Produce json
//
// @Param   name path     string  true  "Nome"
//
// @Success 200  {object} models.PlanoDeAssinatura
//
// @Failure 400  {object} echo.HTTPError
//
// @Failure 500  {object} echo.HTTPError
//
// @Router  /pricing/:name [get]
func (s *Server) PlanoDeAssinaturaReadByNome(c echo.Context) error {
	nome := c.Param("name")

	if err := utils.Validate.Var(nome, "required,min=3,max=120"); err != nil {
		return utils.ErroValidacaoNome
	}

	planoDeAssinatura, err := s.PlanoDeAssinaturaModel.ReadByNome(nome)

	if err != nil {
		slog.Error("PlanoDeAssinaturaReadByNome", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, planoDeAssinatura)
}

// PlanoDeAssinaturaReadAll godoc
//
// @Summary Retorna os planos de assinatura
//
// @Tags    Planos de Assinatura
//
// @Accept  json
//
// @Produce json
//
// @Success 200  {object} []models.PlanoDeAssinatura
//
// @Failure 400  {object} echo.HTTPError
//
// @Failure 500  {object} echo.HTTPError
//
// @Router  /pricing [get]
func (s *Server) PlanoDeAssinaturaReadAll(c echo.Context) error {
	planosDeAssinatura, err := s.PlanoDeAssinaturaModel.ReadAll()

	if err != nil {
		slog.Error("PlanoDeAssinaturaReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, planosDeAssinatura)
}

// PlanoDeAssinaturaCreate godoc
//
// @Summary Cria um plano de assinatura
//
// @Tags    Planos de Assinatura
//
// @Accept  json
//
// @Produce json
//
// @Param   nome                            body     string true "Nome"
//
// @Param   valor_mensal                    body     int    true "Valor Mensal"
//
// @Param   limite_links_mensal             body     int    true "Limite de Links Mensal"
//
// @Param   ordenacao_aleatoria_links       body     bool   true "Ordenação Aleatória de Links"
//
// @Success 200                             {object} map[string]string
//
// @Failure 400                             {object} echo.HTTPError
//
// @Failure 500                             {object} echo.HTTPError
//
// @Router  /pricing [post]
func (s *Server) PlanoDeAssinaturaCreate(c echo.Context) error {
	parametros := struct {
		Nome                    string `json:"nome"`
		ValorMensal             int64  `json:"valor_mensal"`
		LimiteLinksMensal       int64  `json:"limite_links_mensal"`
		OrdenacaoAleatoriaLinks bool   `json:"ordenacao_aleatoria_links"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome, valor mensal, limite de links mensal e ordenação aleatória dos links do plano de assinatura nos parâmetros 'nome', 'valor_mensal', 'limite_links_mensal' e 'ordenacao_aleatoria_links', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Nome, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido no parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.ValorMensal, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um valor mensal válido no parâmetro 'valor_mensal'.")
	}

	if err := utils.Validate.Var(parametros.LimiteLinksMensal, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um limite válido no parâmetro 'limite_links_mensal'.")
	}

	if err := utils.Validate.Var(parametros.OrdenacaoAleatoriaLinks, "required"); err != nil {
		erros = append(erros, "Por favor, forneça um valor válido no parâmetro 'ordenacao_aleatoria_links'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := s.PlanoDeAssinaturaModel.Create(parametros.Nome, parametros.ValorMensal, parametros.LimiteLinksMensal, parametros.OrdenacaoAleatoriaLinks)

	if err != nil {
		slog.Error("PlanoDeAssinaturaCreate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemPlanoDeAssinaturaCriadoComSucesso)
}

// PlanoDeAssinaturaUpdate godoc
//
// @Summary Atualiza um plano de assinatura
//
// @Tags    Planos de Assinatura
//
// @Accept  json
//
// @Produce json
//
// @Param   name                            path     string true  "Nome"
//
// @Param   nome                            body     string false "Nome"
//
// @Param   valor_mensal                    body     int    false "Valor Mensal"
//
// @Param   limite_links_mensal             body     int    false "Limite de Links Mensal"
//
// @Param   ordenacao_aleatoria_links       body     bool   true  "Ordenação Aleatória de Links"
//
// @Success 200                             {object} map[string]string
//
// @Failure 400                             {object} echo.HTTPError
//
// @Failure 500                             {object} echo.HTTPError
//
// @Router  /pricing/:name [patch]
func (s *Server) PlanoDeAssinaturaUpdate(c echo.Context) error {
	type parametrosUpdate struct {
		Nome                    string `json:"nome"`
		ValorMensal             int64  `json:"valor_mensal"`
		LimiteLinksMensal       int64  `json:"limite_links_mensal"`
		OrdenacaoAleatoriaLinks bool   `json:"ordenacao_aleatoria_links"`
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

	if err := utils.Validate.Var(parametros.LimiteLinksMensal, "gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um limite válido no parâmetro 'limite_links_mensal'.")
	}

	if (parametrosUpdate{}) == parametros {
		erros = append(erros, "Por favor, forneça algum valor válido para a atualização.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := s.PlanoDeAssinaturaModel.Update(nome, parametros.Nome, parametros.ValorMensal, parametros.LimiteLinksMensal, parametros.OrdenacaoAleatoriaLinks)

	if err != nil {
		slog.Error("PlanoDeAssinaturaUpdate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemPlanoDeAssinaturaAtualizadoComSucesso)
}

// PlanoDeAssinaturaRemove godoc
//
// @Summary Remove um plano de assinatura
//
// @Tags    Planos de Assinatura
//
// @Accept  json
//
// @Produce json
//
// @Param   name path     string true "Nome"
//
// @Success 200  {object} map[string]string
//
// @Failure 400  {object} echo.HTTPError
//
// @Failure 500  {object} echo.HTTPError
//
// @Router  /pricing/:name [delete]
func (s *Server) PlanoDeAssinaturaRemove(c echo.Context) error {
	nome := c.Param("nome")

	if err := utils.Validate.Var(nome, "required,min=3,max=120"); err != nil {
		return utils.ErroValidacaoNome
	}

	err := s.PlanoDeAssinaturaModel.Remove(nome)

	if err != nil {
		slog.Error("PlanoDeAssinaturaRemove", slog.Any("error", err))
		return err
	}

	return c.JSON(http.StatusOK, utils.MensagemPlanoDeAssinaturaRemovidoComSucesso)
}
