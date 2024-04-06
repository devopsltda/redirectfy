package models

import (
	"database/sql"
	"fmt"
)

type PlanoDeAssinatura struct {
	Id                            int64          `json:"id"`
	Nome                          string         `json:"nome"`
	ValorMensal                   int64          `json:"valor_mensal"`
	LimiteLinksMensal             int64          `json:"limite_links_mensal"`
	OrdenacaoAleatoriaLinks       bool           `json:"ordenacao_aleatoria_links"`
	CriadoEm                      string         `json:"criado_em"`
	AtualizadoEm                  string         `json:"atualizado_em"`
	RemovidoEm                    sql.NullString `json:"removido_em" swaggertype:"integer"`
} // @name PlanoDeAssinatura

type PlanoDeAssinaturaModel struct {
	DB *sql.DB
}

func (pa *PlanoDeAssinaturaModel) ReadByNome(nome string) (PlanoDeAssinatura, error) {
	var planoDeAssinatura PlanoDeAssinatura

	row := pa.DB.QueryRow(
		"SELECT ID, NOME, VALOR_MENSAL, LIMITE_LINKS_MENSAL, ORDENACAO_ALEATORIA_LINKS, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL AND NOME = $1",
		nome,
	)

	if err := row.Scan(
		&planoDeAssinatura.Id,
		&planoDeAssinatura.Nome,
		&planoDeAssinatura.ValorMensal,
		&planoDeAssinatura.LimiteLinksMensal,
		&planoDeAssinatura.OrdenacaoAleatoriaLinks,
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

func (pa *PlanoDeAssinaturaModel) ReadAll() ([]PlanoDeAssinatura, error) {
	var planosDeAssinatura []PlanoDeAssinatura

	rows, err := pa.DB.Query("SELECT ID, NOME, VALOR_MENSAL, LIMITE_LINKS_MENSAL, ORDENACAO_ALEATORIA_LINKS, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL")

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
			&planoDeAssinatura.LimiteLinksMensal,
			&planoDeAssinatura.OrdenacaoAleatoriaLinks,
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

func (pa *PlanoDeAssinaturaModel) Create(nome string, valorMensal, limiteLinksMensal int64, ordenacaoAleatoriaLinks bool) error {
	_, err := pa.DB.Exec(
		"INSERT INTO PLANO_DE_ASSINATURA (NOME, VALOR_MENSAL, LIMITE_LINKS_MENSAL, ORDENACAO_ALEATORIA_LINKS) VALUES ($1, $2, $3, $4)",
		nome,
		valorMensal,
		limiteLinksMensal,
		ordenacaoAleatoriaLinks,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pa *PlanoDeAssinaturaModel) Update(nomeParam, nome string, valorMensal, limiteLinksMensal int64, ordenacaoAleatoriaLinks bool) error {
	sqlQuery := "UPDATE PLANO_DE_ASSINATURA SET ATUALIZADO_EM = CURRENT_TIMESTAMP"

	if nome != "" {
		sqlQuery += ", NOME = '" + nome + "'"
	}

	if valorMensal != 0 {
		sqlQuery += ", VALOR_MENSAL = " + fmt.Sprint(valorMensal)
	}

	if limiteLinksMensal != 0 {
		sqlQuery += ", LIMITE_LINKS_MENSAL = " + fmt.Sprint(limiteLinksMensal)
	}

	var ordenacao int

	if ordenacaoAleatoriaLinks {
		ordenacao = 1
	}

	sqlQuery += ", ORDENACAO_ALEATORIA_LINKS = " + fmt.Sprint(ordenacao)

	sqlQuery += " WHERE REMOVIDO_EM IS NULL AND NOME = $1"

	_, err := pa.DB.Exec(
		sqlQuery,
		nomeParam,
	)

	if err != nil {
		return err
	}

	return nil
}

func (pa *PlanoDeAssinaturaModel) Remove(nome string) error {
	_, err := pa.DB.Exec(
		"UPDATE PLANO_DE_ASSINATURA SET REMOVIDO_EM = CURRENT_TIMESTAMP WHERE NOME = $1",
		nome,
	)

	if err != nil {
		return err
	}

	return nil
}
