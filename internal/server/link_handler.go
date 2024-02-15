package server

import (
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirectify/internal/model"
	"github.com/TheDevOpsCorp/redirectify/internal/util"
	"github.com/labstack/echo/v4"
)

// LinkReadByCodigoHash godoc
//
// @Summary Retorna o link com o código hash fornecido
// @Tags    Link
// @Accept  json
// @Produce json
// @Param   codigo_hash path     string true "Nome de Usuário"
// @Success 200         {object} model.Link
// @Failure 400         {object} util.Erro
// @Failure 500         {object} util.Erro
// @Router  /api/link/:codigo_hash [get]
func (s *Server) LinkReadByCodigoHash(c echo.Context) error {
	codigoHash := c.Param("codigo_hash")

	if err := util.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return util.ErroValidacaoCodigoHash
	}

	link, err := model.LinkReadByCodigoHash(s.db, codigoHash)

	if err != nil {
		log.Printf("LinkReadByCodigoHash: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, link)
}

// LinkReadAll godoc
//
// @Summary Retorna os links
// @Tags    Link
// @Accept  json
// @Produce json
// @Success 200 {object} []model.Link
// @Failure 400 {object} util.Erro
// @Failure 500 {object} util.Erro
// @Router  /api/link [get]
func (s *Server) LinkReadAll(c echo.Context) error {
	links, err := model.LinkReadAll(s.db)

	if err != nil {
		log.Printf("LinkReadAll: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, links)
}

// LinkCreate godoc
//
// @Summary Cria um link
// @Tags    Link
// @Accept  json
// @Produce json
// @Param   nome                      body     string true  "Nome"
// @Param   link_whatsapp             body     string false "Link Whatsapp"
// @Param   link_telegram             body     string false "Link Telegram"
// @Param   ordem_de_redirecionamento body     string true  "Ordem de Redirecionamento"
// @Param   usuario                   body     string int   "Usuário"
// @Success 200                       {object} map[string]string
// @Failure 400                       {object} util.Erro
// @Failure 500                       {object} util.Erro
// @Router  /api/link [post]
func (s *Server) LinkCreate(c echo.Context) error {
	parametros := struct {
		Nome                    string `json:"nome"`
		LinkWhatsapp            string `json:"link_whatsapp"`
		LinkTelegram            string `json:"link_telegram"`
		OrdemDeRedirecionamento string `json:"ordem_de_redirecionamento"`
		Usuario                 int64  `json:"usuario"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome, link do whatsapp, link do telegram, ordem de redirecionamento e usuário nos parâmetro 'nome', 'link_whatsapp', 'link_telegram', 'ordem_de_redirecionamento' e 'usuario', respectivamente.")
	}

	if err := util.Validate.Var(parametros.Nome, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := util.Validate.Var(parametros.LinkWhatsapp, "required,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um link de whatsapp válido para o parâmetro 'link_whatsapp'.")
	}

	if err := util.Validate.Var(parametros.LinkTelegram, "required,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um link de telegram válido para o parâmetro 'link_telegram'.")
	}

	if err := util.Validate.Var(parametros.OrdemDeRedirecionamento, "required,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça uma ordem de redirecionamento válida para o parâmetro 'ordem_de_redirecionamento'.")
	}

	if err := util.Validate.Var(parametros.Usuario, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um usuário válido para o parâmetro 'usuario'.")
	}

	if len(erros) > 0 {
		return util.ErroValidacaoParametro(erros)
	}

	var err error
	var codigoHash string
	codigoHashExiste := true

	for codigoHashExiste {
		codigoHash = util.GeraHashCode(10)

		codigoHashExiste, err = model.LinkCheckIfCodigoHashExists(s.db, codigoHash)

		if err != nil {
			log.Printf("LinkCreate: %v", err)
			return util.ErroBancoDados
		}
	}

	err = model.LinkCreate(
		s.db,
		parametros.Nome,
		codigoHash,
		parametros.LinkWhatsapp,
		parametros.LinkTelegram,
		parametros.OrdemDeRedirecionamento,
		parametros.Usuario,
	)

	if err != nil {
		log.Printf("LinkCreate: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O link foi adicionado com sucesso.",
	})
}

// LinkUpdate godoc
//
// @Summary Atualiza um link
// @Tags    Link
// @Accept  json
// @Produce json
// @Param   codigo_hash               path     string true  "Código Hash"
// @Param   nome                      body     string false "Nome"
// @Param   link_whatsapp             body     string false "Link Whatsapp"
// @Param   link_telegram             body     string false "Link Telegram"
// @Param   ordem_de_redirecionamento body     string false "Ordem de Redirecionamento"
// @Success 200                       {object} map[string]string
// @Failure 400                       {object} util.Erro
// @Failure 500                       {object} util.Erro
// @Router  /api/link/:codigo_hash [patch]
func (s *Server) LinkUpdate(c echo.Context) error {
	parametros := struct {
		Nome                    string `json:"nome"`
		LinkWhatsapp            string `json:"link_whatsapp"`
		LinkTelegram            string `json:"link_telegram"`
		OrdemDeRedirecionamento string `json:"ordem_de_redirecionamento"`
	}{}

	var erros []string

	codigoHash := c.Param("codigo_hash")

	if err := util.Validate.Var(codigoHash, "required,len=10"); err != nil {
		log.Printf("LinkUpdate: %v", err)
		return util.ErroValidacaoCodigoHash
	}

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome, link do whatsapp, link do telegram, ordem de redirecionamento e usuário nos parâmetro 'nome', 'link_whatsapp', 'link_telegram', 'ordem_de_redirecionamento' e 'usuario', respectivamente.")
	}

	if err := util.Validate.Var(parametros.Nome, "min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) para o parâmetro 'nome'.")
	}

	if err := util.Validate.Var(parametros.LinkWhatsapp, "max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um link de whatsapp válido para o parâmetro 'link_whatsapp'.")
	}

	if err := util.Validate.Var(parametros.LinkTelegram, "max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um link de telegram válido para o parâmetro 'link_telegram'.")
	}

	if err := util.Validate.Var(parametros.OrdemDeRedirecionamento, "max=120"); err != nil {
		erros = append(erros, "Por favor, forneça uma ordem de redirecionamento válida para o parâmetro 'ordem_de_redirecionamento'.")
	}

	if len(erros) > 0 {
		return util.ErroValidacaoParametro(erros)
	}

	err := model.LinkUpdate(s.db, parametros.Nome, codigoHash, parametros.LinkWhatsapp, parametros.LinkTelegram, parametros.OrdemDeRedirecionamento)

	if err != nil {
		log.Printf("LinkUpdate: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O link foi atualizado com sucesso.",
	})
}

// LinkRemove godoc
//
// @Summary Remove um link
// @Tags    Link
// @Accept  json
// @Produce json
// @Param   codigo_hash path     string true "Código Hash"
// @Success 200         {object} map[string]string
// @Failure 400         {object} util.Erro
// @Failure 500         {object} util.Erro
// @Router  /api/link/:codigo_hash [delete]
func (s *Server) LinkRemove(c echo.Context) error {
	codigoHash := c.Param("codigo_hash")

	if err := util.Validate.Var(codigoHash, "required,len=10"); err != nil {
		return util.ErroValidacaoCodigoHash
	}

	err := model.LinkRemove(s.db, codigoHash)

	if err != nil {
		log.Printf("LinkRemove: %v", err)
		return util.ErroBancoDados
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O link foi removido com sucesso.",
	})
}
