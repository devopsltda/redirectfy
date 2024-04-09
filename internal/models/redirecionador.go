package models

import (
	"database/sql"
)

type Redirecionador struct {
	Id                      int64          `json:"id"`
	Nome                    string         `json:"nome"`
	CodigoHash              string         `json:"codigo_hash"`
	OrdemDeRedirecionamento string         `json:"ordem_de_redirecionamento"`
	Usuario                 string         `json:"usuario"`
	CriadoEm                string         `json:"criado_em"`
	AtualizadoEm            string         `json:"atualizado_em"`
	RemovidoEm              sql.NullString `json:"removido_em" swaggertype:"string"`
} // @name Redirecionador

type RedirecionadorModel struct {
	DB *sql.DB
}

func (r *RedirecionadorModel) ReadByCodigoHash(codigoHash string) (Redirecionador, error) {
	var redirecionador Redirecionador

	row := r.DB.QueryRow(
		"SELECT ID, NOME, CODIGO_HASH, ORDEM_DE_REDIRECIONAMENTO, USUARIO, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM FROM REDIRECIONADOR WHERE REMOVIDO_EM IS NULL AND CODIGO_HASH = ?",
		codigoHash,
	)

	if err := row.Scan(
		&redirecionador.Id,
		&redirecionador.Nome,
		&redirecionador.CodigoHash,
		&redirecionador.OrdemDeRedirecionamento,
		&redirecionador.Usuario,
		&redirecionador.CriadoEm,
		&redirecionador.AtualizadoEm,
		&redirecionador.RemovidoEm,
	); err != nil {
		return redirecionador, err
	}

	if err := row.Err(); err != nil {
		return redirecionador, err
	}

	return redirecionador, nil
}

func (r *RedirecionadorModel) ReadAll(nomeDeUsuario string) ([]Redirecionador, error) {
	var redirecionadores []Redirecionador

	rows, err := r.DB.Query(
		"SELECT ID, NOME, CODIGO_HASH, ORDEM_DE_REDIRECIONAMENTO, USUARIO, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM FROM REDIRECIONADOR WHERE REMOVIDO_EM IS NULL AND USUARIO = ?",
		nomeDeUsuario,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var redirecionador Redirecionador

		if err := rows.Scan(
			&redirecionador.Id,
			&redirecionador.Nome,
			&redirecionador.CodigoHash,
			&redirecionador.OrdemDeRedirecionamento,
			&redirecionador.Usuario,
			&redirecionador.CriadoEm,
			&redirecionador.AtualizadoEm,
			&redirecionador.RemovidoEm,
		); err != nil {
			return nil, err
		}

		redirecionadores = append(redirecionadores, redirecionador)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return redirecionadores, nil
}

func (r *RedirecionadorModel) CheckIfCodigoHashExists(codigoHash string) (bool, error) {
	row := r.DB.QueryRow(
		"SELECT '' FROM REDIRECIONADOR WHERE REMOVIDO_EM IS NULL AND CODIGO_HASH = ?",
		codigoHash,
	)

	if err := row.Scan(); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			return false, err
		}
	}

	if err := row.Err(); err != nil {
		return false, err
	}

	return true, nil
}

func (r *RedirecionadorModel) Create(nome, codigoHash, ordemDeRedirecionamento, usuario string) (int64, error) {
	result, err := r.DB.Exec(
		"INSERT INTO REDIRECIONADOR (NOME, CODIGO_HASH, ORDEM_DE_REDIRECIONAMENTO, USUARIO) VALUES (?, ?, ?, ?) RETURNING ID",
		nome,
		codigoHash,
		ordemDeRedirecionamento,
		usuario,
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

func (r *RedirecionadorModel) WithinLimit(codigoHash string, quantidadeDeLinks int) (bool, error) {
	var quantidadeLinks int
	var limiteLinksMensal int

	row := r.DB.QueryRow(`SELECT (
																 SELECT COUNT(*)
																 FROM LINK
																			INNER JOIN REDIRECIONADOR ON REDIRECIONADOR.USUARIO = USUARIO.NOME
															 ),
															 USUARIO.LIMITE_LINKS_MENSAL
												FROM (
													SELECT USUARIO.NOME,
																 USUARIO.PLANO_DE_ASSINATURA,
																 PLANO_DE_ASSINATURA.LIMITE_LINKS_MENSAL
													FROM USUARIO
															 INNER JOIN REDIRECIONADOR ON REDIRECIONADOR.USUARIO = USUARIO.NOME_DE_USUARIO
															 INNER JOIN PLANO_DE_ASSINATURA ON PLANO_DE_ASSINATURA.NOME = USUARIO.PLANO_DE_ASSINATURA
													WHERE USUARIO.REMOVIDO_EM IS NULL
																AND REDIRECIONADOR.CODIGO_HASH = ?
												) AS USUARIO`,
		codigoHash,
	)

	if err := row.Scan(
		&quantidadeLinks,
		&limiteLinksMensal,
	); err != nil {
		return false, err
	}

	if err := row.Err(); err != nil {
		return false, err
	}

	if quantidadeLinks > limiteLinksMensal {
		return false, nil
	}

	return true, nil
}

func (r *RedirecionadorModel) Rehash(codigoHashAntigo, codigoHashNovo string) error {
	_, err := r.DB.Exec(
		"UPDATE REDIRECIONADOR SET ATUALIZADO_EM = CURRENT_TIMESTAMP, CODIGO_HASH = ? WHERE REMOVIDO_EM IS NULL AND CODIGO_HASH = ?",
		codigoHashNovo,
		codigoHashAntigo,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RedirecionadorModel) Update(nome, codigoHash, ordemDeRedirecionamento string) error {
	sqlQuery := "UPDATE REDIRECIONADOR SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if nome != "" {
		sqlQuery += ", NOME = '" + nome + "'"
	}

	if ordemDeRedirecionamento != "" {
		sqlQuery += ", ORDEM_DE_REDIRECIONAMENTO = '" + ordemDeRedirecionamento + "'"
	}

	sqlQuery += " WHERE REMOVIDO_EM IS NULL AND CODIGO_HASH = ?"

	_, err := r.DB.Exec(
		sqlQuery,
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RedirecionadorModel) Remove(codigoHash string) error {
	_, err := r.DB.Exec(
		"UPDATE REDIRECIONADOR SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE CODIGO_HASH = ?",
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}
