package server

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/TheDevOpsCorp/redirect-max/internal/model"
	"github.com/labstack/echo/v4"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var symbols []byte = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_")

func generateHashCode(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = symbols[seededRand.Intn(len(symbols))]
	}

	return string(b)
}

func (s *Server) LinkReadByCodigoHash(c echo.Context) error {
	var link model.Link

	row := s.db.QueryRow(
		"SELECT * FROM LINK WHERE REMOVIDO_EM IS NULL AND CODIGO_HASH = $1",
		c.Param("codigo_hash"),
	)

	if err := row.Scan(
		&link.Id,
		&link.Nome,
		&link.CodigoHash,
		&link.LinkWhatsapp,
		&link.LinkTelegram,
		&link.OrdemDeRedirecionamento,
		&link.Usuario.Id,
		&link.CriadoEm,
		&link.AtualizadoEm,
		&link.RemovidoEm,
	); err != nil {
		log.Printf("LinkReadByCodigoHash: %v", err)
		return err
	}

	if err := row.Err(); err != nil {
		log.Printf("LinkReadByCodigoHash: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, link)
}

func (s *Server) LinkReadAll(c echo.Context) error {
	var links []model.Link

	rows, err := s.db.Query("SELECT * FROM LINK WHERE REMOVIDO_EM IS NULL")

	if err != nil {
		log.Printf("LinkReadAll: %v", err)
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var link model.Link

		if err := rows.Scan(
			&link.Id,
			&link.Nome,
			&link.CodigoHash,
			&link.LinkWhatsapp,
			&link.LinkTelegram,
			&link.OrdemDeRedirecionamento,
			&link.Usuario.Id,
			&link.CriadoEm,
			&link.AtualizadoEm,
			&link.RemovidoEm,
		); err != nil {
			log.Printf("LinkReadAll: %v", err)
			return err
		}

		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		log.Printf("LinkReadAll: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, links)
}

func (s *Server) LinkCreate(c echo.Context) error {
	var link model.Link

	if err := c.Bind(&link); err != nil {
		log.Printf("LinkCreate: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"Mensagem": "Requisição teve algum erro",
		})
	}

	var codigo_hash string

	for {
		var codigo_hash_banco string
		codigo_hash = generateHashCode(10)

		row := s.db.QueryRow(
			"SELECT CODIGO_HASH FROM LINK WHERE REMOVIDO_EM IS NULL AND CODIGO_HASH = $1",
			codigo_hash,
		)

		if err := row.Scan(&codigo_hash_banco); err != nil {
			if err == sql.ErrNoRows {
				codigo_hash = codigo_hash_banco
				break
			} else {
				log.Printf("LinkCreate: %v", err)
				return err
			}
		}

		if err := row.Err(); err != nil {
			log.Printf("LinkCreate: %v", err)
			return err
		}
	}

	_, err := s.db.Exec(
		"INSERT INTO LINK (NOME, CODIGO_HASH, LINK_WHATSAPP, LINK_TELEGRAM, ORDEM_DE_REDIRECIONAMENTO, USUARIO, REMOVIDO_EM) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		link.Nome,
		codigo_hash,
		link.LinkWhatsapp,
		link.LinkTelegram,
		link.OrdemDeRedirecionamento,
		link.Usuario.Id,
		nil,
	)

	if err != nil {
		log.Printf("LinkCreate: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Link adicionado com sucesso",
	})
}

func (s *Server) LinkUpdate(c echo.Context) error {
	parametros := struct {
		Nome                    string `json:"nome"`
		LinkWhatsapp            string `json:"link_whatsapp"`
		LinkTelegram            string `json:"link_telegram"`
		OrdemDeRedirecionamento string `json:"ordem_de_redirecionamento"`
		Usuario                 int64  `json:"usuario"`
	}{}

	if err := c.Bind(&parametros); err != nil {
		return err
	}

	sqlQuery := "UPDATE LINK SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if parametros.Nome != "" {
		sqlQuery += ", SET NOME = '" + parametros.Nome + "'"
	}

	if parametros.LinkWhatsapp != "" {
		sqlQuery += ", SET LINK_WHATSAPP = '" + parametros.LinkWhatsapp + "'"
	}

	if parametros.LinkTelegram != "" {
		sqlQuery += ", SET LINK_TELEGRAM = '" + parametros.LinkTelegram + "'"
	}

	if parametros.OrdemDeRedirecionamento != "" {
		sqlQuery += ", SET ORDEM_DE_REDIRECIONAMENTO = '" + parametros.OrdemDeRedirecionamento + "'"
	}

	if parametros.Usuario != 0 {
		sqlQuery += ", SET USUARIO = " + fmt.Sprint(parametros.Usuario)
	}

	sqlQuery += " WHERE CODIGO_HASH = $1"

	_, err := s.db.Exec(
		sqlQuery,
		c.Param("codigo_hash"),
	)

	if err != nil {
		log.Printf("LinkUpdate: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Link atualizado com sucesso",
	})
}

func (s *Server) LinkRemove(c echo.Context) error {
	_, err := s.db.Exec(
		"UPDATE LINK SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE CODIGO_HASH = $1",
		c.Param("codigo_hash"),
	)

	if err != nil {
		log.Printf("LinkRemove: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Mensagem": "Link removido com sucesso",
	})
}
