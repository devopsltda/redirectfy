package model

import "database/sql"

type Historico struct {
	Id               int64          `json:"id"`
	Usuario          sql.NullInt64  `json:"usuario" swaggertype:"integer"`
	ValorOriginal    sql.NullString `json:"valor_original" swaggertype:"string"`
	ValorNovo        sql.NullString `json:"valor_novo" swaggertype:"string"`
	TabelaModificada string         `json:"tabela_modificada"`
	CriadoEm         string         `json:"criado_em"`
} // @name Historico

func HistoricoReadAll(db *sql.DB) ([]Historico, error) {
	var historico []Historico

	rows, err := db.Query("SELECT * FROM HISTORICO")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var registro Historico

		if err := rows.Scan(
			&registro.Id,
			&registro.Usuario,
			&registro.ValorOriginal,
			&registro.ValorNovo,
			&registro.TabelaModificada,
			&registro.CriadoEm,
		); err != nil {
			return nil, err
		}

		historico = append(historico, registro)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return historico, nil
}
