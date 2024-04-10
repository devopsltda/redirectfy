package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"redirectfy/internal/utils"

	_ "redirectfy/internal/models"
)

// HistoricoPlanoDeAssinaturaReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a planos de assinatura no sistema
//
// @Tags    Histórico
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoPlanoDeAssinatura
//
// @Failure 400 {object} *echo.HTTPError
//
// @Failure 500 {object} *echo.HTTPError
//
// @Router  /planos_de_assinatura/historico [get]
func (s *Server) HistoricoPlanoDeAssinaturaReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.PlanoDeAssinaturaReadAll()

	if err != nil {
		utils.ErroLog("HistoricoPlanoDeAssinaturaReadAll", "Erro na leitura do histórico de planos de assinatura", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos planos de assinatura.")
	}

	return c.JSON(http.StatusOK, historico)
}

// HistoricoUsuarioTemporarioReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a usuários temporários no sistema
//
// @Tags    Histórico
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoUsuarioTemporario
//
// @Failure 400 {object} *echo.HTTPError
//
// @Failure 500 {object} *echo.HTTPError
//
// @Router  /usuarios_temporarios/historico [get]
func (s *Server) HistoricoUsuarioTemporarioReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.UsuarioTemporarioReadAll()

	if err != nil {
		utils.ErroLog("HistoricoUsuarioTemporarioReadAll", "Erro na leitura do histórico de usuários temporários", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários temporários.")
	}

	return c.JSON(http.StatusOK, historico)
}

// HistoricoUsuarioReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a usuários no sistema
//
// @Tags    Histórico
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoUsuario
//
// @Failure 400 {object} *echo.HTTPError
//
// @Failure 500 {object} *echo.HTTPError
//
// @Router  /usuarios/historico [get]
func (s *Server) HistoricoUsuarioReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.UsuarioReadAll()

	if err != nil {
		utils.ErroLog("HistoricoUsuarioReadAll", "Erro na leitura do histórico de usuários", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários.")
	}

	return c.JSON(http.StatusOK, historico)
}

// HistoricoRedirecionadorReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a redirecionadores no sistema
//
// @Tags    Histórico
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoRedirecionador
//
// @Failure 400 {object} *echo.HTTPError
//
// @Failure 500 {object} *echo.HTTPError
//
// @Router  /redirecionadores/historico [get]
func (s *Server) HistoricoRedirecionadorReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.RedirecionadorReadAll()

	if err != nil {
		utils.ErroLog("HistoricoRedirecionadorReadAll", "Erro na leitura do histórico de redirecionadores", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários.")
	}

	return c.JSON(http.StatusOK, historico)
}

// HistoricoLinkReadAll godoc
//
// @Summary Retorna o histórico de ações relativas a links no sistema
//
// @Tags    Histórico
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoLink
//
// @Failure 400 {object} *echo.HTTPError
//
// @Failure 500 {object} *echo.HTTPError
//
// @Router  /links/historico [get]
func (s *Server) HistoricoLinkReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.LinkReadAll()

	if err != nil {
		utils.ErroLog("HistoricoLinkReadAll", "Erro na leitura do histórico de links", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos links.")
	}

	return c.JSON(http.StatusOK, historico)
}
