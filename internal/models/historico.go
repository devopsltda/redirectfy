package models

import "database/sql"

type HistoricoPlanoDeAssinatura struct {
	RowId         int64          `json:"rowid"`
	Id            sql.NullInt64  `json:"id"`
	Nome          sql.NullString `json:"nome"`
	ValorMensal   sql.NullInt64  `json:"valor_mensal"`
	Limite        sql.NullInt64  `json:"limite"`
	PeriodoLimite sql.NullString `json:"periodo_limite"`
	CriadoEm      sql.NullString `json:"criado_em"`
	AtualizadoEm  sql.NullString `json:"atualizado_em"`
	RemovidoEm    sql.NullString `json:"removido_em" swaggertype:"integer"`
	Versao        int64          `json:"versao"`
	Bitmask       int64          `json:"bitmask"`
	CriadoEm_     string         `json:"criado_em_"`
} // @name HistoricoPlanoDeAssinatura

type HistoricoUsuario struct {
	RowId             int64          `json:"rowid"`
	Id                sql.NullInt64  `json:"id"`
	Cpf               sql.NullString `json:"cpf"`
	Nome              sql.NullString `json:"nome"`
	NomeDeUsuario     sql.NullString `json:"nome_de_usuario"`
	Email             sql.NullString `json:"email"`
	Senha             sql.NullString `json:"senha"`
	DataDeNascimento  sql.NullString `json:"data_de_nascimento"`
	Autenticado       sql.NullBool   `json:"autenticado"`
	PlanoDeAssinatura sql.NullInt64  `json:"plano_de_assinatura"`
	CriadoEm          sql.NullString `json:"criado_em"`
	AtualizadoEm      sql.NullString `json:"atualizado_em"`
	RemovidoEm        sql.NullString `json:"removido_em" swaggertype:"string"`
	Versao            int64          `json:"versao"`
	Bitmask           int64          `json:"bitmask"`
	CriadoEm_         string         `json:"criado_em_"`
} // @name HistoricoUsuario

type HistoricoEmailAutenticacao struct {
	RowId     int64          `json:"rowid"`
	Id        sql.NullInt64  `json:"id"`
	Valor     sql.NullString `json:"valor"`
	Tipo      sql.NullString `json:"tipo"`
	ExpiraEm  sql.NullString `json:"expira_em"`
	Usuario   sql.NullString `json:"usuario"`
	Versao    int64          `json:"versao"`
	Bitmask   int64          `json:"bitmask"`
	CriadoEm_ string         `json:"criado_em_"`
} // @name HistoricoEmailAutenticacao

type HistoricoLink struct {
	RowId                   int64          `json:"rowid"`
	Id                      sql.NullInt64  `json:"id"`
	Nome                    sql.NullString `json:"nome"`
	CodigoHash              sql.NullString `json:"codigo_hash"`
	LinkWhatsapp            sql.NullString `json:"link_whatsapp"`
	LinkTelegram            sql.NullString `json:"link_telegram"`
	OrdemDeRedirecionamento sql.NullString `json:"ordem_de_redirecionamento"`
	Usuario                 sql.NullInt64  `json:"usuario"`
	CriadoEm                sql.NullString `json:"criado_em"`
	AtualizadoEm            sql.NullString `json:"atualizado_em"`
	RemovidoEm              sql.NullString `json:"removido_em" swaggertype:"string"`
	Versao                  int64          `json:"versao"`
	Bitmask                 int64          `json:"bitmask"`
	CriadoEm_               string         `json:"criado_em_"`
} // @name HistoricoLink

func HistoricoPlanoDeAssinaturaReadAll(db *sql.DB) ([]HistoricoPlanoDeAssinatura, error) {
	var historico []HistoricoPlanoDeAssinatura

	rows, err := db.Query("SELECT * FROM HISTORICO_PLANO_DE_ASSINATURA")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var registro HistoricoPlanoDeAssinatura

		if err := rows.Scan(
			&registro.RowId,
			&registro.Id,
			&registro.Nome,
			&registro.ValorMensal,
			&registro.Limite,
			&registro.PeriodoLimite,
			&registro.CriadoEm,
			&registro.AtualizadoEm,
			&registro.RemovidoEm,
			&registro.Versao,
			&registro.Bitmask,
			&registro.CriadoEm_,
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

func HistoricoUsuarioReadAll(db *sql.DB) ([]HistoricoUsuario, error) {
	var historico []HistoricoUsuario

	rows, err := db.Query("SELECT * FROM HISTORICO_USUARIO")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var registro HistoricoUsuario

		if err := rows.Scan(
			&registro.RowId,
			&registro.Id,
			&registro.Cpf,
			&registro.Nome,
			&registro.NomeDeUsuario,
			&registro.Email,
			&registro.Senha,
			&registro.DataDeNascimento,
			&registro.Autenticado,
			&registro.PlanoDeAssinatura,
			&registro.CriadoEm,
			&registro.AtualizadoEm,
			&registro.RemovidoEm,
			&registro.Versao,
			&registro.Bitmask,
			&registro.CriadoEm_,
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

func HistoricoEmailAutenticacaoReadAll(db *sql.DB) ([]HistoricoEmailAutenticacao, error) {
	var historico []HistoricoEmailAutenticacao

	rows, err := db.Query("SELECT * FROM HISTORICO_EMAIL_AUTENTICACAO")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var registro HistoricoEmailAutenticacao

		if err := rows.Scan(
			&registro.RowId,
			&registro.Id,
			&registro.Valor,
			&registro.Tipo,
			&registro.ExpiraEm,
			&registro.Usuario,
			&registro.Versao,
			&registro.Bitmask,
			&registro.CriadoEm_,
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
func HistoricoLinkReadAll(db *sql.DB) ([]HistoricoLink, error) {
	var historico []HistoricoLink

	rows, err := db.Query("SELECT * FROM HISTORICO_LINK")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var registro HistoricoLink

		if err := rows.Scan(
			&registro.RowId,
			&registro.Id,
			&registro.Nome,
			&registro.CodigoHash,
			&registro.LinkWhatsapp,
			&registro.LinkTelegram,
			&registro.OrdemDeRedirecionamento,
			&registro.Usuario,
			&registro.CriadoEm,
			&registro.AtualizadoEm,
			&registro.RemovidoEm,
			&registro.Versao,
			&registro.Bitmask,
			&registro.CriadoEm_,
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
