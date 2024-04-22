package models

import (
	"database/sql"
)

type UsuarioKirvano struct {
	Id                int64          `json:"id"`
	Cpf               string         `json:"cpf"`
	Nome              string         `json:"nome"`
	NomeDeUsuario     string         `json:"nome_de_usuario"`
	Email             string         `json:"email"`
	PlanoDeAssinatura string         `json:"plano_de_assinatura"`
	CriadoEm          string         `json:"criado_em"`
} // @name UsuarioKirvano

type UsuarioKirvanoModel struct {
	DB *sql.DB
}

func (u *UsuarioKirvanoModel) ReadById(id int64) (UsuarioKirvano, error) {
	var usuario UsuarioKirvano

	row := u.DB.QueryRow(
		"SELECT ID, CPF, NOME, NOME_DE_USUARIO, EMAIL, PLANO_DE_ASSINATURA, CRIADO_EM FROM USUARIO_KIRVANO WHERE REMOVIDO_EM IS NULL AND ID = ?",
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
	); err != nil {
		return usuario, err
	}

	if err := row.Err(); err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (u *UsuarioKirvanoModel) Create(cpf, nome, nomeDeUsuario, email, planoDeAssinatura string) (int64, error) {
	result, err := u.DB.Exec(
		"INSERT INTO USUARIO_KIRVANO (CPF, NOME, NOME_DE_USUARIO, EMAIL, PLANO_DE_ASSINATURA) VALUES (?, ?, ?, ?, ?) RETURNING ID",
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

func (u *UsuarioKirvanoModel) Remove(id int64) error {
	_, err := u.DB.Exec(
		"UPDATE USUARIO_KIRVANO SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE ID = ?",
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
