package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"redirectfy/internal/utils"

	_ "redirectfy/internal/models"
)

func (s *Server) PricingHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.PlanoDeAssinaturaReadAll()

	if err != nil {
		utils.ErroLog("PricingHistory", "Erro na leitura do histórico de planos de assinatura", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos planos de assinatura.")
	}

	return c.JSON(http.StatusOK, historico)
}

func (s *Server) KirvanoHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.UsuarioTemporarioReadAll()

	if err != nil {
		utils.ErroLog("KirvanoHistory", "Erro na leitura do histórico de usuários da Kirvano", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários da Kirvano.")
	}

	return c.JSON(http.StatusOK, historico)
}

func (s *Server) UserHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.UsuarioReadAll()

	if err != nil {
		utils.ErroLog("UserHistory", "Erro na leitura do histórico de usuários", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários.")
	}

	return c.JSON(http.StatusOK, historico)
}

func (s *Server) RedirectHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.RedirecionadorReadAll()

	if err != nil {
		utils.ErroLog("RedirectHistory", "Erro na leitura do histórico de redirecionadores", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos usuários.")
	}

	return c.JSON(http.StatusOK, historico)
}

func (s *Server) LinkHistory(c echo.Context) error {
	historico, err := s.HistoricoModel.LinkReadAll()

	if err != nil {
		utils.ErroLog("LinkHistory", "Erro na leitura do histórico de links", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o histórico dos links.")
	}

	return c.JSON(http.StatusOK, historico)
}
