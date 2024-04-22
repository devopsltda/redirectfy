package server

import (
	"fmt"
	"net/http"
	"strconv"

	"redirectfy/internal/models"
	"redirectfy/internal/utils"

	"github.com/labstack/echo/v4"

	_ "redirectfy/internal/models"
)

// LinkReadById godoc
//
// @Summary Retorna o link com o id fornecido do redirecionador com o código hash fornecido
//
// @Tags    Links
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true "Código Hash"
//
// @Param   id          path     int true "Id"
//
// @Success 200         {object} models.Link
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /api/r/:hash/links/:id [get]
func (s *Server) LinkReadById(c echo.Context) error {
	id := c.Param("id")
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(id, "required,gte=0"); err != nil {
		utils.DebugLog("LinkReadById", "Erro na validação do id do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'id' inserido é inválido, por favor insira um 'id' maior que 0.")
	}

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("LinkReadById", "Erro na validação do código hash do redirecionador do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		utils.ErroLog("LinkReadById", "Erro na transformação do id para inteiro", err)
		return utils.Erro(http.StatusBadRequest, "Houve um erro na transformação do id para um número inteiro. Por favor, insira um id válido.")
	}

	link, err := s.LinkModel.ReadById(parsedId, codigoHash)

	if err != nil {
		utils.ErroLog("LinkReadById", "Erro na leitura do link do redirecionador inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o link do redirecionador inserido.")
	}

	return c.JSON(http.StatusOK, link)
}

// LinkReadByCodigoHash godoc
//
// @Summary Retorna os links do redirecionador com o código hash fornecido
//
// @Tags    Links
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true "Código Hash"
//
// @Success 200         {object} []models.Link
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /api/r/:hash/links [get]
func (s *Server) LinkReadByCodigoHash(c echo.Context) error {
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("LinkReadByCodigoHash", "Erro na validação do código hash do redirecionador dos links", err)
		return utils.Erro(http.StatusBadRequest, "O 'codigo_hash' inserido é inválido, por favor insira um 'codigo_hash' existente com 10 caracteres.")
	}

	links, err := s.LinkModel.ReadByCodigoHash(codigoHash)

	if err != nil {
		utils.ErroLog("LinkReadByCodigoHash", "Erro na leitura dos links do redirecionador inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler os links do redirecionador inserido.")
	}

	return c.JSON(http.StatusOK, links)
}

// LinkCreate godoc
//
// @Summary Cria links no redirecionador com o código hash fornecido
//
// @Tags    Links
//
// @Accept  json
//
// @Produce json
//
// @Param   hash                path     string true "Código Hash"
//
// @Param   links               body     []models.LinkToBatchInsert true "Links"
//
// @Success 200                 {object} map[string]string
//
// @Failure 400                 {object} echo.HTTPError
//
// @Failure 500                 {object} echo.HTTPError
//
// @Router  /api/r/:hash/links [post]
func (s *Server) LinkCreate(c echo.Context) error {
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("LinkCreate", "Erro na validação do código hash do redirecionador dos links", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	parametros := struct {
		Links      []models.LinkToBatchInsert `json:"links"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("LinkCreate", "Não foram inseridos os links na requisição", nil)
		erros = append(erros, "Por favor, forneça os links no parâmetro 'links'.")
	}

	if len(parametros.Links) == 0 {
		utils.DebugLog("LinkCreate", "Erro porque não há nenhum link a ser inserido", nil)
		return utils.Erro(http.StatusInternalServerError, "Não foram passados links no parâmetro 'links'.")
	}

	for i, link := range parametros.Links {
		if err := utils.Validate.Var(link.Nome, "required,min=3,max=120"); err != nil {
			utils.DebugLog("LinkCreate", fmt.Sprintf("Erro no link %d: nome inválido para o parâmetro 'nome'", i+1), nil)
			erros = append(erros, fmt.Sprintf("Link %d: Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.", i+1))
		}

		if err := utils.Validate.Var(link.Link, "required"); err != nil {
			utils.DebugLog("LinkCreate", fmt.Sprintf("Erro no link %d: link inválido para o parâmetro 'link'", i+1), nil)
			erros = append(erros, fmt.Sprintf("Link %d: Por favor, forneça link válido (exemplo: 'https://t.me/+<numero_telefone>' ou 'https://wa.me/<numero_telefone>', a depender da plataforma) para o parâmetro 'link'.", i+1))
		}

		if err := utils.Validate.Var(link.Plataforma, "required,oneof=whatsapp telegram"); err != nil {
			utils.DebugLog("LinkCreate", fmt.Sprintf("Erro no link %d: plataforma inválida para o parâmetro 'plataforma'", i+1), nil)
			erros = append(erros, fmt.Sprintf("Link %d: Por favor, forneça uma plataforma válida ('telegram' ou 'whatsapp') para o parâmetro 'plataforma'.", i+1))
		}
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := s.LinkModel.Create(codigoHash, parametros.Links)

	if err != nil {
		utils.ErroLog("LinkCreate", "Erro na criação dos links do redirecionador inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível criar os links do redirecionador inserido.")
	}

	return c.JSON(http.StatusCreated, "Os links foram criados com sucesso.")
}

// LinkUpdate godoc
//
// @Summary Atualiza um link específico de um redirecionador específico
//
// @Tags    Links
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true  "Código Hash"
//
// @Param   id          path     int    true  "Id"
//
// @Param   nome        body     string false "Nome"
//
// @Param   link        body     string false "Link"
//
// @Param   plataforma  body     string false "Plataforma"
//
// @Success 200         {object} map[string]string
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /api/r/:hash/links/:id [patch]
func (s *Server) LinkUpdate(c echo.Context) error {
	id := c.Param("id")
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(id, "required,gte=0"); err != nil {
		utils.DebugLog("LinkUpdate", "Erro na validação do id do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'id' inserido é inválido, por favor insira um 'id' maior que 0.")
	}

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("LinkUpdate", "Erro na validação do código hash do redirecionador do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		utils.ErroLog("LinkUpdate", "Erro na transformação do id para inteiro", err)
		return utils.Erro(http.StatusBadRequest, "Houve um erro na transformação do id para um número inteiro. Por favor, insira um id válido.")
	}

	type parametrosUpdate struct {
		Nome       string `json:"nome"`
		Link       string `json:"link"`
		Plataforma string `json:"plataforma"`
	}

	parametros := parametrosUpdate{} 

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("LinkUpdate", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça o nome, link e plataforma nos parâmetro 'nome', 'link' e 'plataforma', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Nome, "min=3,max=120"); parametros.Nome != "" && err != nil {
		utils.DebugLog("LinkUpdate", "Erro no nome inválido para o parâmetro 'nome'", nil)
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.Plataforma, "oneof=whatsapp telegram"); parametros.Plataforma != "" && err != nil {
		utils.DebugLog("LinkUpdate", "Erro na plataforma inválida para o parâmetro 'plataforma'", nil)
		erros = append(erros, "Por favor, forneça uma plataforma válida ('telegram' ou 'whatsapp') para o parâmetro 'plataforma'.")
	}

	if (parametrosUpdate{}) == parametros {
		utils.DebugLog("LinkCreate", "Parâmetros para atualização estão vazios", nil)
		erros = append(erros, "Por favor, forneça algum valor válido para a atualização.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err = s.LinkModel.Update(parsedId, codigoHash, parametros.Nome, parametros.Link, parametros.Plataforma)

	if err != nil {
		utils.ErroLog("LinkUpdate", "Erro na atualização do link do redirecionador inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível atualizar o link do redirecionador inserido.")
	}

	return c.JSON(http.StatusOK, "O link foi atualizado com sucesso.")
}

// LinkDisable godoc
//
// @Summary Desativa um link específico de um redirecionador específico
//
// @Tags    Links
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true "Código Hash"
//
// @Param   id          path     int    true "Id"
//
// @Success 200         {object} map[string]string
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /api/r/:hash/links/:id/disable [patch]
func (s *Server) LinkDisable(c echo.Context) error {
	id := c.Param("id")
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(id, "required,gte=0"); err != nil {
		utils.DebugLog("LinkDisable", "Erro na validação do id do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'id' inserido é inválido, por favor insira um 'id' maior que 0.")
	}

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("LinkDisable", "Erro na validação do código hash do redirecionador do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		utils.ErroLog("LinkDisable", "Erro na transformação do id para inteiro", err)
		return utils.Erro(http.StatusBadRequest, "Houve um erro na transformação do id para um número inteiro. Por favor, insira um id válido.")
	}

	err = s.LinkModel.Disable(parsedId, codigoHash)

	if err != nil {
		utils.ErroLog("LinkDisable", "Erro ao desativar link do redirecionador inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível desativar o link do redirecionador inserido.")
	}

	return c.JSON(http.StatusOK, "O link foi desativado com sucesso.")
}

// LinkEnable godoc
//
// @Summary Ativa um link específico de um redirecionador específico
//
// @Tags    Links
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true "Código Hash"
//
// @Param   id          path     int    true "Id"
//
// @Success 200         {object} map[string]string
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /api/r/:hash/links/:id/enable [patch]
func (s *Server) LinkEnable(c echo.Context) error {
	id := c.Param("id")
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(id, "required,gte=0"); err != nil {
		utils.DebugLog("LinkEnable", "Erro na validação do id do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'id' inserido é inválido, por favor insira um 'id' maior que 0.")
	}

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("LinkEnable", "Erro na validação do código hash do redirecionador do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		utils.ErroLog("LinkEnable", "Erro na transformação do id para inteiro", err)
		return utils.Erro(http.StatusBadRequest, "Houve um erro na transformação do id para um número inteiro. Por favor, insira um id válido.")
	}

	err = s.LinkModel.Enable(parsedId, codigoHash)

	if err != nil {
		utils.ErroLog("LinkEnable", "Erro ao ativar o link do redirecionador inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ativar o link do redirecionador inserido.")
	}

	return c.JSON(http.StatusOK, "O link foi ativado com sucesso.")
}

// LinkRemove godoc
//
// @Summary Remove um link específico de um redirecionador específico
//
// @Tags    Links
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true "Código Hash"
//
// @Param   id          path     int    true "Id"
//
// @Success 200         {object} map[string]string
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /api/r/:hash/links/:id [delete]
func (s *Server) LinkRemove(c echo.Context) error {
	id := c.Param("id")
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(id, "required,gte=0"); err != nil {
		utils.DebugLog("LinkRemove", "Erro na validação do id do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'id' inserido é inválido, por favor insira um 'id' maior que 0.")
	}

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("LinkRemove", "Erro na validação do código hash do redirecionador do link", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	parsedId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		utils.ErroLog("LinkRemove", "Erro na transformação do id para inteiro", err)
		return utils.Erro(http.StatusBadRequest, "Houve um erro na transformação do id para um número inteiro. Por favor, insira um id válido.")
	}

	err = s.LinkModel.Remove(parsedId, codigoHash)

	if err != nil {
		utils.ErroLog("LinkRemove", "Erro na remoção do link do redirecionador inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível remover o link do redirecionador inserido.")
	}

	return c.JSON(http.StatusOK, "O link foi removido com sucesso.")
}
