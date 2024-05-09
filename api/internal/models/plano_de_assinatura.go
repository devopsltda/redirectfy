package models

import (
	"database/sql"
)

type PlanoDeAssinatura struct {
	Nome                    string         `json:"nome"`
	Valor                   int64          `json:"valor"`
	LimiteRedirecionadores  int64          `json:"limite_redirecionadores"`
	OrdenacaoAleatoriaLinks bool           `json:"ordenacao_aleatoria_links"`
} // @name PlanoDeAssinatura

type PlanoDeAssinaturaModel struct {
	DB *sql.DB
}

func (pa *PlanoDeAssinaturaModel) ReadByNome(nome string) (PlanoDeAssinatura, error) {
	var planoDeAssinatura PlanoDeAssinatura

	row := pa.DB.QueryRow(
		"SELECT NOME, VALOR, LIMITE_REDIRECIONADORES, ORDENACAO_ALEATORIA_LINKS FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL AND NOME != 'Administrador' && NOME = ?",
		nome,
	)

	if err := row.Scan(
		&planoDeAssinatura.Nome,
		&planoDeAssinatura.Valor,
		&planoDeAssinatura.LimiteRedirecionadores,
		&planoDeAssinatura.OrdenacaoAleatoriaLinks,
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

	rows, err := pa.DB.Query("SELECT NOME, VALOR, LIMITE_REDIRECIONADORES, ORDENACAO_ALEATORIA_LINKS FROM PLANO_DE_ASSINATURA WHERE REMOVIDO_EM IS NULL AND NOME != 'Administrador'")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var planoDeAssinatura PlanoDeAssinatura

		if err := rows.Scan(
			&planoDeAssinatura.Nome,
			&planoDeAssinatura.Valor,
			&planoDeAssinatura.LimiteRedirecionadores,
			&planoDeAssinatura.OrdenacaoAleatoriaLinks,
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
