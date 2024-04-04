package server

import (
	"log/slog"
	"net/http"
	"strconv"

	"redirectify/internal/utils"

	"github.com/labstack/echo/v4"
)

// LinkReadById godoc
//
// @Summary Retorna o link com o id fornecido do redirecionador com o código hash fornecido
// @Tags    Links
// @Accept  json
// @Produce json
// @Param   codigo_hash path     string true "Código Hash"
// @Param   id          path     int true "Id"
// @Success 200         {object} models.Link
// @Failure 400         {object} utils.Erro
// @Failure 500         {object} utils.Erro
// @Router  /v1/api/redirecionadores/:codigo_hash/links/:id [get]
func (s *Server) LinkReadById(c echo.Context) error {
	id := c.Param("id")
	codigoHash := c.Param("codigo_hash")

	if err := utils.Validate.Var(id, "required,gte=0"); err != nil {
		return utils.ErroValidacaoCodigoHash
	}

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return utils.ErroValidacaoCodigoHash
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		slog.Error("LinkReadById", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	link, err := s.LinkModel.ReadById(parsedId, codigoHash)

	if err != nil {
		slog.Error("LinkReadById", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, link)
}

// LinkReadAll godoc
//
// @Summary Retorna os links do redirecionador com o código hash fornecido
// @Tags    Links
// @Accept  json
// @Produce json
// @Param   codigo_hash path     string true "Código Hash"
// @Success 200         {object} []models.Link
// @Failure 400         {object} utils.Erro
// @Failure 500         {object} utils.Erro
// @Router  /v1/api/redirecionadores/:codigo_hash/links [get]
func (s *Server) LinkReadAll(c echo.Context) error {
	codigoHash := c.Param("codigo_hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return utils.ErroValidacaoCodigoHash
	}

	links, err := s.LinkModel.ReadAll(codigoHash)

	if err != nil {
		slog.Error("LinkReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, links)
}

// LinkCreate godoc
//
// @Summary Cria um link no redirecionador com o código hash fornecido
// @Tags    Links
// @Accept  json
// @Produce json
// @Param   codigo_hash         path     string true "Código Hash"
// @Param   nome                body     string true "Nome"
// @Param   link                body     string true "Link"
// @Param   plataforma          body     string true "Plataforma"
// @Success 200                 {object} map[string]string
// @Failure 400                 {object} utils.Erro
// @Failure 500                 {object} utils.Erro
// @Router  /v1/api/redirecionadores/:codigo_hash/links [post]
func (s *Server) LinkCreate(c echo.Context) error {
	parametros := struct {
		CodigoHash string `path:"codigo_hash"`
		Nome       string `json:"nome"`
		Link       string `json:"link"`
		Plataforma string `json:"plataforma"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o código hash, link e plataforma nos parâmetro 'codigo_hash', 'link' e 'plataforma', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.CodigoHash, "required,len=10"); err != nil {
		erros = append(erros, "Por favor, forneça um código hash válido para o parâmetro 'codigo_hash'.")
	}

	if err := utils.Validate.Var(parametros.Nome, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.Link, "required"); err != nil {
		erros = append(erros, "Por favor, forneça um link válido para o parâmetro 'link'.")
	}

	if err := utils.Validate.Var(parametros.Plataforma, "required,oneof=whatsapp telegram"); err != nil {
		erros = append(erros, "Por favor, forneça uma plataforma válida para o parâmetro 'plataforma'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := s.LinkModel.Create(
		parametros.Nome,
		parametros.Link,
		parametros.Plataforma,
		parametros.CodigoHash,
	)

	if err != nil {
		slog.Error("LinkCreate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusCreated, utils.MensagemLinkAtualizadoComSucesso)
}

// LinkUpdate godoc
//
// @Summary Atualiza um link específico de um redirecionador específico
// @Tags    Links
// @Accept  json
// @Produce json
// @Param   codigo_hash path     string true  "Código Hash"
// @Param   id          path     int    true  "Id"
// @Param   nome        body     string false "Nome"
// @Param   link        body     string false "Link"
// @Param   plataforma  body     string false "Plataforma"
// @Success 200         {object} map[string]string
// @Failure 400         {object} utils.Erro
// @Failure 500         {object} utils.Erro
// @Router  /v1/api/redirecionadores/:codigo_hash/link/:id [patch]
func (s *Server) LinkUpdate(c echo.Context) error {
	parametros := struct {
		CodigoHash string `path:"codigo_hash"`
		Id         int64  `path:"id"`
		Nome       string `json:"nome"`
		Link       string `json:"link"`
		Plataforma string `json:"plataforma"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o código hash, id, link e plataforma nos parâmetro 'codigo_hash', 'id', 'link' e 'plataforma', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.CodigoHash, "required,len=10"); err != nil {
		erros = append(erros, "Por favor, forneça um código hash válido para o parâmetro 'codigo_hash'.")
	}

	if err := utils.Validate.Var(parametros.Id, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um id válido para o parâmetro 'id'.")
	}

	if err := utils.Validate.Var(parametros.Nome, "min=3,max=120"); parametros.Nome != "" && err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.Plataforma, "oneof=whatsapp telegram"); parametros.Plataforma != "" && err != nil {
		erros = append(erros, "Por favor, forneça uma plataforma válida para o parâmetro 'plataforma'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := s.LinkModel.Update(parametros.Id, parametros.CodigoHash, parametros.Nome, parametros.Link, parametros.Plataforma)

	if err != nil {
		slog.Error("LinkUpdate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemLinkAtualizadoComSucesso)
}

// LinkRemove godoc
//
// @Summary Remove um link específico de um redirecionador específico
// @Tags    Links
// @Accept  json
// @Produce json
// @Param   codigo_hash path     string true "Código Hash"
// @Param   id          path     int    true "Id"
// @Success 200         {object} map[string]string
// @Failure 400         {object} utils.Erro
// @Failure 500         {object} utils.Erro
// @Router  /v1/api/redirecionadores/:codigo_hash/link/:id [delete]
func (s *Server) LinkRemove(c echo.Context) error {
	id := c.Param("id")
	codigoHash := c.Param("codigo_hash")

	if err := utils.Validate.Var(id, "required,gte=0"); err != nil {
		return utils.ErroValidacaoCodigoHash
	}

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return utils.ErroValidacaoCodigoHash
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		slog.Error("LinkRemove", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.LinkModel.Remove(parsedId, codigoHash)

	if err != nil {
		slog.Error("LinkRemove", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemLinkRemovidoComSucesso)
}
