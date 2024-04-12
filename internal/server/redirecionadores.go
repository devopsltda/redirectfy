package server

import (
	"log/slog"
	"net/http"

	"redirectfy/internal/models"
	"redirectfy/internal/utils"

	"github.com/labstack/echo/v4"

	_ "redirectfy/internal/models"
)

type RedirecionadorReadByCodigoHashResponse struct {
	R models.Redirecionador `json:"redirecionador"`
	L []models.Link         `json:"links"`
}

// RedirecionadorReadByCodigoHash godoc
//
// @Summary Retorna o redirecionador com o código hash fornecido
//
// @Tags    Redirecionadors
//
// @Accept  json
//
// @Produce json
//
// @Param   codigo_hash path     string true "Código Hash"
//
// @Success 200         {object} models.Redirecionador
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /redirecionadores/:codigo_hash [get]
func (s *Server) RedirecionadorReadByCodigoHash(c echo.Context) error {
	codigoHash := c.Param("codigo_hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return utils.ErroValidacaoCodigoHash
	}

	redirecionador, err := s.RedirecionadorModel.ReadByCodigoHash(codigoHash)

	if err != nil {
		slog.Error("RedirecionadorReadByCodigoHash", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	links, err := s.LinkModel.ReadByCodigoHash(codigoHash)

	if err != nil {
		slog.Error("LinkReadByCodigoHash", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	picked_links := models.LinkPicker(links)

	return c.JSON(http.StatusOK, RedirecionadorReadByCodigoHashResponse{R: redirecionador, L: picked_links})
}

// RedirecionadorReadAll godoc
//
// @Summary Retorna os redirecionadores
//
// @Tags    Redirecionadores
//
// @Accept  json
//
// @Produce json
//
// @Param   nome_de_usuario path     string true "Nome de Usuário"
//
// @Success 200             {object} []models.Redirecionador
//
// @Failure 400             {object} echo.HTTPError
//
// @Failure 500             {object} echo.HTTPError
//
// @Router  /redirecionadores/:nome_de_usuario [get]
func (s *Server) RedirecionadorReadAll(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	redirecionadores, err := s.RedirecionadorModel.ReadAll(nomeDeUsuario)

	if err != nil {
		slog.Error("RedirecionadorReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, redirecionadores)
}

// RedirecionadorCreate godoc
//
// @Summary Cria um redirecionador
//
// @Tags    Redirecionadores
//
// @Accept  json
//
// @Produce json
//
// @Param   nome                      body     string true  "Nome"
//
// @Param   ordem_de_redirecionamento body     string true  "Ordem de Redirecionamento"
//
// @Param   nome_de_usuario           body     string true  "Nome de Usuário"
//
// @Success 200                       {object} map[string]string
//
// @Failure 400                       {object} echo.HTTPError
//
// @Failure 500                       {object} echo.HTTPError
//
// @Router  /redirecionadores [post]
func (s *Server) RedirecionadorCreate(c echo.Context) error {
	parametros := struct {
		Nome                    string `json:"nome"`
		OrdemDeRedirecionamento string `json:"ordem_de_redirecionamento"`
		Usuario                 string `json:"nome_de_usuario"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome, redirecionador do whatsapp, redirecionador do telegram, ordem de redirecionamento e usuário nos parâmetro 'nome', 'redirecionador_whatsapp', 'redirecionador_telegram', 'ordem_de_redirecionamento' e 'usuario', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Nome, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.OrdemDeRedirecionamento, "required,max=120,oneof=telegram0x2Cwhatsapp whatsapp0x2Ctelegram"); err != nil {
		erros = append(erros, "Por favor, forneça uma ordem de redirecionamento válida para o parâmetro 'ordem_de_redirecionamento'.")
	}

	if !utils.ValidaNomeDeUsuario(parametros.Usuario) {
		erros = append(erros, "Por favor, forneça um nome de usuário válido para o parâmetro 'nome_de_usuario'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	var err error
	var codigoHash string
	codigoHashExiste := true

	for codigoHashExiste {
		codigoHash = utils.GeraHashCode(10)

		codigoHashExiste, err = s.RedirecionadorModel.CheckIfCodigoHashExists(codigoHash)

		if err != nil {
			slog.Error("RedirecionadorCreate", slog.Any("error", err))
			return utils.ErroBancoDados
		}
	}

	_, err = s.RedirecionadorModel.Create(
		parametros.Nome,
		codigoHash,
		parametros.OrdemDeRedirecionamento,
		parametros.Usuario,
	)

	if err != nil {
		slog.Error("RedirecionadorCreate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusCreated, codigoHash)
}

// RedirecionadorRehash godoc
//
// @Summary Recria o hash de um redirecionador
//
// @Tags    Redirecionadores
//
// @Accept  json
//
// @Produce json
//
// @Param   codigo_hash path     string true "Código Hash"
//
// @Success 200         {object} map[string]string
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /redirecionadores/rehash/:codigo_hash [patch]
func (s *Server) RedirecionadorRehash(c echo.Context) error {
	codigoHash := c.Param("codigo_hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return utils.ErroValidacaoCodigoHash
	}

	var err error
	var codigoHashNovo string
	codigoHashExiste := true

	for codigoHashExiste {
		codigoHashNovo = utils.GeraHashCode(10)

		codigoHashExiste, err = s.RedirecionadorModel.CheckIfCodigoHashExists(codigoHashNovo)

		if err != nil {
			slog.Error("RedirecionadorRehash", slog.Any("error", err))
			return utils.ErroBancoDados
		}
	}

	err = s.RedirecionadorModel.Rehash(codigoHash, codigoHashNovo)

	if err != nil {
		slog.Error("RedirecionadorRehash", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, codigoHashNovo)
}

// RedirecionadorUpdate godoc
//
// @Summary Atualiza um redirecionador
//
// @Tags    Redirecionadores
//
// @Accept  json
//
// @Produce json
//
// @Param   codigo_hash               path     string true  "Código Hash"
//
// @Param   nome                      body     string false "Nome"
//
// @Param   ordem_de_redirecionamento body     string false "Ordem de Redirecionamento"
//
// @Success 200                       {object} map[string]string
//
// @Failure 400                       {object} echo.HTTPError
//
// @Failure 500                       {object} echo.HTTPError
//
// @Router  /redirecionadores/:codigo_hash [patch]
func (s *Server) RedirecionadorUpdate(c echo.Context) error {
	parametros := struct {
		Nome                    string `json:"nome"`
		OrdemDeRedirecionamento string `json:"ordem_de_redirecionamento"`
	}{}

	var erros []string

	codigoHash := c.Param("codigo_hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		slog.Error("RedirecionadorUpdate", slog.Any("error", err))
		return utils.ErroValidacaoCodigoHash
	}

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome, redirecionador do whatsapp, redirecionador do telegram, ordem de redirecionamento e usuário nos parâmetro 'nome', 'redirecionador_whatsapp', 'redirecionador_telegram', 'ordem_de_redirecionamento' e 'usuario', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Nome, "min=3,max=120"); parametros.Nome != "" && err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.OrdemDeRedirecionamento, "max=120,oneof=telegram0x2Cwhatsapp whatsapp0x2Ctelegram"); parametros.OrdemDeRedirecionamento != "" && err != nil {
		erros = append(erros, "Por favor, forneça uma ordem de redirecionamento válida para o parâmetro 'ordem_de_redirecionamento'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := s.RedirecionadorModel.Update(parametros.Nome, codigoHash, parametros.OrdemDeRedirecionamento)

	if err != nil {
		slog.Error("RedirecionadorUpdate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemRedirecionadorAtualizadoComSucesso)
}

// RedirecionadorRemove godoc
//
// @Summary Remove um redirecionador
//
// @Tags    Redirecionadors
//
// @Accept  json
//
// @Produce json
//
// @Param   codigo_hash path     string true "Código Hash"
//
// @Success 200         {object} map[string]string
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /redirecionadores/:codigo_hash [delete]
func (s *Server) RedirecionadorRemove(c echo.Context) error {
	codigoHash := c.Param("codigo_hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return utils.ErroValidacaoCodigoHash
	}

	err := s.RedirecionadorModel.Remove(codigoHash)

	if err != nil {
		slog.Error("RedirecionadorRemove", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemRedirecionadorRemovidoComSucesso)
}
