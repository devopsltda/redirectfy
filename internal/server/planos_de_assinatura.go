package server

import (
	"database/sql"
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
		utils.DebugLog("PlanoDeAssinaturaReadByNome", "Erro na validação do nome do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'name' inserido é inválido, por favor insira um 'name' existente entre 3 e 120 caracteres.")
	}

	planoDeAssinatura, err := s.PlanoDeAssinaturaModel.ReadByNome(nome)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.DebugLog("PlanoDeAssinaturaReadByNome", "Não há nenhum plano de assinatura com esse nome", err)
			return utils.Erro(http.StatusNotFound, "O plano de assinatura com o nome inserido não foi encontrado.")
		}

		utils.ErroLog("PlanoDeAssinaturaReadByNome", "Erro na leitura do plano de assinatura com o nome inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o plano de assinatura com o nome inserido.")
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
		utils.ErroLog("PlanoDeAssinaturaReadAll", "Erro na leitura dos planos de assinatura", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler os planos de assinatura.")
	}

	return c.JSON(http.StatusOK, planosDeAssinatura)
}
