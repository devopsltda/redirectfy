package server

import (
	"database/sql"
	"net/http"
	"strings"
	"time"
	"unicode"

	"redirectfy/internal/auth"
	"redirectfy/internal/utils"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	_ "redirectfy/internal/models"
)

type productData struct {
	Name string `json:"name"`
}

type customerData struct {
	Name     string `json:"name"`
	Document string `json:"document"`
	Email    string `json:"email"`
}

// Esses parâmetros estão dispostos aqui como é documentado
// na configuração de web hooks da Kirvano em:
// https://help.kirvano.com/pt-BR/articles/8141372-o-que-e-e-como-configurar-webhooks
type parametrosKirvano struct {
	Event    string        `json:"event"`
	Customer customerData  `json:"customer"`
	Products []productData `json:"products"`
} // @name ParametrosKirvano

// Essa função é utilizada para criar um nome de usuário com base no email
// fornecido pela Kirvano.
func criaNomeDeUsuario(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c == '@' {
			break
		}

		if unicode.IsLetter(c) || unicode.IsNumber(c) || c == '_' || c == '-' {
			sb.WriteRune(c)
		} else {
			sb.WriteRune('_')
		}
	}

	return sb.String()
}

// UsuarioReadByNomeDeUsuario godoc
//
// @Summary Retorna o usuário com o nome de usuário fornecido
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Success 200             {object} models.Usuario
//
// @Failure 400             {object} echo.HTTPError
//
// @Failure 500             {object} echo.HTTPError
//
// @Router  /api/u [get]
func (s *Server) UsuarioReadByNomeDeUsuario(c echo.Context) error {
	nomeDeUsuario := c.Get("usuario").(*jwt.Token).Claims.(*auth.Claims).NomeDeUsuario

	usuario, err := s.UsuarioModel.ReadByNomeDeUsuario(nomeDeUsuario)

	if err != nil {
		utils.ErroLog("UsuarioReadByNomeDeUsuario", "Erro na leitura do usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o usuário.")
	}

	return c.JSON(http.StatusOK, usuario)
}

// KirvanoCreate godoc
//
// @Summary Cria um usuário da Kirvano
//
// @Tags    Kirvano
//
// @Accept  json
//
// @Produce json
//
// @Param   request             body     parametrosKirvano true "Request"
//
// @Success 200                 {object} map[string]string
//
// @Failure 400                 {object} echo.HTTPError
//
// @Failure 500                 {object} echo.HTTPError
//
// @Router  /api/kirvano [post]
func (s *Server) KirvanoCreate(c echo.Context) error {
	if c.Request().Header.Get("Security-Token") != "abc123" {
				utils.ErroLog("KirvanoCreate", "Erro no token de segurança passado pela Kirvano", nil)
				return utils.Erro(http.StatusUnauthorized, "Você não tem o que é necessário para o usuário.")
	}

	var parametros parametrosKirvano

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("KirvanoCreate", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça o CPF, email, nome e plano de assinatura do usuário nos parâmetros 'customer.document', 'customer.email', 'customer.name' e 'products.name', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Event, "required,oneof=SALE_APPROVED SALE_REFUNDED SUBSCRIPTION_CANCELED SUBSCRIPTION_EXPIRED SUBSCRIPTION_RENEWED"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no nome de evento inválido no parâmetro 'event'", nil)
		erros = append(erros, "Por favor, forneça um evento válido (SALE_APPROVED, SALE_REFUNDED, SUBSCRIPTION_CANCELED, SUBSCRIPTION_EXPIRED ou SUBSCRIPTION_RENEWED) para o parâmetro 'event'.")
	}

	if err := utils.Validate.Var(parametros.Customer.Document, "required,numeric,len=11"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no CPF inválido no parâmetro 'customer.document'", nil)
		erros = append(erros, "Por favor, forneça um CPF válido (texto numérico com 11 dígitos) para o parâmetro 'customer.document'.")
	}

	if err := utils.Validate.Var(parametros.Customer.Name, "required,min=3,max=240"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no nome inválido no parâmetro 'customer.name'", nil)
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 240 caracteres) para o parâmetro 'customer.name'.")
	}

	if err := utils.Validate.Var(parametros.Customer.Email, "required,email"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no email inválido no parâmetro 'customer.email'", nil)
		erros = append(erros, "Por favor, forneça um email válido para o parâmetro 'customer.email'.")
	}

	if len(parametros.Products) != 1 {
		utils.DebugLog("KirvanoCreate", "Erro no parâmetro 'products' que possui mais de um produto", nil)
		erros = append(erros, "Por favor, forneça um único produto no parâmetro 'products'.")
	} else if err := utils.Validate.Var(parametros.Products[0].Name, "required,min=3,max=120"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no nome inválido no parâmetro 'products.name'", nil)
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 120 caracteres) no parâmetro 'products.name'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	nomeDeUsuario := criaNomeDeUsuario(parametros.Customer.Email)

	switch parametros.Event {
	case "SALE_APPROVED":
		_, err := s.UsuarioModel.ReadByNomeDeUsuario(nomeDeUsuario)

		if err == nil {
			err := s.UsuarioModel.UpdatePlanoDeAssinatura(nomeDeUsuario, parametros.Products[0].Name)

			if err != nil {
				utils.ErroLog("KirvanoCreate", "Erro na atualização do usuário", err)
				return utils.Erro(http.StatusInternalServerError, "Não foi possível atualizar o usuário.")
			}

			return c.JSON(http.StatusCreated, "O plano de assinatura do usuário foi atualizado com sucesso.")
		}

		usuarioId, err := s.UsuarioKirvanoModel.Create(
			parametros.Customer.Document,
			parametros.Customer.Name,
			nomeDeUsuario,
			parametros.Customer.Email,
			parametros.Products[0].Name,
		)

		if err != nil {
			utils.ErroLog("KirvanoCreate", "Erro na criação do usuário da Kirvano", err)
			return utils.Erro(http.StatusInternalServerError, "Não foi possível criar o usuário da Kirvano.")
		}

		var valor string
		valorExiste := true

		for valorExiste {
			valor = utils.GeraHashCode(120)

			valorExiste, err = s.EmailAutenticacaoModel.CheckIfValorExists(valor)

			if err != nil {
				utils.ErroLog("KirvanoCreate", "Erro na checagem da existência do valor do email de autenticação", err)
				return utils.Erro(http.StatusInternalServerError, "Não foi possível verificar se havia um valor disponível para o email de autenticação.")
			}
		}

		id, err := s.EmailAutenticacaoModel.Create(valor, "nova_senha", usuarioId)

		if err != nil {
			utils.ErroLog("KirvanoCreate", "Erro na criação do email de autenticação", err)
			return utils.Erro(http.StatusInternalServerError, "Não foi possível criar email de autenticação para o usuário da Kirvano.")
		}

		err = s.email.SendValidacao(id, parametros.Customer.Name, valor, parametros.Customer.Email)

		if err != nil {
			utils.ErroLog("KirvanoCreate", "Erro no envio do email de autenticação", err)
			return utils.Erro(http.StatusInternalServerError, "Não foi possível enviar o email de autenticação para o usuário da Kirvano.")
		}
	case "SALE_REFUNDED", "SUBSCRIPTION_CANCELED", "SUBSCRIPTION_EXPIRED":
		err := s.UsuarioModel.UpdatePlanoDeAssinatura(nomeDeUsuario, "Gratuito")

		if err != nil {
			utils.ErroLog("KirvanoCreate", "Erro na atualização do usuário", err)
			return utils.Erro(http.StatusInternalServerError, "Não foi possível atualizar o usuário.")
		}
	}

	return c.JSON(http.StatusCreated, "O usuário da Kirvano foi criado com sucesso.")
}

// KirvanoToUser godoc
//
// @Summary Cria um usuário a partir de um usuário da Kirvano
//
// @Tags    Kirvano
//
// @Accept  json
//
// @Produce json
//
// @Param   hash               path     string true "Valor"
//
// @Param   senha              body     string true "Senha"
//
// @Param   data_de_nascimento body     string true "Data de Nascimento"
//
// @Success 200                {object} map[string]string
//
// @Failure 400                {object} echo.HTTPError
//
// @Failure 500                {object} echo.HTTPError
//
// @Router  /api/kirvano/to_user/:hash [post]
func (s *Server) KirvanoToUser(c echo.Context) error {
	valor := c.Param("hash")

	if !utils.IsURLSafe(valor) {
		utils.DebugLog("KirvanoToUser", "Erro na validação do código hash do email de autenticação", nil)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente.")
	}

	parametros := struct {
		Senha            string `json:"senha"`
		DataDeNascimento string `json:"data_de_nascimento"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("KirvanoToUser", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça a senha e data de nascimento nos parâmetros 'senha' e 'data_de_nascimento', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Senha, "required,min=8,max=72"); err != nil {
		utils.DebugLog("KirvanoToUser", "Erro na senha inválida no parâmetro 'senha'", nil)
		erros = append(erros, "Por favor, forneça uma senha válida (8 a 72 caracteres) para o parâmetro 'senha'.")
	}

	if err := utils.Validate.Var(parametros.DataDeNascimento, "required,datetime=2006-01-02"); err != nil {
		utils.DebugLog("KirvanoToUser", "Erro na data de nascimento inválida no parâmetro 'data_de_nascimento'", nil)
		erros = append(erros, "Por favor, forneça uma data de nascimento válida (No formato 'YYYY-MM-DD') para o parâmetro 'data_de_nascimento'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	usuario, err := s.EmailAutenticacaoModel.CheckIfValorExistsAndIsValid(valor, "nova_senha")

	if err != nil {
		utils.ErroLog("KirvanoToUser", "Erro na checagem da existência do valor do email de autenticação", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível verificar se existia um email de autenticação com o hash inserido.")
	}

	senhaComHash, err := argon2id.CreateHash(parametros.Senha+utils.Pepper, utils.SenhaParams)

	if err != nil {
		utils.ErroLog("KirvanoToUser", "Erro ao criar senha do usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível criar a senha do usuário.")
	}

	usuarioTemporario, err := s.UsuarioKirvanoModel.ReadById(usuario)

	if err != nil {
		utils.ErroLog("KirvanoToUser", "Erro ao ler o usuário da Kirvano", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o usuário da Kirvano do email de autenticação.")
	}

	_, err = s.UsuarioModel.Create(
		usuarioTemporario.Cpf,
		usuarioTemporario.Nome,
		usuarioTemporario.NomeDeUsuario,
		usuarioTemporario.Email,
		senhaComHash,
		parametros.DataDeNascimento,
		usuarioTemporario.PlanoDeAssinatura,
	)

	if err != nil {
		utils.ErroLog("KirvanoToUser", "Erro ao criar o usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível criar o usuário.")
	}

	err = s.UsuarioKirvanoModel.Remove(usuario)

	if err != nil {
		utils.ErroLog("KirvanoToUser", "Erro ao remover o usuário da Kirvano", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível remover o usuário da Kirvano do email de autenticação.")
	}

	err = s.EmailAutenticacaoModel.Expirar(valor)

	if err != nil {
		utils.ErroLog("KirvanoToUser", "Erro ao expirar o email de autenticação", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível expirar o email de autenticação com o hash inserido.")
	}

	return c.JSON(http.StatusOK, "O usuário foi criado com sucesso.")
}

// UsuarioSolicitarTrocaDeSenha godoc
//
// @Summary Exige a troca de senha de um usuário
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Param   email             body     string true "Email"
//
// @Success 200               {object} map[string]string
//
// @Failure 400               {object} echo.HTTPError
//
// @Failure 500               {object} echo.HTTPError
//
// @Router  /api/u/change_password [patch]
func (s *Server) UsuarioSolicitarTrocaDeSenha(c echo.Context) error {
	parametros := struct {
		Email            string `json:"email"`
	}{}

	var err error
	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("UsuarioSolicitarTrocaDeSenha", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça o email no parâmetro 'email'.")
	}

	if err := utils.Validate.Var(parametros.Email, "required,email"); parametros.Email != "" && err != nil {
		utils.DebugLog("UsuarioSolicitarTrocaDeSenha", "Erro no email inválido no parâmetro 'email'", nil)
		erros = append(erros, "Por favor, forneça o email do usuário válido no parâmetro 'email'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	usuario, err := s.UsuarioModel.ReadByEmail(parametros.Email)

	if err != nil {
		utils.ErroLog("UsuarioSolicitarTrocaDeSenha", "Erro ao procurar um usuário com esse email", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi procurar um usuário com esse email.")
	}

	var valor string
	valorExiste := true

	for valorExiste {
		valor = utils.GeraHashCode(120)

		valorExiste, err = s.EmailAutenticacaoModel.CheckIfValorExists(valor)

		if err != nil {
			utils.ErroLog("UsuarioSolicitarTrocaDeSenha", "Erro na checagem da existência do valor do email de troca de senha", err)
			return utils.Erro(http.StatusInternalServerError, "Não foi possível verificar se havia um valor disponível para o email de troca de senha.")
		}
	}

	id, err := s.EmailAutenticacaoModel.Create(valor, "senha", usuario.Id)

	if err != nil {
		utils.ErroLog("UsuarioSolicitarTrocaDeSenha", "Erro ao criar email de troca de senha", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível criar um email de troca de senha para o usuário.")
	}

	err = s.email.SendTrocaDeSenha(id, usuario.Nome, valor, parametros.Email)

	if err != nil {
		utils.ErroLog("UsuarioSolicitarTrocaDeSenha", "Erro ao enviar email de troca de senha", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível enviar o email de troca de senha para o usuário.")
	}

	return c.JSON(http.StatusOK, "O pedido de troca de senha do usuário foi enviado com sucesso.")
}

// UsuarioTrocaDeSenha godoc
//
// @Summary Troca a senha de um usuário
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Param   hash              path     string true "Valor"
//
// @Param   senha_nova        body     string true "Senha Nova"
//
// @Success 200               {object} map[string]string
//
// @Failure 400               {object} echo.HTTPError
//
// @Failure 500               {object} echo.HTTPError
//
// @Router  /api/u/change_password/:hash [patch]
func (s *Server) UsuarioTrocaDeSenha(c echo.Context) error {
	valor := c.Param("hash")

	if !utils.IsURLSafe(valor) {
		utils.DebugLog("UsuarioTrocaDeSenha", "Erro na validação do código hash do email de autenticação", nil)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente.")
	}

	parametros := struct {
		SenhaNova string `json:"senha_nova"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("UsuarioTrocaDeSenha", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça a senha nova no parâmetro 'senha_nova'.")
	}

	if err := utils.Validate.Var(parametros.SenhaNova, "required,min=8,max=72"); err != nil {
		utils.DebugLog("UsuarioTrocaDeSenha", "Erro na senha inválida no parâmetro 'senha_nova'", nil)
		erros = append(erros, "Por favor, forneça uma senha válida (8 a 72 caracteres) para o parâmetro 'senha_nova'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	usuario, err := s.EmailAutenticacaoModel.CheckIfValorExistsAndIsValid(valor, "senha")

	if err != nil {
		utils.ErroLog("UsuarioTrocaDeSenha", "Erro na checagem da existência do valor do email de troca de senha", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível verificar se existia um email de troca de senha com o hash inserido.")
	}

	senhaComHash, err := argon2id.CreateHash(parametros.SenhaNova+utils.Pepper, utils.SenhaParams)

	if err != nil {
		utils.ErroLog("UsuarioTrocaDeSenha", "Erro ao criar a nova senha do usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível criar a nova senha do usuário.")
	}

	err = s.UsuarioModel.TrocaSenha(usuario, senhaComHash)

	if err != nil {
		utils.ErroLog("UsuarioTrocaDeSenha", "Erro ao trocar a senha do usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível trocar a senha do usuário.")
	}

	err = s.EmailAutenticacaoModel.Expirar(valor)

	if err != nil {
		utils.ErroLog("UsuarioTrocaDeSenha", "Erro ao expirar o email de troca de senha", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível expirar o email de troca de senha com o hash inserido.")
	}

	return c.JSON(http.StatusOK, "A senha do usuário foi trocada com sucesso.")
}

// UsuarioLogin godoc
//
// @Summary Autentica o usuário
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Param   email             body     string true "Email"
//
// @Param   senha             body     string true  "Senha"
//
// @Success 200               {object} map[string]string
//
// @Failure 400               {object} echo.HTTPError
//
// @Failure 401               {object} echo.HTTPError
//
// @Failure 500               {object} echo.HTTPError
//
// @Router  /api/u/login [post]
func (s *Server) UsuarioLogin(c echo.Context) error {
	parametros := struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("UsuarioLogin", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça o nome de usuário do usuário no parâmetro 'nome_de_usuario'.")
	}

	if err := utils.Validate.Var(parametros.Email, "required,email"); parametros.Email != "" && err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no email inválido no parâmetro 'email'", nil)
		erros = append(erros, "Por favor, forneça o email do usuário válido no parâmetro 'email'.")
	}

	if err := utils.Validate.Var(parametros.Senha, "required,min=8,max=72"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro na senha inválida no parâmetro 'senha'", nil)
		erros = append(erros, "Por favor, forneça uma senha válida (texto de 8 a 72 caracteres, podendo conter letras, números e símbolos) para o parâmetro 'senha'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	id, nome, nomeDeUsuario, planoDeAssinatura, senha, err := s.UsuarioModel.Login(parametros.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.DebugLog("UsuarioLogin", "Erro ao ler o email do usuário", err)
			return utils.Erro(http.StatusBadRequest, "Usuário ou senha incorretos.")
		}

		utils.DebugLog("UsuarioLogin", "Erro ao ler os dados do usuário", err)
		return utils.Erro(http.StatusInternalServerError, "Não foi possível ler o usuário com o email inserido.")
	}

	match, err := argon2id.ComparePasswordAndHash(parametros.Senha+utils.Pepper, senha)

	if !match || err != nil {
		utils.DebugLog("UsuarioLogin", "Erro ao validar a senha do usuário", err)
		return utils.Erro(http.StatusBadRequest, "Usuário ou senha incorretos.")
	}

	err = auth.GeraTokensESetaCookies(id, nome, nomeDeUsuario, planoDeAssinatura, c)

	if err != nil {
		utils.DebugLog("UsuarioLogin", "Erro na assinatura do token JWT", err)
		return utils.Erro(http.StatusUnauthorized, "O seu token de autenticação não pôde ser criado.")
	}

	return c.JSON(http.StatusOK, "O usuário foi autenticado com sucesso.")
}

// UsuarioLogout godoc
//
// @Summary Remove cookies de autenticação do usuário
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Success 200               {object} map[string]string
//
// @Router  /api/u/logout [post]
func (s *Server) UsuarioLogout(c echo.Context) error {
	for _, c := range c.Cookies() {
		c.Expires = time.Now().Add(-48 * time.Hour)
	}

	return c.JSON(http.StatusOK, "A autenticação do usuário foi expirada com sucesso.")
}
