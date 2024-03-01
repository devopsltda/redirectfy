package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TheDevOpsCorp/redirectify/internal/utils"
)

type EmailAutenticacao struct {
	Id       int64  `json:"id"`
	Valor    string `json:"valor"`
	Tipo     string `json:"tipo"`
	ExpiraEm string `json:"expira_em"`
	Usuario  int64  `json:"usuario"`
} // @name EmailAutenticacao

func EmailAutenticacaoReadByValor(db *sql.DB, valor string) (EmailAutenticacao, error) {
	var emailAutenticacao EmailAutenticacao

	row := db.QueryRow(
		"SELECT ID, VALOR, TIPO, EXPIRA_EM, USUARIO FROM EMAIL_AUTENTICACAO WHERE VALOR = $1",
		valor,
	)

	if err := row.Scan(
		&emailAutenticacao.Id,
		&emailAutenticacao.Valor,
		&emailAutenticacao.Tipo,
		&emailAutenticacao.ExpiraEm,
		&emailAutenticacao.Usuario,
	); err != nil {
		return emailAutenticacao, err
	}

	if err := row.Err(); err != nil {
		return emailAutenticacao, err
	}

	return emailAutenticacao, nil
}

func EmailAutenticacaoCreate(db *sql.DB, valor, tipo string, usuario int64) error {
	_, err := db.Exec(
		"INSERT INTO EMAIL_AUTENTICACAO (VALOR, TIPO, EXPIRA_EM, USUARIO) VALUES ($1, $2, $3, $4)",
		valor,
		tipo,
		time.Now().Add(time.Minute*time.Duration(utils.TempoExpiracao)).Format("2006-01-02 15:04:05"),
		usuario,
	)

	if err != nil {
		return err
	}

	return nil
}

func EmailAutenticacaoCheckIfValorExists(db *sql.DB, valor string) (bool, error) {
	row := db.QueryRow(
		"SELECT '' FROM EMAIL_AUTENTICACAO WHERE VALOR = $1",
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

func EmailAutenticacaoCheckIfValorExistsAndIsValid(db *sql.DB, valor, tipo string) (int64, error) {
	var tipoRetornado string
	var usuario int64

	row := db.QueryRow(
		"SELECT TIPO, USUARIO FROM EMAIL_AUTENTICACAO WHERE VALOR = $1 AND EXPIRA_EM > CURRENT_TIMESTAMP",
		valor,
	)

	if err := row.Scan(&tipoRetornado, &usuario); err != nil {
		if err == sql.ErrNoRows {
			return 0, err
		} else {
			return 0, err
		}
	}

	if err := row.Err(); err != nil {
		return 0, err
	}

	if tipoRetornado != tipo {
		return 0, fmt.Errorf("Tipo do email de retornado não é '%s'.", tipo)
	}

	return usuario, nil
}

func EmailAutenticacaoExpirar(db *sql.DB, valor string) error {
	_, err := db.Exec(
		"UPDATE EMAIL_AUTENTICACAO SET EXPIRA_EM = CURRENT_TIMESTAMP WHERE VALOR = $1",
		valor,
	)

	if err != nil {
		return err
	}

	return nil
}
