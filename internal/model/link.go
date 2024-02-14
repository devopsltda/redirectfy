package model

import (
	"database/sql"
)

type Link struct {
	Id                      int64          `json:"id"`
	Nome                    string         `json:"nome"`
	CodigoHash              string         `json:"codigo_hash"`
	LinkWhatsapp            string         `json:"link_whatsapp"`
	LinkTelegram            string         `json:"link_telegram"`
	OrdemDeRedirecionamento string         `json:"ordem_de_redirecionamento"`
	Usuario                 int64          `json:"usuario"`
	CriadoEm                string         `json:"criado_em"`
	AtualizadoEm            string         `json:"atualizado_em"`
	RemovidoEm              sql.NullString `json:"removido_em" swaggertype:"string"`
} // @name Link

func LinkReadByCodigoHash(db *sql.DB, codigoHash string) (*Link, error) {
	var link *Link

	row := db.QueryRow(
		"SELECT * FROM LINK WHERE REMOVIDO_EM IS NULL AND CODIGO_HASH = $1",
		codigoHash,
	)

	if err := row.Scan(
		&link.Id,
		&link.Nome,
		&link.CodigoHash,
		&link.LinkWhatsapp,
		&link.LinkTelegram,
		&link.OrdemDeRedirecionamento,
		&link.Usuario,
		&link.CriadoEm,
		&link.AtualizadoEm,
		&link.RemovidoEm,
	); err != nil {
		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return link, nil
}

func LinkReadAll(db *sql.DB) ([]Link, error) {
	var links []Link

	rows, err := db.Query("SELECT * FROM LINK WHERE REMOVIDO_EM IS NULL")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var link Link

		if err := rows.Scan(
			&link.Id,
			&link.Nome,
			&link.CodigoHash,
			&link.LinkWhatsapp,
			&link.LinkTelegram,
			&link.OrdemDeRedirecionamento,
			&link.Usuario,
			&link.CriadoEm,
			&link.AtualizadoEm,
			&link.RemovidoEm,
		); err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}

func LinkCheckIfCodigoHashExists(db *sql.DB, codigoHash string) (bool, error) {
	row := db.QueryRow(
		"SELECT '' FROM LINK WHERE REMOVIDO_EM IS NULL AND CODIGO_HASH = $1",
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

func LinkCreate(db *sql.DB, nome, codigoHash, linkWhatsapp, linkTelegram, ordemDeRedirecionamento string, usuario int64) error {
	_, err := db.Exec(
		"INSERT INTO LINK (NOME, CODIGO_HASH, LINK_WHATSAPP, LINK_TELEGRAM, ORDEM_DE_REDIRECIONAMENTO, USUARIO, REMOVIDO_EM) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		nome,
		codigoHash,
		linkWhatsapp,
		linkTelegram,
		ordemDeRedirecionamento,
		usuario,
		nil,
	)

	if err != nil {
		return err
	}

	return nil
}

func LinkUpdate(db *sql.DB, nome, codigoHash, linkWhatsapp, linkTelegram, ordemDeRedirecionamento string) error {
	sqlQuery := "UPDATE LINK SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if nome != "" {
		sqlQuery += ", SET NOME = '" + nome + "'"
	}

	if linkWhatsapp != "" {
		sqlQuery += ", SET LINK_WHATSAPP = '" + linkWhatsapp + "'"
	}

	if linkTelegram != "" {
		sqlQuery += ", SET LINK_TELEGRAM = '" + linkTelegram + "'"
	}

	if ordemDeRedirecionamento != "" {
		sqlQuery += ", SET ORDEM_DE_REDIRECIONAMENTO = '" + ordemDeRedirecionamento + "'"
	}

	sqlQuery += " WHERE CODIGO_HASH = $1"

	_, err := db.Exec(
		sqlQuery,
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}

func LinkRemove(db *sql.DB, codigoHash string) error {
	_, err := db.Exec(
		"UPDATE LINK SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE CODIGO_HASH = $1",
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}
