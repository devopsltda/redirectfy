package server

import (
	"net/http"
	"strings"

	"redirectfy/internal/models"
	"redirectfy/internal/utils"

	"github.com/labstack/echo/v4"

	_ "redirectfy/internal/models"
)

type RedirecionadorReadByCodigoHashResponse struct {
	R models.Redirecionador `json:"redirecionador"`
	L []models.Link         `json:"links"`
}

// RedirecionadorReadAll godoc
//
// @Summary Retorna os redirecionadores de um usuário específico
//
// @Tags    Redirecionadores
//
// @Accept  json
//
// @Produce json
//
// @Success 200         {object} models.Redirecionador
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /r [get]
func (s *Server) RedirecionadorReadAll(c echo.Context) error {
	cookie, err := c.Cookie("usuario")

	if err != nil {
		utils.DebugLog("RedirecionadorReadAll", "Erro na leitura do nome do usuário pelo cookie", err)
		return utils.Erro(http.StatusBadRequest, "Não foi possível ler o nome de usuário.")
	}

	redirecionadores, err := s.RedirecionadorModel.ReadAll(cookie.Value)

	if err != nil {
		utils.ErroLog("RedirecionadorReadAll", "Erro na leitura dos redirecionadores do usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler os redirecionadores do usuário.")
	}

	return c.JSON(http.StatusOK, redirecionadores)
}

// RedirecionadorReadByCodigoHash godoc
//
// @Summary Retorna o redirecionador com o código hash fornecido
//
// @Tags    Redirecionadores
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true "Código Hash"
//
// @Success 200         {object} models.Redirecionador
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /r/:hash [get]
func (s *Server) RedirecionadorReadByCodigoHash(c echo.Context) error {
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("RedirecionadorReadByCodigoHash", "Erro na validação do código hash do redirecionador", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	redirecionador, err := s.RedirecionadorModel.ReadByCodigoHash(codigoHash)

	if err != nil {
		utils.ErroLog("RedirecionadorReadByCodigoHash", "Erro na leitura do redirecionador do usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler os redirecionadores do usuário.")
	}

	links, err := s.LinkModel.ReadByCodigoHash(codigoHash)

	if err != nil {
		utils.ErroLog("RedirecionadorReadByCodigoHash", "Erro na leitura dos links do redirecionador do usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler os links do redirecionador do usuário.")
	}

	return c.JSON(http.StatusOK, RedirecionadorReadByCodigoHashResponse{
		R: redirecionador,
		L: links,
	})
}

// RedirecionadorLinksToGoTo godoc
//
// @Summary Retorna os links selecionados daquele redirecionador
//
// @Tags    Redirecionadores
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
// @Router  /r/:hash [get]
func (s *Server) RedirecionadorLinksToGoTo(c echo.Context) error {
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("RedirecionadorLinksToGoTo", "Erro na validação do código hash do redirecionador", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	redirecionador, err := s.RedirecionadorModel.ReadByCodigoHash(codigoHash)

	if err != nil {
		utils.ErroLog("RedirecionadorLinksToGoTo", "Erro na leitura do redirecionador", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o redirecionador.")
	}

	links, err := s.LinkModel.ReadByCodigoHash(codigoHash)

	if err != nil {
		utils.ErroLog("RedirecionadorLinksToGoTo", "Erro na leitura dos links do redirecionador", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler os links do redirecionador.")
	}

	usuario, err := s.UsuarioModel.ReadByNomeDeUsuario(redirecionador.Usuario)

	if err != nil {
		utils.ErroLog("RedirecionadorLinksToGoTo", "Erro na leitura do usuário do redirecionador", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o usuário do redirecionador.")
	}

	picked_links := models.LinkPicker(links, strings.HasPrefix(usuario.PlanoDeAssinatura, "Pro"))

	return c.JSON(http.StatusOK, RedirecionadorReadByCodigoHashResponse{
		R: redirecionador,
		L: picked_links,
	})
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
// @Success 200                       {object} map[string]string
//
// @Failure 400                       {object} echo.HTTPError
//
// @Failure 500                       {object} echo.HTTPError
//
// @Router  /r [post]
func (s *Server) RedirecionadorCreate(c echo.Context) error {
	cookie, err := c.Cookie("usuario")

	if err != nil {
		utils.DebugLog("RedirecionadorReadAll", "Erro na leitura do nome do usuário pelo cookie", err)
		return utils.Erro(http.StatusBadRequest, "Não foi possível ler o nome de usuário.")
	}

	parametros := struct {
		Nome                    string `json:"nome"`
		OrdemDeRedirecionamento string `json:"ordem_de_redirecionamento"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("RedirecionadorCreate", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça o nome e ordem de redirecionamento nos parâmetro 'nome' e 'ordem_de_redirecionamento', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Nome, "required,min=3,max=120"); err != nil {
		utils.DebugLog("RedirecionadorCreate", "Erro no nome inválido para o parâmetro 'nome'", nil)
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.OrdemDeRedirecionamento, "required,max=120,oneof=telegram0x2Cwhatsapp whatsapp0x2Ctelegram"); err != nil {
		utils.DebugLog("RedirecionadorCreate", "Erro na ordem de redirecionamento inválida para o parâmetro 'ordem_de_redirecionamento'", nil)
		erros = append(erros, "Por favor, forneça uma ordem de redirecionamento válida ('telegram,whatsapp' ou 'whatsapp,telegram') para o parâmetro 'ordem_de_redirecionamento'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	var codigoHash string
	codigoHashExiste := true

	for codigoHashExiste {
		codigoHash = utils.GeraHashCode(10)

		codigoHashExiste, err = s.RedirecionadorModel.CheckIfCodigoHashExists(codigoHash)

		if err != nil {
			utils.ErroLog("RedirecionadorCreate", "Erro na checagem da existência do código hash", err)
			return utils.Erro(http.StatusInternalServerError, "Não foi possível verificar se havia um código hash disponível para o novo redirecionador.")
		}
	}

	_, err = s.RedirecionadorModel.Create(
		parametros.Nome,
		codigoHash,
		parametros.OrdemDeRedirecionamento,
		cookie.Value,
	)

	if err != nil {
		utils.ErroLog("RedirecionadorCreate", "Erro na criação do redirecionador", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível criar o redirecionador.")
	}

	return c.JSON(http.StatusCreated, codigoHash)
}

// RedirecionadorRefresh godoc
//
// @Summary Recria o hash de um redirecionador
//
// @Tags    Redirecionadores
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true "Código Hash"
//
// @Success 200         {object} map[string]string
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /r/:hash/refresh [patch]
func (s *Server) RedirecionadorRefresh(c echo.Context) error {
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("RedirecionadorRefresh", "Erro na validação do código hash do redirecionador", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	var err error
	var codigoHashNovo string
	codigoHashExiste := true

	for codigoHashExiste {
		codigoHashNovo = utils.GeraHashCode(10)

		codigoHashExiste, err = s.RedirecionadorModel.CheckIfCodigoHashExists(codigoHashNovo)

		if err != nil {
			utils.ErroLog("RedirecionadorRefresh", "Erro na checagem da existência do código hash", err)
			return utils.Erro(http.StatusInternalServerError, "Não foi possível verificar se havia um código hash disponível para o novo redirecionador.")
		}
	}

	err = s.RedirecionadorModel.Rehash(codigoHash, codigoHashNovo)

	if err != nil {
		utils.ErroLog("RedirecionadorRefresh", "Erro na atualização do código hash do redirecionador", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível atualizar o código hash do redirecionador.")
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
// @Param   hash                      path     string true  "Código Hash"
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
// @Router  /r/:hash [patch]
func (s *Server) RedirecionadorUpdate(c echo.Context) error {
	parametros := struct {
		Nome                    string `json:"nome"`
		OrdemDeRedirecionamento string `json:"ordem_de_redirecionamento"`
	}{}

	var erros []string

	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("RedirecionadorUpdate", "Erro na validação do código hash do redirecionador", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("RedirecionadorUpdate", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça o nome e/ou a ordem de redirecionamento nos parâmetros 'nome' e 'ordem_de_redirecionamento', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Nome, "min=3,max=120"); parametros.Nome != "" && err != nil {
		utils.DebugLog("RedirecionadorUpdate", "Erro no nome inválido para o parâmetro 'nome'", nil)
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.OrdemDeRedirecionamento, "max=120,oneof=telegram0x2Cwhatsapp whatsapp0x2Ctelegram"); parametros.OrdemDeRedirecionamento != "" && err != nil {
		utils.DebugLog("RedirecionadorUpdate", "Erro na ordem de redirecionamento inválida para o parâmetro 'ordem_de_redirecionamento'", nil)
		erros = append(erros, "Por favor, forneça uma ordem de redirecionamento válida ('telegram,whatsapp' ou 'whatsapp,telegram') para o parâmetro 'ordem_de_redirecionamento'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	err := s.RedirecionadorModel.Update(parametros.Nome, codigoHash, parametros.OrdemDeRedirecionamento)

	if err != nil {
		utils.ErroLog("RedirecionadorUpdate", "Erro na atualização do redirecionador", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível atualizar o redirecionador.")
	}

	return c.JSON(http.StatusOK, "O redirecionador foi atualizado com sucesso.")
}

// RedirecionadorRemove godoc
//
// @Summary Remove um redirecionador
//
// @Tags    Redirecionadores
//
// @Accept  json
//
// @Produce json
//
// @Param   hash        path     string true "Código Hash"
//
// @Success 200         {object} map[string]string
//
// @Failure 400         {object} echo.HTTPError
//
// @Failure 500         {object} echo.HTTPError
//
// @Router  /r/:hash [delete]
func (s *Server) RedirecionadorRemove(c echo.Context) error {
	codigoHash := c.Param("hash")

	if err := utils.Validate.Var(codigoHash, "required,len=10"); err != nil {
		utils.DebugLog("RedirecionadorRemove", "Erro na validação do código hash do redirecionador", err)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente com 10 caracteres.")
	}

	err := s.RedirecionadorModel.Remove(codigoHash)

	if err != nil {
		utils.ErroLog("RedirecionadorRemove", "Erro na remoção do redirecionador inserido", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível remover o redirecionador inserido.")
	}

	return c.JSON(http.StatusOK, "O redirecionador foi removido com sucesso.")
}
