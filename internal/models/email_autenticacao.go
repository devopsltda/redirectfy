package models

import (
	"database/sql"
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
		"SELECT * FROM EMAIL_AUTENTICACAO WHERE VALOR = $1",
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
