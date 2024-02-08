package server

import (
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
// @Failure 400             {object} echo.HTTPError
// @Failure 500             {object} echo.HTTPError
// @Router  /api/usuario/:nome_de_usuario [get]
func (s *Server) UsuarioReadByNomeDeUsuario(c echo.Context) error {
	var usuario model.Usuario

	row := s.db.QueryRow(
		"SELECT * FROM USUARIO WHERE REMOVIDO_EM IS NULL AND NOME_DE_USUARIO = $1",
		c.Param("nome_de_usuario"),
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
		return err
	}

	if err := row.Err(); err != nil {
		log.Printf("UsuarioReadByNomeDeUsuario: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, usuario)
}

// UsuarioReadAll godoc
//
// @Summary Retorna todos os usuários
// @Tags    Usuário
// @Accept  json
// @Produce json
// @Success 200             {object} []model.Usuario
// @Failure 400             {object} echo.HTTPError
// @Failure 500             {object} echo.HTTPError
// @Router  /api/usuario [get]
func (s *Server) UsuarioReadAll(c echo.Context) error {
	var usuarios []model.Usuario

	rows, err := s.db.Query("SELECT * FROM USUARIO WHERE REMOVIDO_EM IS NULL")

	if err != nil {
		log.Printf("UsuarioReadAll: %v", err)
		return err
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
			return err
		}

		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		log.Printf("UsuarioReadAll: %v", err)
		return err
	}

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
// @Failure 400                 {object} echo.HTTPError
// @Failure 500                 {object} echo.HTTPError
// @Router  /api/usuario [post]
func (s *Server) UsuarioCreate(c echo.Context) error {
	parametros := struct {
		Cpf               string `json:"cpf"`
		Nome              string `json:"nome"`
		NomeDeUsuario     string `json:"nome_de_usuario"`
		Email             string `json:"email"`
		Senha             string `json:"senha"`
		DataDeNascimento  string `json:"data_de_nascimento"`
		PlanoDeAssinatura int64  `json:"plano_de_assinatura"`
	}{}

	if err := c.Bind(&parametros); err != nil {
		log.Printf("UsuarioNew: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Mensagem": "Requisição teve algum erro",
		})
	}

	senhaComHash, err := geraHashSenha(parametros.Senha)

	if err != nil {
		log.Printf("UsuarioCreate: %v", err)
		return err
	}

	parametros.Senha = senhaComHash

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
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Usuário adicionado com sucesso",
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
// @Failure 400                 {object} echo.HTTPError
// @Failure 500                 {object} echo.HTTPError
// @Router  /api/usuario/:nome_de_usuario [patch]
func (s *Server) UsuarioUpdate(c echo.Context) error {
	parametros := struct {
		Cpf               string `json:"cpf"`
		Nome              string `json:"nome"`
		NomeDeUsuario     string `json:"nome_de_usuario"`
		Email             string `json:"email"`
		Senha             string `json:"senha"`
		DataDeNascimento  string `json:"data_de_nascimento"`
		PlanoDeAssinatura int64  `json:"plano_de_assinatura"`
	}{}

	if err := c.Bind(&parametros); err != nil {
		return err
	}

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
			return err
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
		c.Param("nome_de_usuario"),
	)

	if err != nil {
		log.Printf("UsuarioUpdate: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Usuário atualizado com sucesso",
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
// @Failure 400               {object} echo.HTTPError
// @Failure 500               {object} echo.HTTPError
// @Router  /api/usuario/:nome_de_usuario [delete]
func (s *Server) UsuarioRemove(c echo.Context) error {
	_, err := s.db.Exec(
		"UPDATE USUARIO SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE NOME_DE_USUARIO = $1",
		c.Param("nome_de_usuario"),
	)

	if err != nil {
		log.Printf("UsuarioRemove: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Usuário removido com sucesso",
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
// @Failure 400               {object} echo.HTTPError
// @Failure 500               {object} echo.HTTPError
// @Router  /api/usuario/login [post]
func (s *Server) UsuarioLogin(c echo.Context) error {
	parametros := struct {
		NomeDeUsuario string `json:"nome_de_usuario"`
		Email         string `json:"email"`
		Senha         string `json:"senha"`
	}{}

	if err := c.Bind(&parametros); err != nil {
		return err
	}

	row := s.db.QueryRow(
		"SELECT NOME_DE_USUARIO, SENHA FROM USUARIO WHERE REMOVIDO_EM IS NULL AND (NOME_DE_USUARIO = $1 OR EMAIL = $2)",
		parametros.NomeDeUsuario,
		parametros.Email,
	)

	var nomeDeUsuario string
	var senha string

	if err := row.Scan(&nomeDeUsuario, &senha); err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return err
	}

	if err := row.Err(); err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return err
	}

	err := bcrypt.CompareHashAndPassword([]byte(senha), []byte(parametros.Senha))

	if err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return err
	}

	err = auth.GeraTokensESetaCookies(nomeDeUsuario, c)

	if err != nil {
		log.Printf("UsuarioLogin: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, "")
}
