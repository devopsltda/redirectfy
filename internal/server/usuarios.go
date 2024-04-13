package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"unicode"

	"redirectfy/internal/auth"
	"redirectfy/internal/utils"

	"github.com/alexedwards/argon2id"
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

type parametrosKirvano struct {
	Event    string        `json:"event"`
	Customer customerData  `json:"customer"`
	Products []productData `json:"products"`
} // @name ParametrosKirvano

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
// @Param   username        path     string  true  "Nome de Usuário"
//
// @Success 200             {object} models.Usuario
//
// @Failure 400             {object} echo.HTTPError
//
// @Failure 500             {object} echo.HTTPError
//
// @Router  /u/:username [get]
func (s *Server) UsuarioReadByNomeDeUsuario(c echo.Context) error {
	nomeDeUsuario := c.Param("username")

	if !utils.IsURLSafe(nomeDeUsuario) {
		utils.DebugLog("UsuarioReadByNomeDeUsuario", "Erro de validação do nome do usuário no parâmetro 'username'", nil)
		return utils.Erro(http.StatusBadRequest, "Por favor, forneça um nome de usuário válido (3 a 120 caracteres, composto de apenas letras, números e '-' ou '_').")
	}

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
// @Router  /kirvano [post]
func (s *Server) KirvanoCreate(c echo.Context) error {
	var parametros parametrosKirvano

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		utils.DebugLog("KirvanoCreate", "Não foram inseridos os parâmetros na requisição", nil)
		erros = append(erros, "Por favor, forneça o CPF, email, data de nascimento, nome, nome de usuário, senha e plano de assinatura do usuário nos parâmetro 'cpf', 'email', 'data_de_nascimento', 'nome', 'nome_de_usuario', 'senha' e 'plano_de_assinatura', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Event, "required,oneof=SALE_APPROVED SALE_REFUNDED SUBSCRIPTION_CANCELED SUBSCRIPTION_EXPIRED SUBSCRIPTION_RENEWED"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no nome de evento inválido no parâmetro 'event'", nil)
		erros = append(erros, "Por favor, forneça um evento válido (SALE_APPROVED, SALE_REFUNDED, SUBSCRIPTION_CANCELED, SUBSCRIPTION_EXPIRED ou SUBSCRIPTION_RENEWED) para o parâmetro 'event'.")
	}

	if err := utils.Validate.Var(parametros.Customer.Document, "required,numeric,len=11"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no CPF inválido no parâmetro 'customer.document'", nil)
		erros = append(erros, "Por favor, forneça um CPF válido (texto numérico com 11 dígitos) para o parâmetro 'document'.")
	}

	if err := utils.Validate.Var(parametros.Customer.Name, "required,min=3,max=240"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no nome inválido no parâmetro 'customer.name'", nil)
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 240 caracteres) para o parâmetro 'name'.")
	}

	if err := utils.Validate.Var(parametros.Customer.Email, "required,email"); err != nil {
		utils.DebugLog("KirvanoCreate", "Erro no email inválido no parâmetro 'customer.email'", nil)
		erros = append(erros, "Por favor, forneça um email válido para o parâmetro 'email'.")
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
		err := s.UsuarioModel.Update(nomeDeUsuario, "", "", "", "", "", "Gratuito")

		if err != nil {
			utils.ErroLog("KirvanoCreate", "Erro na atualização do usuário", err)
			return utils.Erro(http.StatusInternalServerError, "Não foi possível atualizar o usuário.")
		}
	}

	return c.JSON(http.StatusCreated, "O usuário da Kirvano foi criado com sucesso.")
}

// UsuarioUpdate godoc
//
// @Summary Atualiza um usuário
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Param   username            path     string true "Nome de Usuário"
//
// @Param   cpf                 body     string false "CPF"
//
// @Param   nome                body     string false "Nome"
//
// @Param   nome_de_usuario     body     string false "Nome de Usuário"
//
// @Param   email               body     string false "Email"
//
// @Param   senha               body     string false "Senha"
//
// @Param   data_de_nascimento  body     string false "Data de Nascimento"
//
// @Param   plano_de_assinatura body     string false "Plano de Assinatura"
//
// @Success 200                 {object} map[string]string
//
// @Failure 400                 {object} echo.HTTPError
//
// @Failure 500                 {object} echo.HTTPError
//
// @Router  /u/:username [patch]
func (s *Server) UsuarioUpdate(c echo.Context) error {
	nomeDeUsuario := c.Param("username")

	type parametrosUpdate struct {
		Cpf               string `json:"cpf"`
		Nome              string `json:"nome"`
		NomeDeUsuario     string `json:"nome_de_usuario"`
		Email             string `json:"email"`
		Senha             string `json:"senha"`
		DataDeNascimento  string `json:"data_de_nascimento"`
		PlanoDeAssinatura string `json:"plano_de_assinatura"`
	}

	parametros := parametrosUpdate{}

	if !utils.IsURLSafe(nomeDeUsuario) {
		utils.DebugLog("UsuarioUpdate", "Erro de validação do nome do usuário no parâmetro 'username'", nil)
		return utils.Erro(http.StatusBadRequest, "Por favor, forneça um nome de usuário válido (3 a 120 caracteres, composto de apenas letras, números e '-' ou '_').")
	}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o CPF, nome, nome de usuário, email, senha, data de nascimento e plano de assinatura do usuário nos parâmetro 'cpf', 'nome', 'nome_de_usuario', 'email', 'senha', 'data_de_nascimento' e 'plano_de_assinatura', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Cpf, "numeric,max=11"); parametros.Cpf != "" && err != nil {
		erros = append(erros, "Por favor, forneça um CPF válido (texto numérico com 11 dígitos) para o parâmetro 'cpf'.")
	}

	if err := utils.Validate.Var(parametros.Nome, "min=3,max=240"); parametros.Nome != "" && err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 240 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.NomeDeUsuario, "min=3,max=120"); parametros.NomeDeUsuario != "" && err != nil || !utils.IsURLSafe(parametros.NomeDeUsuario) {
		erros = append(erros, "Por favor, forneça um nome de usuário válido (texto de 3 a 120 caracteres, contendo apenas letras, número, '_' ou '-') para o parâmetro 'nome_de_usuario'.")
	}

	if err := utils.Validate.Var(parametros.Email, "email"); parametros.Email != "" && err != nil {
		erros = append(erros, "Por favor, forneça um email válido para o parâmetro 'email'.")
	}

	if err := utils.Validate.Var(parametros.Senha, "min=8,max=72"); parametros.Senha != "" && err != nil {
		erros = append(erros, "Por favor, forneça uma senha válida (texto de 8 a 72 caracteres, podendo conter letras, números e símbolos) para o parâmetro 'senha'.")
	}

	if err := utils.Validate.Var(parametros.DataDeNascimento, "datetime=2006-01-02"); parametros.DataDeNascimento != "" && err != nil {
		erros = append(erros, "Por favor, forneça uma data de nascimento válida para o parâmetro 'data_de_nascimento'.")
	}

	if err := utils.Validate.Var(parametros.PlanoDeAssinatura, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido para o parâmetro 'plano_de_assinatura'.")
	}

	if (parametrosUpdate{}) == parametros {
		erros = append(erros, "Por favor, forneça algum valor válido para a atualização.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	if parametros.Senha != "" {
		senhaComHash, err := argon2id.CreateHash(parametros.Senha+utils.Pepper, utils.SenhaParams)

		if err != nil {
			slog.Error("UsuarioUpdate", slog.Any("error", err))
			return utils.ErroCriacaoSenha
		}

		parametros.Senha = senhaComHash
	}

	err := s.UsuarioModel.Update(
		parametros.NomeDeUsuario,
		parametros.Cpf,
		parametros.Nome,
		parametros.Email,
		parametros.Senha,
		parametros.DataDeNascimento,
		parametros.PlanoDeAssinatura,
	)

	if err != nil {
		slog.Error("UsuarioUpdate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, "O usuário foi atualizado com sucesso.")
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
// @Param   senha_nova         body     string true "Senha Nova"
//
// @Param   data_de_nascimento body     string true "Data de Nascimento"
//
// @Success 200                {object} map[string]string
//
// @Failure 400                {object} echo.HTTPError
//
// @Failure 500                {object} echo.HTTPError
//
// @Router  /kirvano/to_user/:hash [post]
func (s *Server) KirvanoToUser(c echo.Context) error {
	valor := c.Param("hash")

	if !utils.IsURLSafe(valor) {
		utils.DebugLog("KirvanoToUser", "Erro na validação do código hash do email de autenticação", nil)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente.")
	}

	parametros := struct {
		SenhaNova        string `json:"senha_nova"`
		DataDeNascimento string `json:"data_de_nascimento"`
	}{}

	if err := c.Bind(&parametros); err != nil {
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroCriacaoSenha
	}

	if err := utils.Validate.Var(parametros.SenhaNova, "required,min=8,max=72"); err != nil {
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroCriacaoSenha
	}

	if err := utils.Validate.Var(parametros.DataDeNascimento, "required,datetime=2006-01-02"); err != nil {
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroCriacaoSenha
	}

	usuario, err := s.EmailAutenticacaoModel.CheckIfValorExistsAndIsValid(valor, "nova_senha")
	fmt.Println(usuario)

	if err != nil {
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	senhaComHash, err := argon2id.CreateHash(parametros.SenhaNova+utils.Pepper, utils.SenhaParams)

	if err != nil {
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroCriacaoSenha
	}

	usuarioTemporario, err := s.UsuarioKirvanoModel.ReadById(usuario)

	if err != nil {
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroBancoDados
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
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.UsuarioKirvanoModel.Remove(usuario)

	if err != nil {
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.EmailAutenticacaoModel.Expirar(valor)

	if err != nil {
		slog.Error("KirvanoToUser", slog.Any("error", err))
		return utils.ErroBancoDados
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
// @Param   username          path     string true "Nome de Usuário"
//
// @Success 200               {object} map[string]string
//
// @Failure 400               {object} echo.HTTPError
//
// @Failure 500               {object} echo.HTTPError
//
// @Router  /u/:username/change_password [patch]
func (s *Server) UsuarioSolicitarTrocaDeSenha(c echo.Context) error {
	nomeDeUsuario := c.Param("username")

	if !utils.IsURLSafe(nomeDeUsuario) {
		utils.DebugLog("UsuarioSolicitarTrocaDeSenha", "Erro de validação do nome do usuário no parâmetro 'username'", nil)
		return utils.Erro(http.StatusBadRequest, "Por favor, forneça um nome de usuário válido (3 a 120 caracteres, composto de apenas letras, números e '-' ou '_').")
	}

	var err error

	usuario, err := s.UsuarioModel.ReadByNomeDeUsuario(nomeDeUsuario)

	if err != nil {
		slog.Error("UsuarioSolicitarTrocaDeSenha", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	var valor string
	valorExiste := true

	for valorExiste {
		valor = utils.GeraHashCode(120)

		valorExiste, err = s.EmailAutenticacaoModel.CheckIfValorExists(valor)

		if err != nil {
			slog.Error("UsuarioSolicitarTrocaDeSenha", slog.Any("error", err))
			return utils.ErroBancoDados
		}
	}

	id, err := s.EmailAutenticacaoModel.Create(valor, "senha", usuario.Id)

	if err != nil {
		slog.Error("UsuarioSolicitarTrocaDeSenha", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.email.SendTrocaDeSenha(id, usuario.Nome, valor, usuario.Email)

	if err != nil {
		slog.Error("UsuarioSolicitarTrocaDeSenha", slog.Any("error", err))
		return utils.ErroBancoDados
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
// @Router  /u/change_password/:hash [patch]
func (s *Server) UsuarioTrocaDeSenha(c echo.Context) error {
	valor := c.Param("hash")

	if !utils.IsURLSafe(valor) {
		utils.DebugLog("UsuarioTrocaDeSenha", "Erro na validação do código hash do email de autenticação", nil)
		return utils.Erro(http.StatusBadRequest, "O 'hash' inserido é inválido, por favor insira um 'hash' existente.")
	}

	parametros := struct {
		SenhaNova string `json:"senha_nova"`
	}{}

	if err := c.Bind(&parametros); err != nil {
		slog.Error("UsuarioTrocaDeSenha", slog.Any("error", err))
		return utils.ErroCriacaoSenha
	}

	if err := utils.Validate.Var(parametros.SenhaNova, "required,min=8,max=72"); err != nil {
		slog.Error("UsuarioTrocaDeSenha", slog.Any("error", err))
		return utils.ErroCriacaoSenha
	}

	usuario, err := s.EmailAutenticacaoModel.CheckIfValorExistsAndIsValid(valor, "senha")

	if err != nil {
		slog.Error("UsuarioTrocaDeSenha", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	senhaComHash, err := argon2id.CreateHash(parametros.SenhaNova+utils.Pepper, utils.SenhaParams)

	if err != nil {
		slog.Error("UsuarioTrocaDeSenha", slog.Any("error", err))
		return utils.ErroCriacaoSenha
	}

	err = s.UsuarioModel.TrocaSenha(usuario, senhaComHash)

	if err != nil {
		slog.Error("UsuarioTrocaDeSenha", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.EmailAutenticacaoModel.Expirar(valor)

	if err != nil {
		slog.Error("UsuarioTrocaDeSenha", slog.Any("error", err))
		return utils.ErroBancoDados
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
// @Failure 500               {object} echo.HTTPError
//
// @Router  /u/login [post]
func (s *Server) UsuarioLogin(c echo.Context) error {
	parametros := struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome de usuário do usuário no parâmetro 'nome_de_usuario'.")
	}

	if err := utils.Validate.Var(parametros.Email, "required,email"); parametros.Email != "" && err != nil {
		erros = append(erros, "Por favor, forneça o email do usuário válido no parâmetro 'email'.")
	}

	if err := utils.Validate.Var(parametros.Senha, "required,min=8,max=72"); err != nil {
		erros = append(erros, "Por favor, forneça uma senha válida (texto de 8 a 72 caracteres, podendo conter letras, números e símbolos) para o parâmetro 'senha'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	id, nome, nomeDeUsuario, planoDeAssinatura, senha, err := s.UsuarioModel.Login(parametros.Email)

	if err != nil {
		slog.Error("UsuarioLogin", slog.Any("error", err))
		return utils.ErroLogin
	}

	match, err := argon2id.ComparePasswordAndHash(parametros.Senha+utils.Pepper, senha)

	if !match || err != nil {
		slog.Error("UsuarioLogin", slog.Any("error", err))
		return utils.ErroLogin
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
// @Router  /u/logout [post]
func (s *Server) UsuarioLogout(c echo.Context) error {
	for _, c := range c.Cookies() {
		c.Expires = time.Now()
	}

	return c.JSON(http.StatusOK, "A autenticação do usuário foi expirada com sucesso.")
}
