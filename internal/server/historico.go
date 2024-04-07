package server

import (
	"log/slog"
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
// @Failure 400 {object} Erro
//
// @Failure 500 {object} Erro
//
// @Router  /planos_de_assinatura/historico [get]
func (s *Server) HistoricoPlanoDeAssinaturaReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.PlanoDeAssinaturaReadAll()

	if err != nil {
		slog.Error("HistoricoPlanoDeAssinaturaReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
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
// @Failure 400 {object} Erro
//
// @Failure 500 {object} Erro
//
// @Router  /usuarios_temporarios/historico [get]
func (s *Server) HistoricoUsuarioTemporarioReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.UsuarioTemporarioReadAll()

	if err != nil {
		slog.Error("HistoricoUsuarioTemporarioReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
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
// @Failure 400 {object} Erro
//
// @Failure 500 {object} Erro
//
// @Router  /usuarios/historico [get]
func (s *Server) HistoricoUsuarioReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.UsuarioReadAll()

	if err != nil {
		slog.Error("HistoricoUsuarioReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
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
// @Failure 400 {object} Erro
//
// @Failure 500 {object} Erro
//
// @Router  /redirecionadores/historico [get]
func (s *Server) HistoricoRedirecionadorReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.RedirecionadorReadAll()

	if err != nil {
		slog.Error("HistoricoRedirecionadorReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
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
// @Failure 400 {object} Erro
//
// @Failure 500 {object} Erro
//
// @Router  /links/historico [get]
func (s *Server) HistoricoLinkReadAll(c echo.Context) error {
	historico, err := s.HistoricoModel.LinkReadAll()

	if err != nil {
		slog.Error("HistoricoLinkReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, historico)
}
