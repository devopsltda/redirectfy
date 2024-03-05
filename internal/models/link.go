package models

import (
	"database/sql"
	"errors"
	"math/rand/v2"
)

type Link struct {
	Id             int64          `json:"id"`
	Nome           string         `json:"nome"`
	Link           string         `json:"link"`
	Plataforma     string         `json:"plataforma"`
	Redirecionador string         `json:"redirecionador"`
	CriadoEm       string         `json:"criado_em"`
	AtualizadoEm   string         `json:"atualizado_em"`
	RemovidoEm     sql.NullString `json:"removido_em" swaggertype:"string"`
} // @name Link

func LinkReadById(db *sql.DB, id int64, codigoHash string) (Link, error) {
	var link Link

	row := db.QueryRow(
		"SELECT ID, NOME, LINK, PLATAFORMA, REDIRECIONADOR, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM FROM LINK WHERE REMOVIDO_EM IS NULL AND ID = $1 AND REDIRECIONADOR = $2",
		id,
		codigoHash,
	)

	if err := row.Scan(
		&link.Id,
		&link.Nome,
		&link.Link,
		&link.Plataforma,
		&link.Redirecionador,
		&link.CriadoEm,
		&link.AtualizadoEm,
		&link.RemovidoEm,
	); err != nil {
		return link, err
	}

	if err := row.Err(); err != nil {
		return link, err
	}

	return link, nil
}

func LinkReadAll(db *sql.DB, codigoHash string) ([]Link, error) {
	var links []Link

	rows, err := db.Query(
		"SELECT ID, NOME, LINK, PLATAFORMA, REDIRECIONADOR, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM FROM LINK WHERE REMOVIDO_EM IS NULL AND REDIRECIONADOR = $1",
		codigoHash,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var link Link

		if err := rows.Scan(
			&link.Id,
			&link.Nome,
			&link.Link,
			&link.Plataforma,
			&link.Redirecionador,
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

func LinkCreate(db *sql.DB, nome, link, plataforma, redirecionador string) error {
	_, err := db.Exec(
		"INSERT INTO LINK (NOME, LINK, PLATAFORMA, REDIRECIONADOR) VALUES ($1, $2, $3, $4)",
		nome,
		link,
		plataforma,
		redirecionador,
	)

	if err != nil {
		return err
	}

	return nil
}

func LinkUpdate(db *sql.DB, id int64, codigoHash, nome, link, plataforma string) error {
	sqlQuery := "UPDATE LINK SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if nome != "" {
		sqlQuery += ", NOME = '" + nome + "'"
	}

	if link != "" {
		sqlQuery += ", LINK = '" + link + "'"
	}

	if plataforma != "" {
		sqlQuery += ", PLATAFORMA = '" + plataforma + "'"
	}

	sqlQuery += " WHERE REMOVIDO_EM IS NULL AND ID = $1 AND REDIRECIONADOR = $2"

	_, err := db.Exec(
		sqlQuery,
		id,
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}

func LinkRemove(db *sql.DB, id int64, codigoHash string) error {
	_, err := db.Exec(
		"UPDATE LINK SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE ID = $1 AND REDIRECIONADOR = $2",
		id,
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}

func LinkPicker(links []Link) (*Link, *Link, error) {
	var (
		linkWhatsapp *Link
		linkTelegram *Link
		linksWhatsapp []Link
		linksTelegram []Link
	)

	for _, link := range links {
		switch link.Plataforma {
		case "whatsapp":
			linksWhatsapp = append(linksWhatsapp, link)
		case "telegram":
			linksTelegram = append(linksTelegram, link)
		default:
			return nil, nil, errors.New("Plataforma não identificada nos links")
		}
	}

	switch {
	case len(linksWhatsapp) == 1:
		linkWhatsapp = &linksWhatsapp[0]
	case len(linksWhatsapp) > 0:
		linkWhatsapp = &linksWhatsapp[rand.IntN(len(linksWhatsapp))]
	}

	switch {
	case len(linksTelegram) == 1:
		linkTelegram = &linksTelegram[0]
	case len(linksTelegram) > 0:
		linkTelegram = &linksTelegram[rand.IntN(len(linksTelegram))]
	}

	return linkWhatsapp, linkTelegram, nil
}
