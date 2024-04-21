package models

import (
	"database/sql"
)

type Usuario struct {
	Id                int64  `json:"-"`
	Cpf               string `json:"cpf"`
	Nome              string `json:"nome"`
	NomeDeUsuario     string `json:"nome_de_usuario"`
	Email             string `json:"email"`
	Senha             string `json:"-"`
	DataDeNascimento  string `json:"data_de_nascimento"`
	PlanoDeAssinatura string `json:"plano_de_assinatura"`
} // @name Usuario

type UsuarioModel struct {
	DB *sql.DB
}

func (u *UsuarioModel) ReadByNomeDeUsuario(nomeDeUsuario string) (Usuario, error) {
	var usuario Usuario

	row := u.DB.QueryRow(
		"SELECT ID, CPF, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, PLANO_DE_ASSINATURA FROM USUARIO WHERE REMOVIDO_EM IS NULL AND NOME_DE_USUARIO = ?",
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
	); err != nil {
		return usuario, err
	}

	if err := row.Err(); err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (u *UsuarioModel) ReadAll() ([]Usuario, error) {
	var usuarios []Usuario

	rows, err := u.DB.Query("SELECT ID, CPF, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, PLANO_DE_ASSINATURA FROM USUARIO WHERE REMOVIDO_EM IS NULL")

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

func (u *UsuarioModel) Create(cpf, nome, nomeDeUsuario, email, senha, dataDeNascimento, planoDeAssinatura string) (int64, error) {
	result, err := u.DB.Exec(
		"INSERT INTO USUARIO (CPF, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, PLANO_DE_ASSINATURA) VALUES (?, ?, ?, ?, ?, ?, ?) RETURNING ID",
		cpf,
		nome,
		nomeDeUsuario,
		email,
		senha,
		dataDeNascimento,
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

func (u *UsuarioModel) TrocaSenha(id int64, senha string) error {
	sqlQuery := "UPDATE USUARIO SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if senha != "" {
		sqlQuery += ", SENHA = '" + senha + "'"
	}
	sqlQuery += " WHERE REMOVIDO_EM IS NULL AND ID = ?"

	_, err := u.DB.Exec(
		sqlQuery,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *UsuarioModel) UpdatePlanoDeAssinatura(nomeDeUsuario, planoDeAssinatura string) error {
	_, err := u.DB.Exec("UPDATE USUARIO SET ATUALIZADO_EM = CURRENT_TIMESTAMP, PLANO_DE_ASSINATURA = ? WHERE REMOVIDO_EM IS NULL AND NOME_DE_USUARIO = ?", nomeDeUsuario)

	if err != nil {
		return err
	}

	return nil
}

func (u *UsuarioModel) Remove(nomeDeUsuario string) error {
	_, err := u.DB.Exec(
		"UPDATE USUARIO SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE NOME_DE_USUARIO = ?",
		nomeDeUsuario,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *UsuarioModel) Login(email string) (int64, string, string, string, string, error) {
	row := u.DB.QueryRow(
		"SELECT ID, NOME, NOME_DE_USUARIO, PLANO_DE_ASSINATURA, SENHA FROM USUARIO WHERE REMOVIDO_EM IS NULL AND EMAIL = ?",
		email,
	)

	var idLogado int64
	var nomeLogado string
	var nomeDeUsuarioLogado string
	var planoDeAssinaturaLogado string
	var senha string

	if err := row.Scan(&idLogado, &nomeLogado, &nomeDeUsuarioLogado, &planoDeAssinaturaLogado, &senha); err != nil {
		return 0, "", "", "", "", err
	}

	if err := row.Err(); err != nil {
		return 0, "", "", "", "", err
	}

	return idLogado, nomeLogado, nomeDeUsuarioLogado, planoDeAssinaturaLogado, senha, nil
}
