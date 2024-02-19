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

CREATE TABLE IF NOT EXISTS USUARIO (
    ID INTEGER PRIMARY KEY NOT NULL,
    CPF TEXT UNIQUE NOT NULL,
    NOME TEXT NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 240),
    NOME_DE_USUARIO TEXT UNIQUE NOT NULL CHECK(LENGTH(NOME) >= 3 AND LENGTH(NOME) <= 120),
    EMAIL TEXT UNIQUE NOT NULL,
    SENHA TEXT NOT NULL,
    DATA_DE_NASCIMENTO TEXT NOT NULL,
    AUTENTICADO BOOLEAN NOT NULL CHECK(AUTENTICADO IN (0, 1)) DEFAULT 0,
    PLANO_DE_ASSINATURA INTEGER NOT NULL,
    CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    ATUALIZADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    REMOVIDO_EM TEXT DEFAULT NULL,
    FOREIGN KEY(PLANO_DE_ASSINATURA) REFERENCES PLANO_DE_ASSINATURA(ID)
);

CREATE TABLE IF NOT EXISTS EMAIL_AUTENTICACAO (
    ID INTEGER PRIMARY KEY NOT NULL,
    VALOR TEXT UNIQUE NOT NULL CHECK(LENGTH(VALOR) <= 128),
    TIPO TEXT NOT NULL CHECK(TIPO IN ('senha', 'validacao')),
    EXPIRA_EM TEXT NOT NULL,
    USUARIO INTEGER NOT NULL,
    FOREIGN KEY(USUARIO) REFERENCES USUARIO(ID)
);

CREATE TABLE IF NOT EXISTS LINK (
    ID INTEGER PRIMARY KEY NOT NULL,
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

CREATE TABLE IF NOT EXISTS HISTORICO_USUARIO (
    _ROWID INTEGER,
    ID INTEGER,
    CPF TEXT,
    NOME TEXT,
    NOME_DE_USUARIO TEXT,
    EMAIL TEXT,
    SENHA TEXT,
    DATA_DE_NASCIMENTO,
    AUTENTICADO BOOLEAN,
    PLANO_DE_ASSINATURA INTEGER,
    CRIADO_EM TEXT,
    ATUALIZADO_EM TEXT,
    REMOVIDO_EM TEXT,
    VERSAO INTEGER,
    BITMASK INTEGER,
    _CRIADO_EM TEXT DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY(ID, PLANO_DE_ASSINATURA) REFERENCES USUARIO(ID, PLANO_DE_ASSINATURA)
);

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

CREATE TABLE IF NOT EXISTS HISTORICO_LINK (
    _ROWID INTEGER,
    ID INTEGER,
    NOME TEXT,
    CODIGO_HASH TEXT,
    LINK_WHATSAPP TEXT,
    LINK_TELEGRAM TEXT,
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

-- TRIGGERS
-- Importante notar que esses triggers são relacionados ao histórico, e as mudanças são monitoradas via um BITMASK, onde a soma única dos valores múltiplos de 2 atribuídos
-- a cada coluna determina o que foi alterado e qual o tipo de alteração. Por padrão, quando um registro é removido, seu BITMASK é igual a -1.

-- PLANO_DE_ASSINATURA

CREATE TRIGGER IF NOT EXISTS PLANO_DE_ASSINATURA_INSERIR_HISTORICO
AFTER INSERT ON PLANO_DE_ASSINATURA
BEGIN
    INSERT INTO HISTORICO_PLANO_DE_ASSINATURA(_ROWID, ID, NOME, VALOR_MENSAL, LIMITE, PERIODO_LIMITE, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    VALUES (NEW.ROWID, NEW.ID, NEW.NOME, NEW.VALOR_MENSAL, NEW.LIMITE, NEW.PERIODO_LIMITE, NEW.CRIADO_EM, NEW.ATUALIZADO_EM, NEW.REMOVIDO_EM, 1, 255);
END;

CREATE TRIGGER IF NOT EXISTS PLANO_DE_ASSINATURA_ATUALIZAR_HISTORICO
AFTER UPDATE ON PLANO_DE_ASSINATURA
FOR EACH ROW
BEGIN
    INSERT INTO HISTORICO_PLANO_DE_ASSINATURA(_ROWID, ID, NOME, VALOR_MENSAL, LIMITE, PERIODO_LIMITE, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    SELECT OLD.ROWID,
           CASE WHEN OLD.ID != NEW.ID THEN NEW.ID ELSE NULL END,
           CASE WHEN OLD.NOME != NEW.NOME THEN NEW.NOME ELSE NULL END,
           CASE WHEN OLD.VALOR_MENSAL != NEW.VALOR_MENSAL THEN NEW.VALOR_MENSAL ELSE NULL END,
           CASE WHEN OLD.LIMITE != NEW.LIMITE THEN NEW.LIMITE ELSE NULL END,
           CASE WHEN OLD.PERIODO_LIMITE != NEW.PERIODO_LIMITE THEN NEW.PERIODO_LIMITE ELSE NULL END,
           CASE WHEN OLD.CRIADO_EM != NEW.CRIADO_EM THEN NEW.CRIADO_EM ELSE NULL END,
           CASE WHEN OLD.ATUALIZADO_EM != NEW.ATUALIZADO_EM THEN NEW.ATUALIZADO_EM ELSE NULL END,
           CASE WHEN OLD.REMOVIDO_EM != NEW.REMOVIDO_EM THEN NEW.REMOVIDO_EM ELSE NULL END,
           (SELECT MAX(VERSAO) FROM HISTORICO_PLANO_DE_ASSINATURA WHERE _ROWID = OLD.ROWID) + 1,
           (CASE WHEN OLD.ID != NEW.ID THEN 1 ELSE 0 END) +
           (CASE WHEN OLD.NOME != NEW.NOME THEN 2 ELSE 0 END) +
           (CASE WHEN OLD.VALOR_MENSAL != NEW.VALOR_MENSAL THEN 4 ELSE 0 END) +
           (CASE WHEN OLD.LIMITE != NEW.LIMITE THEN 8 ELSE 0 END) +
           (CASE WHEN OLD.PERIODO_LIMITE != NEW.PERIODO_LIMITE THEN 16 ELSE 0 END) +
           (CASE WHEN OLD.CRIADO_EM != NEW.CRIADO_EM THEN 32 ELSE 0 END) +
           (CASE WHEN OLD.ATUALIZADO_EM != NEW.ATUALIZADO_EM THEN 64 ELSE 0 END) +
           (CASE WHEN OLD.REMOVIDO_EM != NEW.REMOVIDO_EM THEN 128 ELSE 0 END);
END;

CREATE TRIGGER IF NOT EXISTS PLANO_DE_ASSINATURA_REMOVER_HISTORICO
AFTER DELETE ON PLANO_DE_ASSINATURA
BEGIN
    INSERT INTO HISTORICO_PLANO_DE_ASSINATURA(_ROWID, ID, NOME, VALOR_MENSAL, LIMITE, PERIODO_LIMITE, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    VALUES (OLD.ROWID, OLD.ID, OLD.NOME, OLD.VALOR_MENSAL, OLD.LIMITE, OLD.PERIODO_LIMITE, OLD.CRIADO_EM, OLD.ATUALIZADO_EM, OLD.REMOVIDO_EM, (SELECT COALESCE(MAX(VERSAO), 0) FROM HISTORICO_PLANO_DE_ASSINATURA WHERE _ROWID = OLD.ROWID) + 1, -1);
END;

-- USUARIO

CREATE TRIGGER IF NOT EXISTS USUARIO_INSERIR_HISTORICO
AFTER INSERT ON USUARIO
BEGIN
    INSERT INTO HISTORICO_USUARIO(_ROWID, ID, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, AUTENTICADO, PLANO_DE_ASSINATURA, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    VALUES (NEW.ROWID, NEW.ID, NEW.NOME, NEW.NOME_DE_USUARIO, NEW.EMAIL, NEW.SENHA, NEW.DATA_DE_NASCIMENTO, NEW.AUTENTICADO, NEW.PLANO_DE_ASSINATURA, NEW.CRIADO_EM, NEW.ATUALIZADO_EM, NEW.REMOVIDO_EM, 1, 4095);
END;

CREATE TRIGGER IF NOT EXISTS USUARIO_ATUALIZAR_HISTORICO
AFTER UPDATE ON USUARIO
FOR EACH ROW
BEGIN
    INSERT INTO HISTORICO_USUARIO(_ROWID, ID, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, AUTENTICADO, PLANO_DE_ASSINATURA, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    SELECT OLD.ROWID,
           CASE WHEN OLD.ID != NEW.ID THEN NEW.ID ELSE NULL END,
           CASE WHEN OLD.NOME != NEW.NOME THEN NEW.NOME ELSE NULL END,
           CASE WHEN OLD.NOME_DE_USUARIO != NEW.NOME_DE_USUARIO THEN NEW.NOME_DE_USUARIO ELSE NULL END,
           CASE WHEN OLD.EMAIL != NEW.EMAIL THEN NEW.EMAIL ELSE NULL END,
           CASE WHEN OLD.SENHA != NEW.SENHA THEN NEW.SENHA ELSE NULL END,
           CASE WHEN OLD.DATA_DE_NASCIMENTO != NEW.DATA_DE_NASCIMENTO THEN NEW.DATA_DE_NASCIMENTO ELSE NULL END,
           CASE WHEN OLD.AUTENTICADO != NEW.AUTENTICADO THEN NEW.AUTENTICADO ELSE NULL END,
           CASE WHEN OLD.PLANO_DE_ASSINATURA != NEW.PLANO_DE_ASSINATURA THEN NEW.PLANO_DE_ASSINATURA ELSE NULL END,
           CASE WHEN OLD.CRIADO_EM != NEW.CRIADO_EM THEN NEW.CRIADO_EM ELSE NULL END,
           CASE WHEN OLD.ATUALIZADO_EM != NEW.ATUALIZADO_EM THEN NEW.ATUALIZADO_EM ELSE NULL END,
           CASE WHEN OLD.REMOVIDO_EM != NEW.REMOVIDO_EM THEN NEW.REMOVIDO_EM ELSE NULL END,
           (SELECT MAX(VERSAO) FROM HISTORICO_USUARIO WHERE _ROWID = OLD.ROWID) + 1,
           (CASE WHEN OLD.ID != NEW.ID THEN 1 ELSE 0 END) +
           (CASE WHEN OLD.NOME != NEW.NOME THEN 2 ELSE 0 END) +
           (CASE WHEN OLD.NOME_DE_USUARIO != NEW.NOME_DE_USUARIO THEN 4 ELSE 0 END) +
           (CASE WHEN OLD.EMAIL != NEW.EMAIL THEN 8 ELSE 0 END) +
           (CASE WHEN OLD.SENHA != NEW.SENHA THEN 16 ELSE 0 END) +
           (CASE WHEN OLD.DATA_DE_NASCIMENTO != NEW.DATA_DE_NASCIMENTO THEN 32 ELSE 0 END) +
           (CASE WHEN OLD.AUTENTICADO != NEW.AUTENTICADO THEN 64 ELSE 0 END) +
           (CASE WHEN OLD.PLANO_DE_ASSINATURA != NEW.PLANO_DE_ASSINATURA THEN 128 ELSE 0 END) +
           (CASE WHEN OLD.CRIADO_EM != NEW.CRIADO_EM THEN 256 ELSE 0 END) +
           (CASE WHEN OLD.ATUALIZADO_EM != NEW.ATUALIZADO_EM THEN 512 ELSE 0 END) +
           (CASE WHEN OLD.REMOVIDO_EM != NEW.REMOVIDO_EM THEN 1024 ELSE 0 END);
END;

CREATE TRIGGER IF NOT EXISTS USUARIO_REMOVER_HISTORICO
AFTER DELETE ON USUARIO
BEGIN
    INSERT INTO HISTORICO_USUARIO(_ROWID, ID, NOME, NOME_DE_USUARIO, EMAIL, SENHA, DATA_DE_NASCIMENTO, AUTENTICADO, PLANO_DE_ASSINATURA, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    VALUES (OLD.ROWID, OLD.ID, OLD.NOME, OLD.NOME_DE_USUARIO, OLD.EMAIL, OLD.SENHA, OLD.DATA_DE_NASCIMENTO, OLD.AUTENTICADO, OLD.PLANO_DE_ASSINATURA, OLD.CRIADO_EM, OLD.ATUALIZADO_EM, OLD.REMOVIDO_EM, (SELECT COALESCE(MAX(VERSAO), 0) FROM HISTORICO_USUARIO WHERE _ROWID = OLD.ROWID) + 1, -1);
END;

-- EMAIL_AUTENTICACAO

CREATE TRIGGER IF NOT EXISTS EMAIL_AUTENTICACAO_INSERIR_HISTORICO
AFTER INSERT ON EMAIL_AUTENTICACAO
BEGIN
    INSERT INTO HISTORICO_EMAIL_AUTENTICACAO(_ROWID, ID, VALOR, TIPO, EXPIRA_EM, USUARIO, VERSAO, BITMASK)
    VALUES (NEW.ROWID, NEW.ID, NEW.VALOR, NEW.TIPO, NEW.EXPIRA_EM, NEW.USUARIO, 1, 31);
END;

CREATE TRIGGER IF NOT EXISTS EMAIL_AUTENTICACAO_ATUALIZAR_HISTORICO
AFTER UPDATE ON EMAIL_AUTENTICACAO
FOR EACH ROW
BEGIN
    INSERT INTO HISTORICO_EMAIL_AUTENTICACAO(_ROWID, ID, VALOR, TIPO, EXPIRA_EM, USUARIO, VERSAO, BITMASK)
    SELECT OLD.ROWID,
           CASE WHEN OLD.ID != NEW.ID THEN NEW.ID ELSE NULL END,
           CASE WHEN OLD.VALOR != NEW.VALOR THEN NEW.VALOR ELSE NULL END,
           CASE WHEN OLD.TIPO != NEW.TIPO THEN NEW.TIPO ELSE NULL END,
           CASE WHEN OLD.EXPIRA_EM != NEW.EXPIRA_EM THEN NEW.EXPIRA_EM ELSE NULL END,
           CASE WHEN OLD.USUARIO != NEW.USUARIO THEN NEW.USUARIO ELSE NULL END,
           (SELECT MAX(VERSAO) FROM HISTORICO_EMAIL_AUTENTICACAO WHERE _ROWID = OLD.ROWID) + 1,
           (CASE WHEN OLD.ID != NEW.ID THEN 1 ELSE 0 END) +
           (CASE WHEN OLD.VALOR != NEW.VALOR THEN 2 ELSE 0 END) +
           (CASE WHEN OLD.TIPO != NEW.TIPO THEN 4 ELSE 0 END) +
           (CASE WHEN OLD.EXPIRA_EM != NEW.EXPIRA_EM THEN 8 ELSE 0 END) +
           (CASE WHEN OLD.USUARIO != NEW.USUARIO THEN 16 ELSE 0 END);
END;

CREATE TRIGGER IF NOT EXISTS EMAIL_AUTENTICACAO_REMOVER_HISTORICO
AFTER DELETE ON EMAIL_AUTENTICACAO
BEGIN
    INSERT INTO HISTORICO_EMAIL_AUTENTICACAO(_ROWID, ID, VALOR, TIPO, EXPIRA_EM, USUARIO, VERSAO, BITMASK)
    VALUES (OLD.ROWID, OLD.ID, OLD.VALOR, OLD.TIPO, OLD.EXPIRA_EM, OLD.USUARIO, (SELECT COALESCE(MAX(VERSAO), 0) FROM HISTORICO_EMAIL_AUTENTICACAO WHERE _ROWID = OLD.ROWID) + 1, -1);
END;

-- LINK

CREATE TRIGGER IF NOT EXISTS LINK_INSERIR_HISTORICO
AFTER INSERT ON LINK
BEGIN
    INSERT INTO HISTORICO_LINK(_ROWID, ID, NOME, CODIGO_HASH, LINK_WHATSAPP, LINK_TELEGRAM, ORDEM_DE_REDIRECIONAMENTO, USUARIO, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    VALUES (NEW.ROWID, NEW.ID, NEW.NOME, NEW.CODIGO_HASH, NEW.LINK_WHATSAPP, NEW.LINK_TELEGRAM, NEW.ORDEM_DE_REDIRECIONAMENTO, NEW.USUARIO, NEW.CRIADO_EM, NEW.ATUALIZADO_EM, NEW.REMOVIDO_EM, 1, 1023);
END;

CREATE TRIGGER IF NOT EXISTS LINK_ATUALIZAR_HISTORICO
AFTER UPDATE ON LINK
FOR EACH ROW
BEGIN
    INSERT INTO HISTORICO_LINK(_ROWID, ID, NOME, CODIGO_HASH, LINK_WHATSAPP, LINK_TELEGRAM, ORDEM_DE_REDIRECIONAMENTO, USUARIO, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    SELECT OLD.ROWID,
           CASE WHEN OLD.ID != NEW.ID THEN NEW.ID ELSE NULL END,
           CASE WHEN OLD.NOME != NEW.NOME THEN NEW.NOME ELSE NULL END,
           CASE WHEN OLD.CODIGO_HASH != NEW.CODIGO_HASH THEN NEW.CODIGO_HASH ELSE NULL END,
           CASE WHEN OLD.LINK_WHATSAPP != NEW.LINK_WHATSAPP THEN NEW.LINK_WHATSAPP ELSE NULL END,
           CASE WHEN OLD.LINK_TELEGRAM != NEW.LINK_TELEGRAM THEN NEW.LINK_TELEGRAM ELSE NULL END,
           CASE WHEN OLD.ORDEM_DE_REDIRECIONAMENTO != NEW.ORDEM_DE_REDIRECIONAMENTO THEN NEW.ORDEM_DE_REDIRECIONAMENTO ELSE NULL END,
           CASE WHEN OLD.USUARIO != NEW.USUARIO THEN NEW.USUARIO ELSE NULL END,
           CASE WHEN OLD.CRIADO_EM != NEW.CRIADO_EM THEN NEW.CRIADO_EM ELSE NULL END,
           CASE WHEN OLD.ATUALIZADO_EM != NEW.ATUALIZADO_EM THEN NEW.ATUALIZADO_EM ELSE NULL END,
           CASE WHEN OLD.REMOVIDO_EM != NEW.REMOVIDO_EM THEN NEW.REMOVIDO_EM ELSE NULL END,
           (SELECT MAX(VERSAO) FROM HISTORICO_LINK WHERE _ROWID = OLD.ROWID) + 1,
           (CASE WHEN OLD.ID != NEW.ID THEN 1 ELSE 0 END) +
           (CASE WHEN OLD.NOME != NEW.NOME THEN 2 ELSE 0 END) +
           (CASE WHEN OLD.CODIGO_HASH != NEW.CODIGO_HASH THEN 4 ELSE 0 END) +
           (CASE WHEN OLD.LINK_WHATSAPP != NEW.LINK_WHATSAPP THEN 8 ELSE 0 END) +
           (CASE WHEN OLD.LINK_TELEGRAM != NEW.LINK_TELEGRAM THEN 16 ELSE 0 END) +
           (CASE WHEN OLD.ORDEM_DE_REDIRECIONAMENTO != NEW.ORDEM_DE_REDIRECIONAMENTO THEN 32 ELSE 0 END) +
           (CASE WHEN OLD.USUARIO != NEW.USUARIO THEN 64 ELSE 0 END) +
           (CASE WHEN OLD.CRIADO_EM != NEW.CRIADO_EM THEN 128 ELSE 0 END) +
           (CASE WHEN OLD.ATUALIZADO_EM != NEW.ATUALIZADO_EM THEN 256 ELSE 0 END) +
           (CASE WHEN OLD.REMOVIDO_EM != NEW.REMOVIDO_EM THEN 512 ELSE 0 END);
END;

CREATE TRIGGER IF NOT EXISTS LINK_REMOVER_HISTORICO
AFTER DELETE ON LINK
BEGIN
    INSERT INTO HISTORICO_LINK(_ROWID, ID, NOME, CODIGO_HASH, LINK_WHATSAPP, LINK_TELEGRAM, ORDEM_DE_REDIRECIONAMENTO, USUARIO, CRIADO_EM, ATUALIZADO_EM, REMOVIDO_EM, VERSAO, BITMASK)
    VALUES (OLD.ROWID, OLD.ID, OLD.NOME, OLD.CODIGO_HASH, OLD.LINK_WHATSAPP, OLD.LINK_TELEGRAM, OLD.ORDEM_DE_REDIRECIONAMENTO, OLD.USUARIO, OLD.CRIADO_EM, OLD.ATUALIZADO_EM, OLD.REMOVIDO_EM, (SELECT COALESCE(MAX(VERSAO), 0) FROM HISTORICO_LINK WHERE _ROWID = OLD.ROWID) + 1, -1);
END;
