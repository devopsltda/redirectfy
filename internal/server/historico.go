package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"redirectfy/internal/utils"

	_ "redirectfy/internal/models"
)

// PricingHistory godoc
//
// @Summary Retorna o histórico de ações relativas a planos de assinatura no sistema
//
// @Tags    Admin
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoPlanoDeAssinatura
//
// @Failure 400 {object} echo.HTTPError
//
// @Failure 500 {object} echo.HTTPError
//
// @Router  /admin/pricing_history [get]
func (s *Server) PricingHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.PlanoDeAssinaturaReadAll()

	if err != nil {
		utils.ErroLog("PricingHistory", "Erro na leitura do histórico de planos de assinatura", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos planos de assinatura.")
	}

	return c.JSON(http.StatusOK, historico)
}

// KirvanoHistory godoc
//
// @Summary Retorna o histórico de ações relativas a usuários da Kirvano no sistema
//
// @Tags    Admin
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoUsuarioTemporario
//
// @Failure 400 {object} echo.HTTPError
//
// @Failure 500 {object} echo.HTTPError
//
// @Router  /admin/kirvano_history [get]
func (s *Server) KirvanoHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.UsuarioTemporarioReadAll()

	if err != nil {
		utils.ErroLog("KirvanoHistory", "Erro na leitura do histórico de usuários da Kirvano", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários da Kirvano.")
	}

	return c.JSON(http.StatusOK, historico)
}

// UserHistory godoc
//
// @Summary Retorna o histórico de ações relativas a usuários no sistema
//
// @Tags    Admin
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoUsuario
//
// @Failure 400 {object} echo.HTTPError
//
// @Failure 500 {object} echo.HTTPError
//
// @Router  /admin/user_history [get]
func (s *Server) UserHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.UsuarioReadAll()

	if err != nil {
		utils.ErroLog("UserHistory", "Erro na leitura do histórico de usuários", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários.")
	}

	return c.JSON(http.StatusOK, historico)
}

// RedirectHistory godoc
//
// @Summary Retorna o histórico de ações relativas a redirecionadores no sistema
//
// @Tags    Admin
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoRedirecionador
//
// @Failure 400 {object} echo.HTTPError
//
// @Failure 500 {object} echo.HTTPError
//
// @Router  /admin/redirect_history [get]
func (s *Server) RedirectHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.RedirecionadorReadAll()

	if err != nil {
		utils.ErroLog("RedirectHistory", "Erro na leitura do histórico de redirecionadores", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários.")
	}

	return c.JSON(http.StatusOK, historico)
}

// LinkHistory godoc
//
// @Summary Retorna o histórico de ações relativas a links no sistema
//
// @Tags    Admin
//
// @Accept  json
//
// @Produce json
//
// @Success 200 {object} []models.HistoricoLink
//
// @Failure 400 {object} echo.HTTPError
//
// @Failure 500 {object} echo.HTTPError
//
// @Router  /admin/link_history [get]
func (s *Server) LinkHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.LinkReadAll()

	if err != nil {
		utils.ErroLog("LinkHistory", "Erro na leitura do histórico de links", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos links.")
	}

	return c.JSON(http.StatusOK, historico)
}
