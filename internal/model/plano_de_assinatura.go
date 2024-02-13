package model

import (
	"database/sql"
	"fmt"
)

type PlanoDeAssinatura struct {
	Id           int64          `json:"id"`
	Nome         string         `json:"nome"`
	ValorMensal  int64          `json:"valor_mensal"`
	CriadoEm     string         `json:"criado_em"`
	AtualizadoEm string         `json:"atualizado_em"`
	RemovidoEm   sql.NullString `json:"removido_em" swaggertype:"integer"`
} // @name PlanoDeAssinatura

func PlanoDeAssinaturaReadByNome(db *sql.DB, nome string) (*PlanoDeAssinatura, error) {
	var planoDeAssinatura *PlanoDeAssinatura

	row := db.QueryRow(
		"SELECT * FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL AND NOME = $1",
		nome,
	)

	if err := row.Scan(
		&planoDeAssinatura.Id,
		&planoDeAssinatura.Nome,
		&planoDeAssinatura.ValorMensal,
		&planoDeAssinatura.CriadoEm,
		&planoDeAssinatura.AtualizadoEm,
		&planoDeAssinatura.RemovidoEm,
	); err != nil {
		return nil, err
	}

	if err := row.Err(); err != nil {
		return nil, err
	}

	return planoDeAssinatura, nil
}

func PlanoDeAssinaturaReadAll(db *sql.DB) ([]PlanoDeAssinatura, error) {
	var planosDeAssinatura []PlanoDeAssinatura

	rows, err := db.Query("SELECT * FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var planoDeAssinatura PlanoDeAssinatura

		if err := rows.Scan(
			&planoDeAssinatura.Id,
			&planoDeAssinatura.Nome,
			&planoDeAssinatura.ValorMensal,
			&planoDeAssinatura.CriadoEm,
			&planoDeAssinatura.AtualizadoEm,
			&planoDeAssinatura.RemovidoEm,
		); err != nil {
			return nil, err
		}

		planosDeAssinatura = append(planosDeAssinatura, planoDeAssinatura)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return planosDeAssinatura, nil
}

func PlanoDeAssinaturaCreate(db *sql.DB, nome string, valorMensal int64) error {
	_, err := db.Exec(
		"INSERT INTO PLANO_DE_ASSINATURA (NOME, VALOR_MENSAL, REMOVIDO_EM) VALUES ($1, $2, NULL)",
		nome,
		valorMensal,
	)

	if err != nil {
		return err
	}

	return nil
}

func PlanoDeAssinaturaUpdate(db *sql.DB, nomeParam, nome string, valorMensal int64) error {
	sqlQuery := "UPDATE PLANO_DE_ASSINATURA SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if nome != "" {
		sqlQuery += ", SET NOME = '" + nome + "'"
	}

	if valorMensal != 0 {
		sqlQuery += ", SET VALOR_MENSAL = " + fmt.Sprint(valorMensal)
	}

	sqlQuery += " WHERE NOME = $1"

	_, err := db.Exec(
		sqlQuery,
		nomeParam,
	)

	if err != nil {
		return err
	}

	return nil
}

func PlanoDeAssinaturaRemove(db *sql.DB, nome string) error {
	_, err := db.Exec(
		"UPDATE PLANO_DE_ASSINATURA SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE NOME = $1",
		nome,
	)

	if err != nil {
		return err
	}

	return nil
}
