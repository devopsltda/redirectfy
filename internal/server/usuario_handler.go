package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/TheDevOpsCorp/redirect-max/internal/auth"
	"github.com/TheDevOpsCorp/redirect-max/internal/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/joho/godotenv/autoload"
)

func geraHashSenha(senha string) (string, error) {
	senhaBytes := []byte(senha)

	senhaComHash, err := bcrypt.GenerateFromPassword(senhaBytes, bcrypt.MinCost)

	return string(senhaComHash), err
}

// UsuarioReadByNomeDeUsuario godoc
//
// @Summary Retorna o usuário com o nome de usuário fornecido
// @Tags    Usuário
// @Accept  json
// @Produce json
// @Param   nome_de_usuario path     string  true  "Nome de Usuário"
// @Success 200             {object} model.Usuario
// @Failure 400             {object} Erro
// @Failure 500             {object} Erro
// @Router  /api/usuario/:nome_de_usuario [get]
func (s *Server) UsuarioReadByNomeDeUsuario(c echo.Context) error {
	/*** Parâmetros ***/
	nomeDeUsuario := c.Param("nome_de_usuario")
	/*** Parâmetros ***/

	/*** Validação ***/
	if !validaNomeDeUsuario(nomeDeUsuario) {
		return ErroValidacaoNomeDeUsuario
	}
	/*** Validação ***/

	/*** Banco de Dados ***/
	var usuario model.Usuario

	row := s.db.QueryRow(
		"SELECT * FROM USUARIO WHERE REMOVIDO_EM IS NULL AND NOME_DE_USUARIO = $1",
		nomeDeUsuario,
	)

	if err := row.Scan(
		&usuario.Id,
		&usuario.Cpf,
		&usuario.Nome,
		&usuario.NomeDeUsuario,
		&usuario.Email,
		&usuario.Senha,
		&usuario.DataDeNascimento,
		&usuario.PlanoDeAssinatura,
		&usuario.CriadoEm,
		&usuario.AtualizadoEm,
		&usuario.RemovidoEm,
	); err != nil {
		log.Printf("UsuarioReadByNomeDeUsuario: %v", err)
		return ErroConsultaLinhaBancoDados
	}

	if err := row.Err(); err != nil {
		log.Printf("UsuarioReadByNomeDeUsuario: %v", err)
		return ErroRedeOuResultadoBancoDados
	}
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, usuario)
}

// UsuarioReadAll godoc
//
// @Summary Retorna todos os usuários
// @Tags    Usuário
// @Accept  json
// @Produce json
// @Success 200             {object} []model.Usuario
// @Failure 400             {object} Erro
// @Failure 500             {object} Erro
// @Router  /api/usuario [get]
func (s *Server) UsuarioReadAll(c echo.Context) error {
	/*** Parâmetros ***/
	/*** Parâmetros ***/

	/*** Validação ***/
	/*** Validação ***/

	/*** Banco de Dados ***/
	var usuarios []model.Usuario

	rows, err := s.db.Query("SELECT * FROM USUARIO WHERE REMOVIDO_EM IS NULL")

	if err != nil {
		log.Printf("UsuarioReadAll: %v", err)
		return ErroConsultaBancoDados
	}

	defer rows.Close()

	for rows.Next() {
		var usuario model.Usuario

		if err := rows.Scan(
			&usuario.Id,
			&usuario.Cpf,
			&usuario.Nome,
			&usuario.NomeDeUsuario,
			&usuario.Email,
			&usuario.Senha,
			&usuario.DataDeNascimento,
			&usuario.PlanoDeAssinatura,
			&usuario.CriadoEm,
			&usuario.AtualizadoEm,
			&usuario.RemovidoEm,
		); err != nil {
			log.Printf("UsuarioReadAll: %v", err)
			return ErroConsultaLinhaBancoDados
		}

		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		log.Printf("UsuarioReadAll: %v", err)
		return ErroRedeOuResultadoBancoDados
	}
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, usuarios)
}

// UsuarioCreate godoc
//
// @Summary Cria um usuário
// @Tags    Usuário
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
// @Failure 400                 {object} Erro
// @Failure 500                 {object} Erro
// @Router  /api/usuario [post]
func (s *Server) UsuarioCreate(c echo.Context) error {
	/*** Parâmetros ***/
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
	/*** Parâmetros ***/

	/*** Validação ***/
	if !validaNomeDeUsuario(nomeDeUsuario) {
		return ErroValidacaoNomeDeUsuario
	}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o CPF, email, data de nascimento, nome, nome de usuário, senha e plano de assinatura do usuário nos parâmetro 'cpf', 'email', 'data_de_nascimento', 'nome', 'nome_de_usuario', 'senha' e 'plano_de_assinatura', respectivamente.")
	}

	if err := Validate.Var(parametros.Cpf, "required,numeric,len=11"); err != nil {
		erros = append(erros, "Por favor, forneça um CPF válido (texto numérico com 11 dígitos) para o parâmetro 'cpf'.")
	}

	if err := Validate.Var(parametros.Nome, "required,min=3,max=240"); err != nil {
		erros = append(erros, "Por favor, forneça um nome válido (texto de 3 a 240 caracteres) para o parâmetro 'nome'.")
	}

	if err := Validate.Var(parametros.NomeDeUsuario, "required,min=3,max=120"); err != nil {
		erros = append(erros, "Por favor, forneça um nome de usuário válido (texto de 3 a 120 caracteres) para o parâmetro 'nome_de_usuario'.")
	}

	if err := Validate.Var(parametros.Email, "required,email"); err != nil {
		erros = append(erros, "Por favor, forneça um email válido para o parâmetro 'email'.")
	}

	if err := Validate.Var(parametros.Senha, "required,min=8,max=72"); err != nil {
		erros = append(erros, "Por favor, forneça uma senha válida (texto de 8 a 72 caracteres, podendo conter letras, números e símbolos) para o parâmetro 'senha'.")
	}

	if err := Validate.Var(parametros.DataDeNascimento, "required,datetime"); err != nil {
		erros = append(erros, "Por favor, forneça uma data de nascimento válida para o parâmetro 'data_de_nascimento'.")
	}

	if err := Validate.Var(parametros.PlanoDeAssinatura, "required,gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um plano de assinatura válido para o parâmetro 'plano_de_assinatura'.")
	}

	if len(erros) > 0 {
		return ErroValidacaoParametro(erros)
	}

	senhaComHash, err := geraHashSenha(parametros.Senha)

	if err != nil {
		log.Printf("UsuarioCreate: %v", err)
		return ErroCriacaoSenha
	}

	parametros.Senha = senhaComHash
	/*** Validação ***/

	/*** Banco de Dados ***/
	_, err = s.db.Exec(
		"INSERT INTO USUARIO (CPF, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, PLANO_DE_ASSINATURA, REMOVIDO_EM) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		parametros.Cpf,
		parametros.Nome,
		parametros.NomeDeUsuario,
		parametros.Email,
		parametros.Senha,
		parametros.DataDeNascimento,
		parametros.PlanoDeAssinatura,
		nil,
	)

	if err != nil {
		log.Printf("UsuarioCreate: %v", err)
		return ErroExecucaoQueryBancoDados
	}
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O novo usuário foi adicionado com sucesso.",
	})
}

// UsuarioUpdate godoc
//
// @Summary Atualiza um usuário
// @Tags    Usuário
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
// @Failure 400                 {object} Erro
// @Failure 500                 {object} Erro
// @Router  /api/usuario/:nome_de_usuario [patch]
func (s *Server) UsuarioUpdate(c echo.Context) error {
	/*** Parâmetros ***/
	nomeDeUsuario := c.Param("nome_de_usuario")

	parametros := struct {
		Cpf                string `json:"cpf"`
		Nome               string `json:"nome"`
		NomeDeUsuario      string `json:"nome_de_usuario"`
		Email              string `json:"email"`
		Senha              string `json:"senha"`
		DataDeNascimento   string `json:"data_de_nascimento"`
		PlanoDeAssinatura  int64  `json:"plano_de_assinatura"`
	}{}
	/*** Parâmetros ***/

	/*** Validação ***/
	if !validaNomeDeUsuario(nomeDeUsuario) {
		return ErroValidacaoNomeDeUsuario
	}

	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o CPF, nome, nome de usuário, email, senha, data de nascimento e plano de assinatura do usuário nos parâmetro 'cpf', 'nome', 'nome_de_usuario', 'email', 'senha', 'data_de_nascimento' e 'plano_de_assinatura', respectivamente.")
	}

	if err := Validate.Var(parametros.Cpf, "numeric,max=11"); err != nil {
		erros = append(erros, "Por favor, forneça um CPF válido (texto numérico com 11 dígitos) para o parâmetro 'cpf'.")
	}

	if err := Validate.Var(parametros.Email, "email"); err != nil {
		erros = append(erros, "Por favor, forneça um email válido para o parâmetro 'email'.")
	}

	if err := Validate.Var(parametros.DataDeNascimento, "datetime"); err != nil {
		erros = append(erros, "Por favor, forneça uma data de nascimento válida para o parâmetro 'data_de_nascimento'.")
	}

	if err := Validate.Var(parametros.PlanoDeAssinatura, "gte=0"); err != nil {
		erros = append(erros, "Por favor, forneça um plano de assinatura válido para o parâmetro 'plano_de_assinatura'.")
	}

	if len(erros) > 0 {
		return ErroValidacaoParametro(erros)
	}
	/*** Validação ***/

	/*** Banco de Dados ***/
	sqlQuery := "UPDATE USUARIO SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if parametros.Cpf != "" {
		sqlQuery += ", SET CPF = '" + parametros.Cpf + "'"
	}

	if parametros.Nome != "" {
		sqlQuery += ", SET NOME = '" + parametros.Nome + "'"
	}

	if parametros.NomeDeUsuario != "" {
		sqlQuery += ", SET NOME_DE_USUARIO = '" + parametros.NomeDeUsuario + "'"
	}

	if parametros.Email != "" {
		sqlQuery += ", SET EMAIL = '" + parametros.Email + "'"
	}

	if parametros.Senha != "" {
		senhaComHash, err := geraHashSenha(parametros.Senha)

		if err != nil {
			log.Printf("UsuarioUpdate: %v", err)
			return ErroCriacaoSenha
		}

		parametros.Senha = senhaComHash
		sqlQuery += ", SET SENHA = '" + parametros.Senha + "'"
	}

	if parametros.DataDeNascimento != "" {
		sqlQuery += ", SET DATA_DE_NASCIMENTO = '" + parametros.DataDeNascimento + "'"
	}

	if parametros.PlanoDeAssinatura != 0 {
		sqlQuery += ", SET PLANO_DE_ASSINATURA = " + fmt.Sprint(parametros.PlanoDeAssinatura)
	}

	sqlQuery += " WHERE NOME_DE_USUARIO = $1"

	_, err := s.db.Exec(
		sqlQuery,
		nomeDeUsuario,
	)

	if err != nil {
		log.Printf("UsuarioUpdate: %v", err)
		return ErroExecucaoQueryBancoDados
	}
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O usuário foi atualizado com sucesso.",
	})
}

// UsuarioRemove godoc
//
// @Summary Remove um usuário
// @Tags    Usuário
// @Accept  json
// @Produce json
// @Param   nome_de_usuario   path     string true "Nome de Usuário"
// @Success 200               {object} map[string]string
// @Failure 400               {object} Erro
// @Failure 500               {object} Erro
// @Router  /api/usuario/:nome_de_usuario [delete]
func (s *Server) UsuarioRemove(c echo.Context) error {
	/*** Parâmetros ***/
	nomeDeUsuario := c.Param("nome_de_usuario")
	/*** Parâmetros ***/

	/*** Validação ***/
	if !validaNomeDeUsuario(nomeDeUsuario) {
		return ErroValidacaoNomeDeUsuario
	}
	/*** Validação ***/

	/*** Banco de Dados ***/
	_, err := s.db.Exec(
		"UPDATE USUARIO SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE NOME_DE_USUARIO = $1",
		nomeDeUsuario,
	)

	if err != nil {
		log.Printf("UsuarioRemove: %v", err)
		return ErroExecucaoQueryBancoDados
	}
	/*** Banco de Dados ***/

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O usuário foi removido com sucesso.",
	})
}

// UsuarioLogin godoc
//
// @Summary Autentica o usuário
// @Tags    Usuário
// @Accept  json
// @Produce json
// @Param   nome_de_usuario   body     string false "Nome de Usuário"
// @Param   email             body     string false "Email"
// @Param   senha             body     string true  "Senha"
// @Success 200               {object} map[string]string
// @Failure 400               {object} Erro
// @Failure 500               {object} Erro
// @Router  /api/usuario/login [post]
func (s *Server) UsuarioLogin(c echo.Context) error {
	/*** Parâmetros ***/
	parametros := struct {
		NomeDeUsuario string `json:"nome_de_usuario"`
		Email         string `json:"email"`
		Senha         string `json:"senha"`
	}{}
	/*** Parâmetros ***/

	/*** Validação ***/
	var erros []string

	if err := c.Bind(&parametros); err != nil {
		erros = append(erros, "Por favor, forneça o nome de usuário do usuário no parâmetro 'nome_de_usuario'.")
	}

	if err := Validate.Var(parametros.Email, "email"); err != nil {
		erros = append(erros, "Por favor, forneça o email do usuário válido no parâmetro 'email'.")
	}

	if !validaNomeDeUsuario(parametros.NomeDeUsuario) {
		erros = append(erros, "Por favor, forneça um nome de usuário válido no parâmetro 'nome_de_usuario'.")
	}


	if parametros.NomeDeUsuario == "" && parametros.Email == "" {
		erros = append(erros, "Por favor, forneça o nome de usuário ou o email do usuário para realizar o login.")
	}

	if len(erros) > 0 {
		return ErroValidacaoParametro(erros)
	}
	/*** Validação ***/

	/*** Banco de Dados ***/
	var login string
	var loginValue string

	if parametros.Email != "" {
		login = "EMAIL = $1"
		loginValue = parametros.Email
	}

	if parametros.NomeDeUsuario != "" {
		login = "NOME_DE_USUARIO = $1"
		loginValue = parametros.NomeDeUsuario
	}

	row := s.db.QueryRow(
		"SELECT ID, NOME, NOME_DE_USUARIO, SENHA FROM USUARIO WHERE REMOVIDO_EM IS NULL AND " + login,
		loginValue,
	)

	var id int64
	var nome string
	var nomeDeUsuario string
	var senha string

	if err := row.Scan(&id, &nome, &nomeDeUsuario, &senha); err != nil {
		if err == sql.ErrNoRows {
			log.Printf("UsuarioLogin: %v", err)
			return ErroUsuario
		}

		log.Printf("UsuarioLogin: %v", err)
		return ErroConsultaLinhaBancoDados
	}

	if err := row.Err(); err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return ErroRedeOuResultadoBancoDados
	}
	/*** Banco de Dados ***/

	err := bcrypt.CompareHashAndPassword([]byte(senha), []byte(parametros.Senha))

	if err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return ErroSenha
	}

	err = auth.GeraTokensESetaCookies(id, nome, nomeDeUsuario, c)

	if err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return ErroAssinaturaJWT
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "O usuário foi logado com sucesso.",
	})
}
