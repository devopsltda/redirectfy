package models

import (
	"database/sql"
	"fmt"
	"math/rand/v2"
)

type LinkToBatchInsert struct {
	Nome       string `json:"nome"`
	Link       string `json:"link"`
	Plataforma string `json:"plataforma"`
}

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

type LinkModel struct {
	DB *sql.DB
}

func (l *LinkModel) ReadById(id int64, codigoHash string) (Link, error) {
	var link Link

	row := l.DB.QueryRow(
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

func (l *LinkModel) ReadByCodigoHash(codigoHash string) ([]Link, error) {
	var links []Link

	rows, err := l.DB.Query(
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

func (l *LinkModel) Create(redirecionador string, links []LinkToBatchInsert) error {
	var values string

	for _, link := range links[:len(links)-1] {
		values += fmt.Sprintf("('%s', '%s', '%s', '%s'),", link.Nome, link.Link, link.Plataforma, redirecionador)
	}

	lastLink := links[len(links)-1]
	values += fmt.Sprintf("('%s', '%s', '%s', '%s')", lastLink.Nome, lastLink.Link, lastLink.Plataforma, redirecionador)

	_, err := l.DB.Exec(fmt.Sprintf("INSERT INTO LINK (NOME, LINK, PLATAFORMA, REDIRECIONADOR) VALUES %s", values))

	if err != nil {
		return err
	}

	return nil
}

func (l *LinkModel) Update(id int64, codigoHash, nome, link, plataforma string) error {
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

	_, err := l.DB.Exec(
		sqlQuery,
		id,
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}

func (l *LinkModel) Remove(id int64, codigoHash string) error {
	_, err := l.DB.Exec(
		"UPDATE LINK SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE ID = $1 AND REDIRECIONADOR = $2",
		id,
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}

func LinkPicker(links []Link) (picked_links []Link) {
	var (
		linkWhatsapp  Link
		linkTelegram  Link
		linksWhatsapp []Link
		linksTelegram []Link
	)

	for _, link := range links {
		switch link.Plataforma {
		case "whatsapp":
			linksWhatsapp = append(linksWhatsapp, link)
		case "telegram":
			linksTelegram = append(linksTelegram, link)
		}
	}

	switch {
	case len(linksWhatsapp) == 1:
		linkWhatsapp = linksWhatsapp[0]
	case len(linksWhatsapp) > 0:
		linkWhatsapp = linksWhatsapp[rand.IntN(len(linksWhatsapp))]
	}

	switch {
	case len(linksTelegram) == 1:
		linkTelegram = linksTelegram[0]
	case len(linksTelegram) > 0:
		linkTelegram = linksTelegram[rand.IntN(len(linksTelegram))]
	}

	return []Link{linkWhatsapp, linkTelegram}
}
