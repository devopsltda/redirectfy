package models

import (
	"database/sql"
	"fmt"
	"time"

	"redirectfy/internal/utils"
)

type EmailAutenticacao struct {
	Id             int64  `json:"-"`
	Valor          string `json:"-"`
	Tipo           string `json:"-"`
	ExpiraEm       string `json:"-"`
	UsuarioKirvano int64  `json:"-"`
} // @name EmailAutenticacao

type EmailAutenticacaoModel struct {
	DB *sql.DB
}

func (ea *EmailAutenticacaoModel) Create(valor, tipo string, usuarioKirvano int64) (int64, error) {
	result, err := ea.DB.Exec(
		"INSERT INTO EMAIL_AUTENTICACAO (VALOR, TIPO, EXPIRA_EM, USUARIO_KIRVANO) VALUES (?, ?, ?, ?) RETURNING ID",
		valor,
		tipo,
		time.Now().In(time.FixedZone("GMT", 0)).Add(time.Minute*time.Duration(utils.TempoExpiracao)).Format("2006-01-02 03:04:05"),
		usuarioKirvano,
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

func (ea *EmailAutenticacaoModel) CheckIfValorExists(valor string) (bool, error) {
	row := ea.DB.QueryRow(
		"SELECT '' FROM EMAIL_AUTENTICACAO WHERE VALOR = ?",
		valor,
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

func (ea *EmailAutenticacaoModel) CheckIfValorExistsAndIsValid(valor, tipo string) (int64, error) {
	var tipoRetornado string
	var usuarioKirvano int64

	row := ea.DB.QueryRow(
		"SELECT TIPO, USUARIO_KIRVANO FROM EMAIL_AUTENTICACAO WHERE VALOR = ?",
		valor,
	)

	if err := row.Scan(&tipoRetornado, &usuarioKirvano); err != nil {
		return 0, err
	}

	if err := row.Err(); err != nil {
		return 0, err
	}

	if tipoRetornado != tipo {
		return 0, fmt.Errorf("Tipo do email de retornado não é '%s'.", tipo)
	}

	return usuarioKirvano, nil
}

func (ea *EmailAutenticacaoModel) Expirar(valor string) error {
	_, err := ea.DB.Exec(
		"UPDATE EMAIL_AUTENTICACAO SET EXPIRA_EM = CURRENT_TIMESTAMP WHERE VALOR = ?",
		valor,
	)

	if err != nil {
		return err
	}

	return nil
}
