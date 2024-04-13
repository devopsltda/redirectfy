-- DDL

CREATE TABLE IF NOT EXISTS PLANO_DE_ASSINATURA (
    ID INTEGER PRIMARY KEY,
    NOME TEXT UNIQUE NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    VALOR INTEGER NOT NULL,
    LIMITE_LINKS INTEGER NOT NULL,
    ORDENACAO_ALEATORIA_LINKS BOOL CHECK(ORDENACAO_ALEATORIA_LINKS IN (0, 1)) DEFAULT 0 NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL
);

CREATE INDEX IF NOT EXISTS IDX_PLANO_DE_ASSINATURA ON PLANO_DE_ASSINATURA (NOME);

CREATE TABLE IF NOT EXISTS USUARIO_KIRVANO (
    ID INTEGER PRIMARY KEY,
    CPF TEXT UNIQUE NOT NULL,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 240),
    NOME_DE_USUARIO TEXT UNIQUE NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    EMAIL TEXT UNIQUE NOT NULL,
    PLANO_DE_ASSINATURA TEXT NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(PLANO_DE_ASSINATURA) REFERENCES PLANO_DE_ASSINATURA(NOME)
);

CREATE INDEX IF NOT EXISTS IDX_USUARIO_KIRVANO ON USUARIO_KIRVANO (ID);

CREATE TABLE IF NOT EXISTS USUARIO (
    ID INTEGER PRIMARY KEY,
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

CREATE INDEX IF NOT EXISTS IDX_USUARIO ON USUARIO (EMAIL, NOME_DE_USUARIO);

CREATE TABLE IF NOT EXISTS EMAIL_AUTENTICACAO (
    ID INTEGER PRIMARY KEY,
    VALOR TEXT UNIQUE NOT NULL CHECK(LENGTH(VALOR) <= 128),
    TIPO TEXT NOT NULL CHECK(TIPO IN ('senha', 'nova_senha')),
    EXPIRA_EM TEXT NOT NULL,
    USUARIO_TEMPORARIO INTEGER NOT NULL,
    FOREIGN KEY(USUARIO_TEMPORARIO) REFERENCES USUARIO_TEMPORARIO(ID)
);

CREATE INDEX IF NOT EXISTS IDX_EMAIL_AUTENTICACAO ON EMAIL_AUTENTICACAO (VALOR);

CREATE TABLE IF NOT EXISTS REDIRECIONADOR (
    ID INTEGER PRIMARY KEY,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    CODIGO_HASH TEXT UNIQUE NOT NULL CHECK(LENGTH(CODIGO_HASH) >= 10),
    ORDEM_DE_REDIRECIONAMENTO TEXT NOT NULL,
    USUARIO TEXT NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(USUARIO) REFERENCES USUARIO(NOME_DE_USUARIO)
);

CREATE INDEX IF NOT EXISTS IDX_REDIRECIONADOR ON REDIRECIONADOR (CODIGO_HASH);

CREATE TABLE IF NOT EXISTS LINK (
    ID INTEGER PRIMARY KEY,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    LINK TEXT NOT NULL,
    PLATAFORMA TEXT NOT NULL CHECK(PLATAFORMA IN ('telegram', 'whatsapp')),
    REDIRECIONADOR TEXT NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(REDIRECIONADOR) REFERENCES REDIRECIONADOR(CODIGO_HASH)
);

CREATE INDEX IF NOT EXISTS IDX_LINK ON LINK (ID, REDIRECIONADOR);

CREATE TABLE IF NOT EXISTS HISTORICO_PLANO_DE_ASSINATURA (
    _ROWID INTEGER,
    ID INTEGER,
    NOME TEXT,
    VALOR INTEGER,
    LIMITE_LINKS INTEGER,
    ORDENACAO_ALEATORIA_LINKS BOOL,
    CRIADO_EM TEXT,
    ATUALIZADO_EM TEXT,
    REMOVIDO_EM TEXT,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID) REFERENCES PLANO_DE_ASSINATURA(ID)
);

CREATE INDEX IF NOT EXISTS IDX_HISTORICO_PLANO_DE_ASSINATURA ON HISTORICO_PLANO_DE_ASSINATURA (_ROWID);

CREATE TABLE IF NOT EXISTS HISTORICO_USUARIO_KIRVANO (
    _ROWID INTEGER,
    ID INTEGER,
    CPF TEXT,
    NOME TEXT,
    NOME_DE_USUARIO TEXT,
    EMAIL TEXT,
    PLANO_DE_ASSINATURA TEXT,
    CRIADO_EM TEXT,
    REMOVIDO_EM TEXT,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID) REFERENCES USUARIO_KIRVANO(ID)
);

CREATE INDEX IF NOT EXISTS IDX_HISTORICO_USUARIO_KIRVANO ON HISTORICO_USUARIO_KIRVANO (_ROWID);

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
    FOREIGN KEY(ID) REFERENCES USUARIO(ID)
);

CREATE INDEX IF NOT EXISTS IDX_HISTORICO_USUARIO ON HISTORICO_USUARIO (_ROWID);

CREATE TABLE IF NOT EXISTS HISTORICO_EMAIL_AUTENTICACAO (
    _ROWID INTEGER,
    ID INTEGER,
    VALOR TEXT,
    TIPO TEXT,
    EXPIRA_EM TEXT,
    USUARIO_TEMPORARIO INTEGER,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID) REFERENCES EMAIL_AUTENTICACAO(ID)
);

CREATE INDEX IF NOT EXISTS IDX_HISTORICO_EMAIL_AUTENTICACAO ON HISTORICO_EMAIL_AUTENTICACAO (_ROWID);

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
    FOREIGN KEY(ID) REFERENCES LINK(ID)
);

CREATE INDEX IF NOT EXISTS IDX_HISTORICO_REDIRECIONADOR ON HISTORICO_REDIRECIONADOR (_ROWID);

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
    FOREIGN KEY(ID) REFERENCES LINK(ID)
);

CREATE INDEX IF NOT EXISTS IDX_HISTORICO_LINK ON HISTORICO_LINK (_ROWID);
