package models

import "database/sql"

type HistoricoPlanoDeAssinatura struct {
	RowId                        int64          `json:"rowid"`
	Id                           sql.NullInt64  `json:"id"`
	Nome                         sql.NullString `json:"nome"`
	ValorMensal                  sql.NullInt64  `json:"valor_mensal"`
	LimiteLinksMensal            sql.NullInt64  `json:"limite_links_mensal"`
	LimiteRedirecionadoresMensal sql.NullInt64  `json:"limite_redirecionadores_mensal"`
	OrdenacaoAleatoriaLinks      sql.NullBool   `json:"ordenacao_aleatoria_links"`
	CriadoEm                     sql.NullString `json:"criado_em"`
	AtualizadoEm                 sql.NullString `json:"atualizado_em"`
	RemovidoEm                   sql.NullString `json:"removido_em" swaggertype:"integer"`
	Versao                       int64          `json:"versao"`
	Bitmask                      int64          `json:"bitmask"`
	CriadoEm_                    string         `json:"criado_em_"`
} // @name HistoricoPlanoDeAssinatura

type HistoricoUsuarioTemporario struct {
	RowId             int64          `json:"rowid"`
	Id                sql.NullInt64  `json:"id"`
	Cpf               sql.NullString `json:"cpf"`
	Nome              sql.NullString `json:"nome"`
	NomeDeUsuario     sql.NullString `json:"nome_de_usuario"`
	Email             sql.NullString `json:"email"`
	PlanoDeAssinatura sql.NullInt64  `json:"plano_de_assinatura"`
	CriadoEm          sql.NullString `json:"criado_em"`
	RemovidoEm        sql.NullString `json:"removido_em" swaggertype:"string"`
	Versao            int64          `json:"versao"`
	Bitmask           int64          `json:"bitmask"`
	CriadoEm_         string         `json:"criado_em_"`
} // @name HistoricoUsuarioTemporario

type HistoricoUsuario struct {
	RowId             int64          `json:"rowid"`
	Id                sql.NullInt64  `json:"id"`
	Cpf               sql.NullString `json:"cpf"`
	Nome              sql.NullString `json:"nome"`
	NomeDeUsuario     sql.NullString `json:"nome_de_usuario"`
	Email             sql.NullString `json:"email"`
	Senha             sql.NullString `json:"senha"`
	DataDeNascimento  sql.NullString `json:"data_de_nascimento"`
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

type HistoricoRedirecionador struct {
	RowId                   int64          `json:"rowid"`
	Id                      sql.NullInt64  `json:"id"`
	Nome                    sql.NullString `json:"nome"`
	CodigoHash              sql.NullString `json:"codigo_hash"`
	OrdemDeRedirecionamento sql.NullString `json:"ordem_de_redirecionamento"`
	Usuario                 sql.NullInt64  `json:"usuario"`
	CriadoEm                sql.NullString `json:"criado_em"`
	AtualizadoEm            sql.NullString `json:"atualizado_em"`
	RemovidoEm              sql.NullString `json:"removido_em" swaggertype:"string"`
	Versao                  int64          `json:"versao"`
	Bitmask                 int64          `json:"bitmask"`
	CriadoEm_               string         `json:"criado_em_"`
} // @name HistoricoRedirecionador

type HistoricoLink struct {
	RowId          int64          `json:"rowid"`
	Id             sql.NullInt64  `json:"id"`
	Link           sql.NullString `json:"link"`
	Plataforma     sql.NullString `json:"plataforma"`
	Redirecionador sql.NullInt64  `json:"redirecionador"`
	CriadoEm       sql.NullString `json:"criado_em"`
	AtualizadoEm   sql.NullString `json:"atualizado_em"`
	RemovidoEm     sql.NullString `json:"removido_em" swaggertype:"string"`
	Versao         int64          `json:"versao"`
	Bitmask        int64          `json:"bitmask"`
	CriadoEm_      string         `json:"criado_em_"`
} // @name HistoricoLink

type HistoricoModel struct {
	DB *sql.DB
}

func (h *HistoricoModel) PlanoDeAssinaturaReadAll() ([]HistoricoPlanoDeAssinatura, error) {
	var historico []HistoricoPlanoDeAssinatura

	rows, err := h.DB.Query("SELECT _ROWID, ID, NOME, VALOR_MENSAL, LIMITE_LINKS_MENSAL, LINKS_REDIRECIONAMENTOS_MENSAL, ORDENACAO_ALEATORIA_LINKS, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK, _CRIADO_EM FROM HISTORICO_PLANO_DE_ASSINATURA")

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
			&registro.LimiteLinksMensal,
			&registro.LimiteRedirecionadoresMensal,
			&registro.OrdenacaoAleatoriaLinks,
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

func (h *HistoricoModel) UsuarioTemporarioReadAll() ([]HistoricoUsuarioTemporario, error) {
	var historico []HistoricoUsuarioTemporario

	rows, err := h.DB.Query("SELECT _ROWID, ID, CPF, NOME, NOME_DE_USUARIO, EMAIL, PLANO_DE_ASSINATURA, CRIADO_EM, REMOVIDO_EM, VERSAO, BITMASK, _CRIADO_EM FROM HISTORICO_USUARIO_TEMPORARIO")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var registro HistoricoUsuarioTemporario

		if err := rows.Scan(
			&registro.RowId,
			&registro.Id,
			&registro.Cpf,
			&registro.Nome,
			&registro.NomeDeUsuario,
			&registro.Email,
			&registro.PlanoDeAssinatura,
			&registro.CriadoEm,
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

func (h *HistoricoModel) UsuarioReadAll() ([]HistoricoUsuario, error) {
	var historico []HistoricoUsuario

	rows, err := h.DB.Query("SELECT _ROWID, ID, CPF, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, PLANO_DE_ASSINATURA, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK, _CRIADO_EM FROM HISTORICO_USUARIO")

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

func (h *HistoricoModel) EmailAutenticacaoReadAll() ([]HistoricoEmailAutenticacao, error) {
	var historico []HistoricoEmailAutenticacao

	rows, err := h.DB.Query("SELECT _ROWID, ID, VALOR, TIPO, EXPIRA_EM, USUARIO, VERSAO, BITMASK, _CRIADO_EM FROM HISTORICO_EMAIL_AUTENTICACAO")

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

func (h *HistoricoModel) RedirecionadorReadAll() ([]HistoricoRedirecionador, error) {
	var historico []HistoricoRedirecionador

	rows, err := h.DB.Query("SELECT _ROWID, ID, NOME, CODIGO_HASH, ORDEM_DE_REDIRECIONAMENTO, USUARIO, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK, _CRIADO_EM FROM HISTORICO_REDIRECIONADOR")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var registro HistoricoRedirecionador

		if err := rows.Scan(
			&registro.RowId,
			&registro.Id,
			&registro.Nome,
			&registro.CodigoHash,
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

func (h *HistoricoModel) LinkReadAll() ([]HistoricoLink, error) {
	var historico []HistoricoLink

	rows, err := h.DB.Query("SELECT _ROWID, ID, LINK, PLATAFORMA, REDIRECIONADOR, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK, _CRIADO_EM FROM HISTORICO_LINK")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var registro HistoricoLink

		if err := rows.Scan(
			&registro.RowId,
			&registro.Id,
			&registro.Link,
			&registro.Plataforma,
			&registro.Redirecionador,
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
