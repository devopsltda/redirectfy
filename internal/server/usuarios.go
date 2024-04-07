package server

import (
	"log/slog"
	"net/http"
	"strings"
	"unicode"

	"redirectfy/internal/auth"
	"redirectfy/internal/utils"

	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"

	_ "redirectfy/internal/models"
)

type produto struct {
	Name string `json:"name"`
} // @name Produto

func criaNomeDeUsuario(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) || c == '_' || c == '-' {
			sb.WriteRune(c)
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
// @Param   nome_de_usuario path     string  true  "Nome de Usuário"
//
// @Success 200             {object} models.Usuario
//
// @Failure 400             {object} utils.Erro
//
// @Failure 500             {object} utils.Erro
//
// @Router  /usuarios/:nome_de_usuario [get]
func (s *Server) UsuarioReadByNomeDeUsuario(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	usuario, err := s.UsuarioModel.ReadByNomeDeUsuario(nomeDeUsuario)

	if err != nil {
		slog.Error("UsuarioReadByNomeDeUsuario", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, usuario)
}

// UsuarioReadAll godoc
//
// @Summary Retorna todos os usuários
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Success 200             {object} []models.Usuario
//
// @Failure 400             {object} utils.Erro
//
// @Failure 500             {object} utils.Erro
//
// @Router  /usuarios [get]
func (s *Server) UsuarioReadAll(c echo.Context) error {
	usuarios, err := s.UsuarioModel.ReadAll()

	if err != nil {
		slog.Error("UsuarioReadAll", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, usuarios)
}

// UsuarioTemporarioCreate godoc
//
// @Summary Cria um usuário temporário
//
// @Tags    UsuáriosTemporários
//
// @Accept  json
//
// @Produce json
//
// @Param   customer.document   body     string        true "CPF"
//
// @Param   customer.name       body     string        true "Nome"
//
// @Param   customer.email      body     string        true "Email"
//
// @Param   products            body     []produto    true "Plano de Assinatura"
//
// @Success 200                 {object} map[string]string
//
// @Failure 400                 {object} utils.Erro
//
// @Failure 500                 {object} utils.Erro
//
// @Router  /usuarios_temporarios [post]
func (s *Server) UsuarioTemporarioCreate(c echo.Context) error {
	parametros := struct {
		Cpf                string    `json:"customer.document"`
		Nome               string    `json:"customer.name"`
		Email              string    `json:"customer.email"`
		PlanosDeAssinatura []produto `json:"products"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o CPF, email, data de nascimento, nome, nome de usuário, senha e plano de assinatura do usuário nos parâmetro 'cpf', 'email', 'data_de_nascimento', 'nome', 'nome_de_usuario', 'senha' e 'plano_de_assinatura', respectivamente.")
	}

	if err := utils.Validate.Var(parametros.Cpf, "required,numeric,len=11"); err != nil {
		erros = append(erros, "Por favor, forneça um CPF válido (texto numérico com 11 dígitos) para o parâmetro 'cpf'.")
	}

	if err := utils.Validate.Var(parametros.Nome, "required,min=3,max=240"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 240 caracteres) para o parâmetro 'nome'.")
	}

	if err := utils.Validate.Var(parametros.Email, "required,email"); err != nil {
		erros = append(erros, "Por favor, forneça um email válido para o parâmetro 'email'.")
	}

	if err := utils.Validate.Var(parametros.PlanosDeAssinatura[0].Name, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido para o parâmetro 'plano_de_assinatura'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	nomeDeUsuario := criaNomeDeUsuario(parametros.Nome)

	usuarioId, err := s.UsuarioTemporarioModel.Create(
		parametros.Cpf,
		parametros.Nome,
		nomeDeUsuario,
		parametros.Email,
		parametros.PlanosDeAssinatura[0].Name,
	)

	if err != nil {
		slog.Error("UsuarioTemporarioCreate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	var valor string
	valorExiste := true

	for valorExiste {
		valor = utils.GeraHashCode(120)

		valorExiste, err = s.EmailAutenticacaoModel.CheckIfValorExists(valor)

		if err != nil {
			slog.Error("UsuarioTemporarioCreate", slog.Any("error", err))
			return utils.ErroBancoDados
		}
	}

	id, err := s.EmailAutenticacaoModel.Create(valor, "senha", usuarioId)

	if err != nil {
		slog.Error("UsuarioTemporarioCreate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.email.SendValidacao(id, parametros.Nome, valor, parametros.Email)

	if err != nil {
		slog.Error("UsuarioTemporarioCreate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioCriadoComSucesso)
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
// @Param   nome_de_usuario     path     string true "Nome de Usuário"
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
// @Failure 400                 {object} utils.Erro
//
// @Failure 500                 {object} utils.Erro
//
// @Router  /usuarios/:nome_de_usuario [patch]
func (s *Server) UsuarioUpdate(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

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

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
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

	if err := utils.Validate.Var(parametros.NomeDeUsuario, "min=3,max=120"); parametros.NomeDeUsuario != "" && err != nil || !utils.ValidaNomeDeUsuario(parametros.NomeDeUsuario) {
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
		parametros.Cpf,
		parametros.Nome,
		parametros.NomeDeUsuario,
		parametros.Email,
		parametros.Senha,
		parametros.DataDeNascimento,
		parametros.PlanoDeAssinatura,
	)

	if err != nil {
		slog.Error("UsuarioUpdate", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioAtualizadoComSucesso)
}

// UsuarioAutenticado godoc
//
// @Summary Autentica um usuário
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Param   valor              path     string true "Valor"
//
// @Param   senha              body     string true "Senha"
//
// @Param   data_de_nascimento body     string true "Data de Nascimento"
//
// @Success 200                {object} map[string]string
//
// @Failure 400                {object} utils.Erro
//
// @Failure 500                {object} utils.Erro
//
// @Router  /usuarios/autentica/:valor [patch]
func (s *Server) UsuarioAutenticado(c echo.Context) error {
	type parametrosUpdate struct {
		Senha            string `json:"senha"`
		DataDeNascimento string `json:"data_de_nascimento"`
	}

	parametros := parametrosUpdate{}

	valor := c.Param("valor")

	if !utils.ValidaNomeDeUsuario(valor) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	var erros []string

	if err := utils.Validate.Var(parametros.Senha, "required,min=8,max=72"); parametros.Senha != "" && err != nil {
		erros = append(erros, "Por favor, forneça uma senha válida (texto de 8 a 72 caracteres, podendo conter letras, números e símbolos) para o parâmetro 'senha'.")
	}

	if err := utils.Validate.Var(parametros.DataDeNascimento, "required,datetime=2006-01-02"); parametros.DataDeNascimento != "" && err != nil {
		erros = append(erros, "Por favor, forneça uma data de nascimento válida para o parâmetro 'data_de_nascimento'.")
	}

	if (parametrosUpdate{}) == parametros {
		erros = append(erros, "Por favor, forneça algum valor válido para a atualização.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	senha, err := argon2id.CreateHash(parametros.Senha+utils.Pepper, utils.SenhaParams)

	if err != nil {
		slog.Error("UsuarioCreate", slog.Any("error", err))
		return utils.ErroCriacaoSenha
	}

	usuarioTemporarioId, err := s.EmailAutenticacaoModel.CheckIfValorExistsAndIsValid(valor, "senha")

	if err != nil {
		slog.Error("UsuarioAutenticado", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	usuarioTemporario, err := s.UsuarioTemporarioModel.ReadById(usuarioTemporarioId)

	if err != nil {
		slog.Error("UsuarioAutenticado", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	_, err = s.UsuarioModel.Create(
		usuarioTemporario.Cpf,
		usuarioTemporario.Nome,
		usuarioTemporario.NomeDeUsuario,
		usuarioTemporario.Email,
		senha,
		parametros.DataDeNascimento,
		usuarioTemporario.PlanoDeAssinatura,
	)

	if err != nil {
		slog.Error("UsuarioAutenticado", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.EmailAutenticacaoModel.Expirar(valor)

	if err != nil {
		slog.Error("UsuarioAutenticado", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.UsuarioTemporarioModel.Remove(usuarioTemporarioId)

	if err != nil {
		slog.Error("UsuarioAutenticado", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioAutenticadoComSucesso)
}

// UsuarioTrocaDeSenhaExigir godoc
//
// @Summary Exige a troca de senha de um usuário
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Param   nome_de_usuario   path     string true "Nome de Usuário"
//
// @Success 200               {object} map[string]string
//
// @Failure 400               {object} utils.Erro
//
// @Failure 500               {object} utils.Erro
//
// @Router  /usuarios/:nome_de_usuario/troca_de_senha [patch]
func (s *Server) UsuarioTrocaDeSenhaExigir(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	var err error

	usuario, err := s.UsuarioModel.ReadByNomeDeUsuario(nomeDeUsuario)

	if err != nil {
		slog.Error("UsuarioTrocaDeSenhaExigir", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	var valor string
	valorExiste := true

	for valorExiste {
		valor = utils.GeraHashCode(120)

		valorExiste, err = s.EmailAutenticacaoModel.CheckIfValorExists(valor)

		if err != nil {
			slog.Error("UsuarioTrocaDeSenhaExigir", slog.Any("error", err))
			return utils.ErroBancoDados
		}
	}

	id, err := s.EmailAutenticacaoModel.Create(valor, "senha", usuario.Id)

	if err != nil {
		slog.Error("UsuarioTrocaDeSenhaExigir", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	err = s.email.SendTrocaDeSenha(id, usuario.Nome, valor, usuario.Email)

	if err != nil {
		slog.Error("UsuarioTrocaDeSenhaExigir", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioSenhaTrocadaComSucesso)
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
// @Param   valor             path     string true "Valor"
//
// @Param   senha_nova        body     string true "Senha Nova"
//
// @Success 200               {object} map[string]string
//
// @Failure 400               {object} utils.Erro
//
// @Failure 500               {object} utils.Erro
//
// @Router  /usuarios/troca_de_senha/:valor [patch]
func (s *Server) UsuarioTrocaDeSenha(c echo.Context) error {
	valor := c.Param("valor")

	if !utils.ValidaNomeDeUsuario(valor) {
		return utils.ErroValidacaoNomeDeUsuario
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

	return c.JSON(http.StatusOK, utils.MensagemUsuarioSenhaTrocadaComSucesso)
}

// UsuarioRemove godoc
//
// @Summary Remove um usuário
//
// @Tags    Usuários
//
// @Accept  json
//
// @Produce json
//
// @Param   nome_de_usuario   path     string true "Nome de Usuário"
//
// @Success 200               {object} map[string]string
//
// @Failure 400               {object} utils.Erro
//
// @Failure 500               {object} utils.Erro
//
// @Router  /usuarios/:nome_de_usuario [delete]
func (s *Server) UsuarioRemove(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	err := s.UsuarioModel.Remove(nomeDeUsuario)

	if err != nil {
		slog.Error("UsuarioRemove", slog.Any("error", err))
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioRemovidoComSucesso)
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
// @Failure 400               {object} utils.Erro
//
// @Failure 500               {object} utils.Erro
//
// @Router  /usuarios/login [post]
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
		slog.Error("UsuarioLogin", slog.Any("error", err))
		return utils.ErroAssinaturaJWT
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioLogadoComSucesso)
}
