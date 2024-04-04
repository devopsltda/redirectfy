package models

import (
	"database/sql"
	"fmt"
	"time"

	"redirectify/internal/utils"
)

type EmailAutenticacao struct {
	Id       int64  `json:"id"`
	Valor    string `json:"valor"`
	Tipo     string `json:"tipo"`
	ExpiraEm string `json:"expira_em"`
	Usuario  int64  `json:"usuario"`
} // @name EmailAutenticacao

type EmailAutenticacaoModel struct {
	DB *sql.DB
}

func (ea *EmailAutenticacaoModel) ReadByValor(valor string) (EmailAutenticacao, error) {
	var emailAutenticacao EmailAutenticacao

	row := ea.DB.QueryRow(
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

func (ea *EmailAutenticacaoModel) Create(valor, tipo string, usuario int64) (int64, error) {
	result, err := ea.DB.Exec(
		"INSERT INTO EMAIL_AUTENTICACAO (VALOR, TIPO, EXPIRA_EM, USUARIO) VALUES ($1, $2, $3, $4) RETURNING ID",
		valor,
		tipo,
		time.Now().In(time.FixedZone("GMT", 0)).Add(time.Minute*time.Duration(utils.TempoExpiracao)).Format("2006-01-02 03:04:05"),
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

func (ea *EmailAutenticacaoModel) CheckIfValorExists(valor string) (bool, error) {
	row := ea.DB.QueryRow(
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

func (ea *EmailAutenticacaoModel) CheckIfValorExistsAndIsValid(valor, tipo string) (int64, error) {
	var tipoRetornado string
	var usuario int64

	row := ea.DB.QueryRow(
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

func (ea *EmailAutenticacaoModel) Expirar(valor string) error {
	_, err := ea.DB.Exec(
		"UPDATE EMAIL_AUTENTICACAO SET EXPIRA_EM = CURRENT_TIMESTAMP WHERE VALOR = $1",
		valor,
	)

	if err != nil {
		return err
	}

	return nil
}
