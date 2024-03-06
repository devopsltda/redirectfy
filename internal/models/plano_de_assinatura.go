package models

import (
	"database/sql"
	"fmt"
)

type PlanoDeAssinatura struct {
	Id            int64          `json:"id"`
	Nome          string         `json:"nome"`
	ValorMensal   int64          `json:"valor_mensal"`
	Limite        int64          `json:"limite"`
	PeriodoLimite string         `json:"periodo_limite"`
	CriadoEm      string         `json:"criado_em"`
	AtualizadoEm  string         `json:"atualizado_em"`
	RemovidoEm    sql.NullString `json:"removido_em" swaggertype:"integer"`
} // @name PlanoDeAssinatura

func PlanoDeAssinaturaReadByNome(db *sql.DB, nome string) (PlanoDeAssinatura, error) {
	var planoDeAssinatura PlanoDeAssinatura

	row := db.QueryRow(
		"SELECT ID, NOME, VALOR_MENSAL, LIMITE, PERIODO_LIMITE, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL AND NOME = $1",
		nome,
	)

	if err := row.Scan(
		&planoDeAssinatura.Id,
		&planoDeAssinatura.Nome,
		&planoDeAssinatura.ValorMensal,
		&planoDeAssinatura.Limite,
		&planoDeAssinatura.PeriodoLimite,
		&planoDeAssinatura.CriadoEm,
		&planoDeAssinatura.AtualizadoEm,
		&planoDeAssinatura.RemovidoEm,
	); err != nil {
		return planoDeAssinatura, err
	}

	if err := row.Err(); err != nil {
		return planoDeAssinatura, err
	}

	return planoDeAssinatura, nil
}

func PlanoDeAssinaturaReadAll(db *sql.DB) ([]PlanoDeAssinatura, error) {
	var planosDeAssinatura []PlanoDeAssinatura

	rows, err := db.Query("SELECT ID, NOME, VALOR_MENSAL, LIMITE, PERIODO_LIMITE, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL")

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
			&planoDeAssinatura.Limite,
			&planoDeAssinatura.PeriodoLimite,
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

func PlanoDeAssinaturaCreate(db *sql.DB, nome string, valorMensal, limite int64, periodoLimite string) error {
	_, err := db.Exec(
		"INSERT INTO PLANO_DE_ASSINATURA (NOME, VALOR_MENSAL, LIMITE, PERIODO_LIMITE) VALUES ($1, $2, $3, $4)",
		nome,
		valorMensal,
		limite,
		periodoLimite,
	)

	if err != nil {
		return err
	}

	return nil
}

func PlanoDeAssinaturaUpdate(db *sql.DB, nomeParam, nome string, valorMensal, limite int64, periodoLimite string) error {
	sqlQuery := "UPDATE PLANO_DE_ASSINATURA SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if nome != "" {
		sqlQuery += ", NOME = '" + nome + "'"
	}

	if valorMensal != 0 {
		sqlQuery += ", VALOR_MENSAL = " + fmt.Sprint(valorMensal)
	}

	if limite != 0 {
		sqlQuery += ", LIMITE = " + fmt.Sprint(limite)
	}

	if periodoLimite != "" {
		sqlQuery += ", PERIODO_LIMITE = '" + periodoLimite + "'"
	}

	sqlQuery += " WHERE REMOVIDO_EM IS NULL AND NOME = $1"

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
