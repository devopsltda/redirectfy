package model

import "database/sql"

type PlanoDeAssinatura struct {
	Id           int64          `json:"id"`
	Nome         string         `json:"nome"`
	ValorMensal  int64          `json:"valor_mensal"`
	CriadoEm     string         `json:"criado_em"`
	AtualizadoEm string         `json:"atualizado_em"`
	RemovidoEm   sql.NullString `json:"removido_em"`
}

type PlanoDeAssinaturaFK struct {
	Id    int64             `json:"id"`
	Dados PlanoDeAssinatura `json:"dados"`
}

type Usuario struct {
	Id                int64               `json:"id"`
	Cpf               string              `json:"cpf"`
	Nome              string              `json:"nome"`
	NomeDeUsuario     string              `json:"nome_de_usuario"`
	Email             string              `json:"email"`
	Senha             string              `json:"senha"`
	DataDeNascimento  string              `json:"data_de_nascimento"`
	PlanoDeAssinatura PlanoDeAssinaturaFK `json:"plano_de_assinatura"`
	CriadoEm          string              `json:"criado_em"`
	AtualizadoEm      string              `json:"atualizado_em"`
	RemovidoEm        sql.NullString      `json:"removido_em"`
}

type UsuarioFK struct {
	Id    int64   `json:"id"`
	Dados Usuario `json:"dados"`
}

type Link struct {
	Id                      int64          `json:"id"`
	Nome                    string         `json:"nome"`
	CodigoHash              string         `json:"codigo_hash"`
	LinkWhatsapp            sql.NullString `json:"link_whatsapp"`
	LinkTelegram            sql.NullString `json:"link_telegram"`
	OrdemDeRedirecionamento string         `json:"ordem_de_redirecionamento"`
	Usuario                 UsuarioFK      `json:"usuario"`
	CriadoEm                string         `json:"criado_em"`
	AtualizadoEm            string         `json:"atualizado_em"`
	RemovidoEm              sql.NullString `json:"removido_em"`
}

type Historico struct {
	Id               int64          `json:"id"`
	Usuario          sql.NullInt64  `json:"usuario"`
	ValorOriginal    sql.NullString `json:"valor_original"`
	ValorNovo        sql.NullString `json:"valor_novo"`
	TabelaModificada string         `json:"tabela_modificada"`
	CriadoEm         string         `json:"criado_em"`
}
