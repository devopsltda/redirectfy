package model

import "database/sql"

type PlanoDeAssinatura struct {
	Id           int64          `json:"id"`
	Nome         string         `json:"nome"`
	ValorMensal  int64          `json:"valor_mensal"`
	CriadoEm     string         `json:"criado_em"`
	AtualizadoEm string         `json:"atualizado_em"`
	RemovidoEm   sql.NullString `json:"removido_em" swaggertype:"integer"`
} // @name PlanoDeAssinatura

type Usuario struct {
	Id                int64               `json:"id"`
	Cpf               string              `json:"cpf"`
	Nome              string              `json:"nome"`
	NomeDeUsuario     string              `json:"nome_de_usuario"`
	Email             string              `json:"email"`
	Senha             string              `json:"senha"`
	DataDeNascimento  string              `json:"data_de_nascimento"`
	PlanoDeAssinatura int64               `json:"plano_de_assinatura"`
	CriadoEm          string              `json:"criado_em"`
	AtualizadoEm      string              `json:"atualizado_em"`
	RemovidoEm        sql.NullString      `json:"removido_em" swaggertype:"string"`
} // @name Usuario

type Link struct {
	Id                      int64          `json:"id"`
	Nome                    string         `json:"nome"`
	CodigoHash              string         `json:"codigo_hash"`
	LinkWhatsapp            string         `json:"link_whatsapp"`
	LinkTelegram            string         `json:"link_telegram"`
	OrdemDeRedirecionamento string         `json:"ordem_de_redirecionamento"`
	Usuario                 int64          `json:"usuario"`
	CriadoEm                string         `json:"criado_em"`
	AtualizadoEm            string         `json:"atualizado_em"`
	RemovidoEm              sql.NullString `json:"removido_em" swaggertype:"string"`
} // @name Link

type Historico struct {
	Id               int64          `json:"id"`
	Usuario          sql.NullInt64  `json:"usuario" swaggertype:"integer"`
	ValorOriginal    sql.NullString `json:"valor_original" swaggertype:"string"`
	ValorNovo        sql.NullString `json:"valor_novo" swaggertype:"string"`
	TabelaModificada string         `json:"tabela_modificada"`
	CriadoEm         string         `json:"criado_em"`
} // @name Historico
