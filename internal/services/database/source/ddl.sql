-- DDL

CREATE TABLE IF NOT EXISTS PLANO_DE_ASSINATURA (
    ID INTEGER PRIMARY KEY NOT NULL,
    NOME TEXT UNIQUE NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    VALOR_MENSAL INTEGER NOT NULL,
    LIMITE INTEGER NOT NULL,
    PERIODO_LIMITE TEXT NOT NULL CHECK(PERIODO_LIMITE IN ('s', 'm', 'h', 'd', 'M')),
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS USUARIO_TEMPORARIO (
    ID INTEGER PRIMARY KEY NOT NULL,
    CPF TEXT UNIQUE NOT NULL,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 240),
    NOME_DE_USUARIO TEXT UNIQUE NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    EMAIL TEXT UNIQUE NOT NULL,
    PLANO_DE_ASSINATURA TEXT NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(PLANO_DE_ASSINATURA) REFERENCES PLANO_DE_ASSINATURA(NOME)
);

CREATE TABLE IF NOT EXISTS USUARIO (
    ID INTEGER PRIMARY KEY NOT NULL,
    CPF TEXT UNIQUE NOT NULL,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 240),
    NOME_DE_USUARIO TEXT UNIQUE NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    EMAIL TEXT UNIQUE NOT NULL,
    SENHA TEXT NOT NULL,
    DATA_DE_NASCIMENTO TEXT NOT NULL,
    PLANO_DE_ASSINATURA TEXT NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(PLANO_DE_ASSINATURA) REFERENCES PLANO_DE_ASSINATURA(NOME)
);

CREATE TABLE IF NOT EXISTS EMAIL_AUTENTICACAO (
    ID INTEGER PRIMARY KEY NOT NULL,
    VALOR TEXT UNIQUE NOT NULL CHECK(LENGTH(VALOR) <= 128),
    TIPO TEXT NOT NULL CHECK(TIPO IN ('senha', 'validacao')),
    EXPIRA_EM TEXT NOT NULL,
    USUARIO INTEGER NOT NULL,
    FOREIGN KEY(USUARIO) REFERENCES USUARIO(ID)
);

CREATE TABLE IF NOT EXISTS REDIRECIONADOR (
    ID INTEGER PRIMARY KEY NOT NULL,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    CODIGO_HASH TEXT UNIQUE NOT NULL CHECK(LENGTH(CODIGO_HASH) >= 10),
    ORDEM_DE_REDIRECIONAMENTO TEXT NOT NULL,
    USUARIO INTEGER NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(USUARIO) REFERENCES USUARIO(ID)
);

CREATE TABLE IF NOT EXISTS LINK (
    ID INTEGER PRIMARY KEY NOT NULL,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    LINK TEXT NOT NULL,
    PLATAFORMA TEXT NOT NULL CHECK(PLATAFORMA IN ('telegram', 'whatsapp')),
    REDIRECIONADOR TEXT NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(REDIRECIONADOR) REFERENCES REDIRECIONADOR(CODIGO_HASH)
);

CREATE TABLE IF NOT EXISTS HISTORICO_PLANO_DE_ASSINATURA (
    _ROWID INTEGER,
    ID INTEGER,
    NOME TEXT,
    VALOR_MENSAL INTEGER,
    LIMITE INTEGER,
    PERIODO_LIMITE TEXT,
    CRIADO_EM TEXT,
    ATUALIZADO_EM TEXT,
    REMOVIDO_EM TEXT,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID) REFERENCES PLANO_DE_ASSINATURA(ID)
);

CREATE INDEX IDX_HISTORICO_PLANO_DE_ASSINATURA ON HISTORICO_PLANO_DE_ASSINATURA (_ROWID);

CREATE TABLE IF NOT EXISTS HISTORICO_USUARIO_TEMPORARIO (
    _ROWID INTEGER,
    ID INTEGER,
    CPF TEXT,
    NOME TEXT,
    NOME_DE_USUARIO TEXT,
    EMAIL TEXT,
    PLANO_DE_ASSINATURA TEXT,
    CRIADO_EM TEXT,
    ATUALIZADO_EM TEXT,
    REMOVIDO_EM TEXT,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID, PLANO_DE_ASSINATURA) REFERENCES USUARIO(ID, PLANO_DE_ASSINATURA)
);

CREATE INDEX IDX_HISTORICO_USUARIO_TEMPORARIO ON HISTORICO_USUARIO_TEMPORARIO (_ROWID);

CREATE TABLE IF NOT EXISTS HISTORICO_USUARIO (
    _ROWID INTEGER,
    ID INTEGER,
    CPF TEXT,
    NOME TEXT,
    NOME_DE_USUARIO TEXT,
    EMAIL TEXT,
    SENHA TEXT,
    DATA_DE_NASCIMENTO,
    PLANO_DE_ASSINATURA TEXT,
    CRIADO_EM TEXT,
    ATUALIZADO_EM TEXT,
    REMOVIDO_EM TEXT,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID, PLANO_DE_ASSINATURA) REFERENCES USUARIO(ID, PLANO_DE_ASSINATURA)
);

CREATE INDEX IDX_HISTORICO_USUARIO ON HISTORICO_USUARIO (_ROWID);

CREATE TABLE IF NOT EXISTS HISTORICO_EMAIL_AUTENTICACAO (
    _ROWID INTEGER,
    ID INTEGER,
    VALOR TEXT,
    TIPO TEXT,
    EXPIRA_EM TEXT,
    USUARIO INTEGER,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID, USUARIO) REFERENCES EMAIL_AUTENTICACAO(ID, USUARIO)
);

CREATE INDEX IDX_HISTORICO_EMAIL_AUTENTICACAO ON HISTORICO_EMAIL_AUTENTICACAO (_ROWID);

CREATE TABLE IF NOT EXISTS HISTORICO_REDIRECIONADOR (
    _ROWID INTEGER,
    ID INTEGER,
    NOME TEXT,
    CODIGO_HASH TEXT,
    ORDEM_DE_REDIRECIONAMENTO TEXT,
    USUARIO INTEGER,
    CRIADO_EM TEXT,
    ATUALIZADO_EM TEXT,
    REMOVIDO_EM TEXT,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID, USUARIO) REFERENCES LINK(ID, USUARIO)
);

CREATE INDEX IDX_HISTORICO_REDIRECIONADOR ON HISTORICO_REDIRECIONADOR (_ROWID);

CREATE TABLE IF NOT EXISTS HISTORICO_LINK (
    _ROWID INTEGER,
    ID INTEGER,
    NOME TEXT,
    LINK TEXT,
    PLATAFORMA TEXT,
    REDIRECIONADOR TEXT,
    CRIADO_EM TEXT,
    ATUALIZADO_EM TEXT,
    REMOVIDO_EM TEXT,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID, REDIRECIONADOR) REFERENCES LINK(ID, REDIRECIONADOR)
);

CREATE INDEX IDX_HISTORICO_LINK ON HISTORICO_LINK (_ROWID);
