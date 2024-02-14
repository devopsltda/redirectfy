CREATE TABLE IF NOT EXISTS PLANO_DE_ASSINATURA (
    ID INTEGER PRIMARY KEY,
    NOME TEXT UNIQUE NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    VALOR_MENSAL INTEGER NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS USUARIO (
    ID INTEGER PRIMARY KEY,
    CPF TEXT UNIQUE NOT NULL,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 240),
    NOME_DE_USUARIO TEXT UNIQUE NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    EMAIL TEXT UNIQUE NOT NULL,
    SENHA TEXT NOT NULL,
    DATA_DE_NASCIMENTO TEXT NOT NULL,
    PLANO_DE_ASSINATURA INTEGER NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(PLANO_DE_ASSINATURA) REFERENCES PLANO_DE_ASSINATURA(ID)
);

CREATE TABLE IF NOT EXISTS LINK (
    ID INTEGER PRIMARY KEY,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    CODIGO_HASH TEXT UNIQUE NOT NULL CHECK(LENGTH(CODIGO_HASH) >= 10),
    LINK_WHATSAPP TEXT,
    LINK_TELEGRAM TEXT,
    ORDEM_DE_REDIRECIONAMENTO TEXT NOT NULL,
    USUARIO INTEGER NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(USUARIO) REFERENCES USUARIO(ID)
);

CREATE TABLE IF NOT EXISTS HISTORICO (
    ID INTEGER PRIMARY KEY,
    USUARIO INTEGER DEFAULT NULL,
    VALOR_ORIGINAL TEXT DEFAULT NULL,
    VALOR_NOVO TEXT DEFAULT NULL,
    TABELA_MODIFICADA TEXT NOT NULL,
    CRIADO_EM TEXT NOT NULL,
    FOREIGN KEY(USUARIO) REFERENCES USUARIO(ID)
);
