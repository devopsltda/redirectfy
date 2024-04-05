package models

import (
	"database/sql"
)

type UsuarioTemporario struct {
	Id                int64          `json:"id"`
	Cpf               string         `json:"cpf"`
	Nome              string         `json:"nome"`
	NomeDeUsuario     string         `json:"nome_de_usuario"`
	Email             string         `json:"email"`
	PlanoDeAssinatura string         `json:"plano_de_assinatura"`
	CriadoEm          string         `json:"criado_em"`
	RemovidoEm        sql.NullString `json:"removido_em" swaggertype:"string"`
} // @name UsuarioTemporario

type UsuarioTemporarioModel struct {
	DB *sql.DB
}

func (u *UsuarioTemporarioModel) ReadById(id int64) (UsuarioTemporario, error) {
	var usuario UsuarioTemporario

	row := u.DB.QueryRow(
		"SELECT ID, CPF, NOME, NOME_DE_USUARIO, EMAIL, PLANO_DE_ASSINATURA, CRIADO_EM, REMOVIDO_EM FROM USUARIO_TEMPORARIO WHERE REMOVIDO_EM IS NULL AND ID = $1",
		id,
	)

	if err := row.Scan(
		&usuario.Id,
		&usuario.Cpf,
		&usuario.Nome,
		&usuario.NomeDeUsuario,
		&usuario.Email,
		&usuario.PlanoDeAssinatura,
		&usuario.CriadoEm,
		&usuario.RemovidoEm,
	); err != nil {
		return usuario, err
	}

	if err := row.Err(); err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (u *UsuarioTemporarioModel) Create(cpf, nome, nomeDeUsuario, email, planoDeAssinatura string) (int64, error) {
	result, err := u.DB.Exec(
		"INSERT INTO USUARIO_TEMPORARIO (CPF, NOME, NOME_DE_USUARIO, EMAIL, PLANO_DE_ASSINATURA) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING ID",
		cpf,
		nome,
		nomeDeUsuario,
		email,
		planoDeAssinatura,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UsuarioTemporarioModel) Remove(id int64) error {
	_, err := u.DB.Exec(
		"UPDATE USUARIO_TEMPORARIO SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE ID = $1",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
