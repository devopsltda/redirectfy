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
	Id             int64  `json:"id"`
	Nome           string `json:"nome"`
	Link           string `json:"link"`
	Ativo          bool   `json:"ativo"`
	Plataforma     string `json:"plataforma"`
	Redirecionador string `json:"redirecionador"`
} // @name Link

type LinkModel struct {
	DB *sql.DB
}

func (l *LinkModel) ReadById(id int64, codigoHash string) (Link, error) {
	var link Link

	row := l.DB.QueryRow(
		"SELECT ID, NOME, LINK, ATIVO, PLATAFORMA, REDIRECIONADOR FROM LINK WHERE REMOVIDO_EM IS NULL AND ID = ? AND REDIRECIONADOR = ?",
		id,
		codigoHash,
	)

	if err := row.Scan(
		&link.Id,
		&link.Nome,
		&link.Link,
		&link.Ativo,
		&link.Plataforma,
		&link.Redirecionador,
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
		"SELECT ID, NOME, LINK, ATIVO, PLATAFORMA, REDIRECIONADOR FROM LINK WHERE REMOVIDO_EM IS NULL AND REDIRECIONADOR = ?",
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
			&link.Ativo,
			&link.Plataforma,
			&link.Redirecionador,
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

	sqlQuery += " WHERE REMOVIDO_EM IS NULL AND ID = ? AND REDIRECIONADOR = ?"

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

func (l *LinkModel) Enable(id int64, codigoHash string) error {
	_, err := l.DB.Exec(
		"UPDATE LINK SET ATIVO = 1 WHERE ID = ? AND REDIRECIONADOR = ?",
		id,
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}

func (l *LinkModel) Disable(id int64, codigoHash string) error {
	_, err := l.DB.Exec(
		"UPDATE LINK SET ATIVO = 0 WHERE ID = ? AND REDIRECIONADOR = ?",
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
		"UPDATE LINK SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE ID = ? AND REDIRECIONADOR = ?",
		id,
		codigoHash,
	)

	if err != nil {
		return err
	}

	return nil
}

// Essa função seleciona o link do Whatsapp e do Telegram, caso seja possível,
// ou algum dos dois. Também é possível que nenhum link seja escolhido, caso
// nenhum link seja passado, no entanto, existe um trigger no banco de dados
// que apenas torna isso possível em um caso excepcional ou com alterações
// externas ao banco.
func LinkPicker(links []Link, isPro bool) (picked_links []Link) {
	var (
		linksWhatsapp []Link
		linksTelegram []Link
	)

	for _, link := range links {
		if !link.Ativo {
			continue
		}

		switch link.Plataforma {
		case "whatsapp":
			linksWhatsapp = append(linksWhatsapp, link)
		case "telegram":
			linksTelegram = append(linksTelegram, link)
		}
	}

	// Caso o usuário não seja de algum plano "Pro" ou "Administrador", não há
	// nenhum tipo de randomização de links selecionados. Caso contrário, os
	// dois links são escolhidos de forma aleatória entre os links disponíveis.
	if !isPro {
		if len(linksWhatsapp) > 0 {
			picked_links = append(picked_links, linksWhatsapp[0])
		}

		if len(linksTelegram) > 0 {
			picked_links = append(picked_links, linksTelegram[0])
		}

		return
	}

	switch {
	case len(linksWhatsapp) == 1:
		picked_links = append(picked_links, linksWhatsapp[0])
	case len(linksWhatsapp) > 1:
		picked_links = append(picked_links, linksWhatsapp[rand.IntN(len(linksWhatsapp))])
	}

	switch {
	case len(linksTelegram) == 1:
		picked_links = append(picked_links, linksTelegram[0])
	case len(linksTelegram) > 1:
		picked_links = append(picked_links, linksTelegram[rand.IntN(len(linksTelegram))])
	}

	return
}
