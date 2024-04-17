package models

import (
	"database/sql"
)

type PlanoDeAssinatura struct {
	Id                      int64          `json:"id"`
	Nome                    string         `json:"nome"`
	Valor                   int64          `json:"valor"`
	LimiteRedirecionadores  int64          `json:"limite_redirecionadores"`
	OrdenacaoAleatoriaLinks bool           `json:"ordenacao_aleatoria_links"`
	CriadoEm                string         `json:"criado_em"`
	AtualizadoEm            string         `json:"atualizado_em"`
} // @name PlanoDeAssinatura

type PlanoDeAssinaturaModel struct {
	DB *sql.DB
}

func (pa *PlanoDeAssinaturaModel) ReadByNome(nome string) (PlanoDeAssinatura, error) {
	var planoDeAssinatura PlanoDeAssinatura

	row := pa.DB.QueryRow(
		"SELECT ID, NOME, VALOR, LIMITE_REDIRECIONADORES, ORDENACAO_ALEATORIA_LINKS, CRIADO_EM, ATUALIZADO_EM FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL AND NOME = ?",
		nome,
	)

	if err := row.Scan(
		&planoDeAssinatura.Id,
		&planoDeAssinatura.Nome,
		&planoDeAssinatura.Valor,
		&planoDeAssinatura.LimiteRedirecionadores,
		&planoDeAssinatura.OrdenacaoAleatoriaLinks,
		&planoDeAssinatura.CriadoEm,
		&planoDeAssinatura.AtualizadoEm,
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

	rows, err := pa.DB.Query("SELECT ID, NOME, VALOR, LIMITE_REDIRECIONADORES, ORDENACAO_ALEATORIA_LINKS, CRIADO_EM, ATUALIZADO_EM FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var planoDeAssinatura PlanoDeAssinatura

		if err := rows.Scan(
			&planoDeAssinatura.Id,
			&planoDeAssinatura.Nome,
			&planoDeAssinatura.Valor,
			&planoDeAssinatura.LimiteRedirecionadores,
			&planoDeAssinatura.OrdenacaoAleatoriaLinks,
			&planoDeAssinatura.CriadoEm,
			&planoDeAssinatura.AtualizadoEm,
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
