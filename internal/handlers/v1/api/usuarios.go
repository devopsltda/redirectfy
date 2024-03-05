package api

import (
	"log"
	"net/http"

	"redirectify/internal/auth"
	"redirectify/internal/models"
	"redirectify/internal/services/database"
	"redirectify/internal/services/email"
	"redirectify/internal/utils"

	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"
)

// UsuarioReadByNomeDeUsuario godoc
//
// @Summary Retorna o usuário com o nome de usuário fornecido
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Param   nome_de_usuario path     string  true  "Nome de Usuário"
// @Success 200             {object} models.Usuario
// @Failure 400             {object} utils.Erro
// @Failure 500             {object} utils.Erro
// @Router  /v1/api/usuarios/:nome_de_usuario [get]
func UsuarioReadByNomeDeUsuario(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	usuario, err := models.UsuarioReadByNomeDeUsuario(database.Db, nomeDeUsuario)

	if err != nil {
		log.Printf("UsuarioReadByNomeDeUsuario: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, usuario)
}

// UsuarioReadAll godoc
//
// @Summary Retorna todos os usuários
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Success 200             {object} []models.Usuario
// @Failure 400             {object} utils.Erro
// @Failure 500             {object} utils.Erro
// @Router  /v1/api/usuarios [get]
func UsuarioReadAll(c echo.Context) error {
	usuarios, err := models.UsuarioReadAll(database.Db)

	if err != nil {
		log.Printf("UsuarioReadAll: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, usuarios)
}

// UsuarioCreate godoc
//
// @Summary Cria um usuário
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Param   cpf                 body     string true "CPF"
// @Param   nome                body     string true "Nome"
// @Param   nome_de_usuario     body     string true "Nome de Usuário"
// @Param   email               body     string true "Email"
// @Param   senha               body     string true "Senha"
// @Param   data_de_nascimento  body     string true "Data de Nascimento"
// @Param   plano_de_assinatura body     int    true "Plano de Assinatura"
// @Success 200                 {object} map[string]string
// @Failure 400                 {object} utils.Erro
// @Failure 500                 {object} utils.Erro
// @Router  /v1/api/usuarios [post]
func UsuarioCreate(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	parametros := struct {
		Cpf               string `json:"cpf"`
		Nome              string `json:"nome"`
		NomeDeUsuario     string `json:"nome_de_usuario"`
		Email             string `json:"email"`
		Senha             string `json:"senha"`
		DataDeNascimento  string `json:"data_de_nascimento"`
		PlanoDeAssinatura int64  `json:"plano_de_assinatura"`
	}{}

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
	}

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

	if err := utils.Validate.Var(parametros.NomeDeUsuario, "required,min=3,max=120"); err != nil || !utils.ValidaNomeDeUsuario(parametros.NomeDeUsuario) {
		erros = append(erros, "Por favor, forneça um nome de usuário válido (texto de 3 a 120 caracteres, contendo apenas letras, número, '_' ou '-') para o parâmetro 'nome_de_usuario'.")
	}

	if err := utils.Validate.Var(parametros.Email, "required,email"); err != nil {
		erros = append(erros, "Por favor, forneça um email válido para o parâmetro 'email'.")
	}

	if err := utils.Validate.Var(parametros.Senha, "required,min=8,max=72"); err != nil {
		erros = append(erros, "Por favor, forneça uma senha válida (texto de 8 a 72 caracteres, podendo conter letras, números e símbolos) para o parâmetro 'senha'.")
	}

	if err := utils.Validate.Var(parametros.DataDeNascimento, "required,datetime=2006-01-02"); err != nil {
		erros = append(erros, "Por favor, forneça uma data de nascimento válida para o parâmetro 'data_de_nascimento'.")
	}

	if err := utils.Validate.Var(parametros.PlanoDeAssinatura, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um plano de assinatura válido para o parâmetro 'plano_de_assinatura'.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	senhaComHash, err := argon2id.CreateHash(parametros.Senha+utils.Pepper, utils.SenhaParams)

	if err != nil {
		log.Printf("UsuarioCreate: %v", err)
		return utils.ErroCriacaoSenha
	}

	parametros.Senha = senhaComHash

	usuarioId, err := models.UsuarioCreate(
		database.Db,
		parametros.Cpf,
		parametros.Nome,
		parametros.NomeDeUsuario,
		parametros.Email,
		parametros.Senha,
		parametros.DataDeNascimento,
		parametros.PlanoDeAssinatura,
	)

	if err != nil {
		log.Printf("UsuarioCreate: %v", err)
		return utils.ErroBancoDados
	}

	var valor string
	valorExiste := true

	for valorExiste {
		valor = utils.GeraHashCode(120)

		valorExiste, err = models.EmailAutenticacaoCheckIfValorExists(database.Db, valor)

		if err != nil {
			log.Printf("UsuarioCreate: %v", err)
			return utils.ErroBancoDados
		}
	}

	id, err := models.EmailAutenticacaoCreate(database.Db, valor, "validacao", usuarioId)

	if err != nil {
		log.Printf("UsuarioCreate: %v", err)
		return utils.ErroBancoDados
	}

	err = email.SendEmailValidacao(id, parametros.Nome, valor, parametros.Email)

	if err != nil {
		log.Printf("UsuarioCreate: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioCriadoComSucesso)
}

// UsuarioUpdate godoc
//
// @Summary Atualiza um usuário
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Param   nome_de_usuario     path     string true "Nome de Usuário"
// @Param   cpf                 body     string false "CPF"
// @Param   nome                body     string false "Nome"
// @Param   nome_de_usuario     body     string false "Nome de Usuário"
// @Param   email               body     string false "Email"
// @Param   senha               body     string false "Senha"
// @Param   data_de_nascimento  body     string false "Data de Nascimento"
// @Param   plano_de_assinatura body     int    false "Plano de Assinatura"
// @Success 200                 {object} map[string]string
// @Failure 400                 {object} utils.Erro
// @Failure 500                 {object} utils.Erro
// @Router  /v1/api/usuarios/:nome_de_usuario [patch]
func UsuarioUpdate(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	type parametrosUpdate struct {
		Cpf               string `json:"cpf"`
		Nome              string `json:"nome"`
		NomeDeUsuario     string `json:"nome_de_usuario"`
		Email             string `json:"email"`
		Senha             string `json:"senha"`
		DataDeNascimento  string `json:"data_de_nascimento"`
		PlanoDeAssinatura int64  `json:"plano_de_assinatura"`
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

	if err := utils.Validate.Var(parametros.PlanoDeAssinatura, "gte=0"); parametros.PlanoDeAssinatura != 0 && err != nil {
		erros = append(erros, "Por favor, forneça um plano de assinatura válido para o parâmetro 'plano_de_assinatura'.")
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
			log.Printf("UsuarioUpdate: %v", err)
			return utils.ErroCriacaoSenha
		}

		parametros.Senha = senhaComHash
	}

	err := models.UsuarioUpdate(
		database.Db,
		parametros.Cpf,
		parametros.Nome,
		parametros.NomeDeUsuario,
		parametros.Email,
		parametros.Senha,
		parametros.DataDeNascimento,
		parametros.PlanoDeAssinatura,
	)

	if err != nil {
		log.Printf("UsuarioUpdate: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioAtualizadoComSucesso)
}

// UsuarioAutenticado godoc
//
// @Summary Autentica um usuário
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Param   valor             path     string true "Valor"
// @Success 200               {object} map[string]string
// @Failure 400               {object} utils.Erro
// @Failure 500               {object} utils.Erro
// @Router  /v1/api/usuarios/autentica/:valor [patch]
func UsuarioAutenticado(c echo.Context) error {
	valor := c.Param("valor")

	if !utils.ValidaNomeDeUsuario(valor) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	usuario, err := models.EmailAutenticacaoCheckIfValorExistsAndIsValid(database.Db, valor, "validacao")

	if err != nil {
		log.Printf("UsuarioAutenticado: %v", err)
		return utils.ErroBancoDados
	}

	err = models.UsuarioAutenticado(database.Db, usuario)

	if err != nil {
		log.Printf("UsuarioAutenticado: %v", err)
		return utils.ErroBancoDados
	}

	err = models.EmailAutenticacaoExpirar(database.Db, valor)

	if err != nil {
		log.Printf("UsuarioAutenticado: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioAutenticadoComSucesso)
}

// UsuarioTrocaDeSenhaExigir godoc
//
// @Summary Exige a troca de senha de um usuário
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Param   nome_de_usuario   path     string true "Nome de Usuário"
// @Success 200               {object} map[string]string
// @Failure 400               {object} utils.Erro
// @Failure 500               {object} utils.Erro
// @Router  /v1/api/usuarios/:nome_de_usuario/troca_de_senha [patch]
func UsuarioTrocaDeSenhaExigir(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	var err error

	usuario, err := models.UsuarioReadByNomeDeUsuario(database.Db, nomeDeUsuario)

	if err != nil {
		log.Printf("UsuarioTrocaDeSenhaExigir: %v", err)
		return utils.ErroBancoDados
	}

	var valor string
	valorExiste := true

	for valorExiste {
		valor = utils.GeraHashCode(120)

		valorExiste, err = models.EmailAutenticacaoCheckIfValorExists(database.Db, valor)

		if err != nil {
			log.Printf("UsuarioTrocaDeSenhaExigir: %v", err)
			return utils.ErroBancoDados
		}
	}

	id, err := models.EmailAutenticacaoCreate(database.Db, valor, "senha", usuario.Id)

	if err != nil {
		log.Printf("UsuarioTrocaDeSenhaExigir: %v", err)
		return utils.ErroBancoDados
	}

	err = email.SendEmailTrocaDeSenha(id, usuario.Nome, valor, usuario.Email)

	if err != nil {
		log.Printf("UsuarioTrocaDeSenhaExigir: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioSenhaTrocadaComSucesso)
}

// UsuarioTrocaDeSenha godoc
//
// @Summary Troca a senha de um usuário
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Param   valor             path     string true "Valor"
// @Param   senha_nova        body     string true "Senha Nova"
// @Success 200               {object} map[string]string
// @Failure 400               {object} utils.Erro
// @Failure 500               {object} utils.Erro
// @Router  /v1/api/usuarios/troca_de_senha/:valor [patch]
func UsuarioTrocaDeSenha(c echo.Context) error {
	valor := c.Param("valor")

	if !utils.ValidaNomeDeUsuario(valor) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	parametros := struct {
		SenhaNova string `json:"senha_nova"`
	}{}

	if err := c.Bind(&parametros); err != nil {
		log.Printf("UsuarioTrocaDeSenha: %v", err)
		return utils.ErroCriacaoSenha
	}

	if err := utils.Validate.Var(parametros.SenhaNova, "required,min=8,max=72"); err != nil {
		log.Printf("UsuarioTrocaDeSenha: %v", err)
		return utils.ErroCriacaoSenha
	}

	usuario, err := models.EmailAutenticacaoCheckIfValorExistsAndIsValid(database.Db, valor, "senha")

	if err != nil {
		log.Printf("UsuarioTrocaDeSenha: %v", err)
		return utils.ErroBancoDados
	}

	senhaComHash, err := argon2id.CreateHash(parametros.SenhaNova+utils.Pepper, utils.SenhaParams)

	if err != nil {
		log.Printf("UsuarioTrocaDeSenha: %v", err)
		return utils.ErroCriacaoSenha
	}

	err = models.UsuarioTrocaSenha(database.Db, usuario, senhaComHash)

	if err != nil {
		log.Printf("UsuarioTrocaDeSenha: %v", err)
		return utils.ErroBancoDados
	}

	err = models.EmailAutenticacaoExpirar(database.Db, valor)

	if err != nil {
		log.Printf("UsuarioTrocaDeSenha: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioSenhaTrocadaComSucesso)
}

// UsuarioRemove godoc
//
// @Summary Remove um usuário
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Param   nome_de_usuario   path     string true "Nome de Usuário"
// @Success 200               {object} map[string]string
// @Failure 400               {object} utils.Erro
// @Failure 500               {object} utils.Erro
// @Router  /v1/api/usuarios/:nome_de_usuario [delete]
func UsuarioRemove(c echo.Context) error {
	nomeDeUsuario := c.Param("nome_de_usuario")

	if !utils.ValidaNomeDeUsuario(nomeDeUsuario) {
		return utils.ErroValidacaoNomeDeUsuario
	}

	err := models.UsuarioRemove(database.Db, nomeDeUsuario)

	if err != nil {
		log.Printf("UsuarioRemove: %v", err)
		return utils.ErroBancoDados
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioRemovidoComSucesso)
}

// UsuarioLogin godoc
//
// @Summary Autentica o usuário
// @Tags    Usuários
// @Accept  json
// @Produce json
// @Param   nome_de_usuario   body     string false "Nome de Usuário"
// @Param   email             body     string false "Email"
// @Param   senha             body     string true  "Senha"
// @Success 200               {object} map[string]string
// @Failure 400               {object} utils.Erro
// @Failure 500               {object} utils.Erro
// @Router  /v1/api/usuarios/login [post]
func UsuarioLogin(c echo.Context) error {
	parametros := struct {
		NomeDeUsuario string `json:"nome_de_usuario"`
		Email         string `json:"email"`
		Senha         string `json:"senha"`
	}{}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome de usuário do usuário no parâmetro 'nome_de_usuario'.")
	}

	if err := utils.Validate.Var(parametros.Email, "email"); parametros.Email != "" && err != nil {
		erros = append(erros, "Por favor, forneça o email do usuário válido no parâmetro 'email'.")
	}

	if err := utils.Validate.Var(parametros.NomeDeUsuario, "min=3,max=120"); parametros.NomeDeUsuario != "" && err != nil || !utils.ValidaNomeDeUsuario(parametros.NomeDeUsuario) {
		erros = append(erros, "Por favor, forneça um nome de usuário válido (texto de 3 a 120 caracteres, contendo apenas letras, número, '_' ou '-') para o parâmetro 'nome_de_usuario'.")
	}

	if err := utils.Validate.Var(parametros.Senha, "min=8,max=72"); err != nil {
		erros = append(erros, "Por favor, forneça uma senha válida (texto de 8 a 72 caracteres, podendo conter letras, números e símbolos) para o parâmetro 'senha'.")
	}

	if parametros.NomeDeUsuario == "" && parametros.Email == "" {
		erros = append(erros, "Por favor, forneça o nome de usuário ou o email do usuário para realizar o login.")
	}

	if len(erros) > 0 {
		return utils.ErroValidacaoParametro(erros)
	}

	id, nome, nomeDeUsuario, autenticado, senha, err := models.UsuarioLogin(database.Db, parametros.Email, parametros.NomeDeUsuario)

	if err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return utils.ErroLogin
	}

	match, err := argon2id.ComparePasswordAndHash(parametros.Senha+utils.Pepper, senha)

	if !match || err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return utils.ErroLogin
	}

	err = auth.GeraTokensESetaCookies(id, nome, nomeDeUsuario, autenticado, c)

	if err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return utils.ErroAssinaturaJWT
	}

	return c.JSON(http.StatusOK, utils.MensagemUsuarioLogadoComSucesso)
}
