package model

import (
	"database/sql"
	"fmt"
)

type Usuario struct {
	Id                int64               `json:"id"`
	Cpf               string              `json:"cpf"`
	Nome              string              `json:"nome"`
	NomeDeUsuario     string              `json:"nome_de_usuario"`
	Email             string              `json:"email"`
	Senha             string              `json:"senha"`
	DataDeNascimento  string              `json:"data_de_nascimento"`
	PlanoDeAssinatura int64               `json:"plano_de_assinatura"`
	CriadoEm          string              `json:"criado_em"`
	AtualizadoEm      string              `json:"atualizado_em"`
	RemovidoEm        sql.NullString      `json:"removido_em" swaggertype:"string"`
} // @name Usuario

func UsuarioReadByNomeDeUsuario(db *sql.DB, nomeDeUsuario string) (*Usuario, error) {
	var usuario *Usuario

	row := db.QueryRow(
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
		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}
	
	return usuario, nil
}

func UsuarioReadAll(db *sql.DB) ([]Usuario, error) {
	var usuarios []Usuario

	rows, err := db.Query("SELECT * FROM USUARIO WHERE REMOVIDO_EM IS NULL")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var usuario Usuario

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
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}

func UsuarioCreate(db *sql.DB, cpf, nome, nomeDeUsuario, email, senha, dataDeNascimento string, planoDeAssinatura int64) error {
	_, err := db.Exec(
		"INSERT INTO USUARIO (CPF, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, PLANO_DE_ASSINATURA, REMOVIDO_EM) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		cpf,
		nome,
		nomeDeUsuario,
		email,
		senha,
		dataDeNascimento,
		planoDeAssinatura,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func UsuarioUpdate(db *sql.DB, cpf, nome, nomeDeUsuario, email, senha, dataDeNascimento string, planoDeAssinatura int64) error {
	sqlQuery := "UPDATE USUARIO SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if cpf != "" {
		sqlQuery += ", SET CPF = '" + cpf + "'"
	}

	if nome != "" {
		sqlQuery += ", SET NOME = '" + nome + "'"
	}

	if nomeDeUsuario != "" {
		sqlQuery += ", SET NOME_DE_USUARIO = '" + nomeDeUsuario + "'"
	}

	if email != "" {
		sqlQuery += ", SET EMAIL = '" + email + "'"
	}

	if senha != "" {
		sqlQuery += ", SET SENHA = '" + senha + "'"
	}

	if dataDeNascimento != "" {
		sqlQuery += ", SET DATA_DE_NASCIMENTO = '" + dataDeNascimento + "'"
	}

	if planoDeAssinatura != 0 {
		sqlQuery += ", SET PLANO_DE_ASSINATURA = " + fmt.Sprint(planoDeAssinatura)
	}

	sqlQuery += " WHERE REMOVIDO_EM IS NULL AND NOME_DE_USUARIO = $1"

	_, err := db.Exec(
		sqlQuery,
		nomeDeUsuario,
	)

	if err != nil {
		return err
	}

	return nil
}

func UsuarioRemove(db *sql.DB, nomeDeUsuario string) error {
	_, err := db.Exec(
		"UPDATE USUARIO SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE NOME_DE_USUARIO = $1",
		nomeDeUsuario,
	)

	if err != nil {
		return err
	}

	return nil
}

func UsuarioLogin(db *sql.DB, email, nomeDeUsuario string) (*int64, *string, *string, *string, error) {
	var login string
	var loginValue string

	if email != "" {
		login = "EMAIL = $1"
		loginValue = email
	}

	if nomeDeUsuario != "" {
		login = "NOME_DE_USUARIO = $1"
		loginValue = nomeDeUsuario
	}

	row := db.QueryRow(
		"SELECT ID, NOME, NOME_DE_USUARIO, SENHA FROM USUARIO WHERE REMOVIDO_EM IS NULL AND " + login,
		loginValue,
	)

	var idLogado int64
	var nomeLogado string
	var nomeDeUsuarioLogado string
	var senha string

	if err := row.Scan(&idLogado, &nomeLogado, &nomeDeUsuarioLogado, &senha); err != nil {
		return nil, nil, nil, nil, err
	}

	if err := row.Err(); err != nil {
		return nil, nil, nil, nil, err
	}

	return &idLogado, &nomeLogado, &nomeDeUsuarioLogado, &senha, nil
}
