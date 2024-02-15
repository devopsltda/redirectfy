# Redirectify

O Redirectify é uma solução para redirecionamento de links que abrange tanto uma API para uso programático como uma aplicação Web responsiva para o usuário final.

## Roadmap

Acompanhe o desenvolvimento [aqui](https://coda.io/d/Mapa-de-Desenvolvimento-Redirect_d3Jz7W_oyZx/Mapa_suI7Y#_luOTT).

## Requisitos para o Ambiente de Desenvolvimento

O ambiente de desenvolvimento pode ser configurado via Docker ou construído a partir do código-fonte, e os requisitos são os seguintes:

### Docker

Apenas tenha o Docker instalado.

### Source

- Make (utilizado para rodar os scripts da aplicação)
    - No Linux, procure o pacote específico para o seu sistema, se ele não vir por padrão
    - No [Windows](https://gnuwin32.sourceforge.net/packages/make.htm)
- [Go](https://go.dev/doc/install) (>=1.21)
    - [Air](https://github.com/cosmtrek/air) (responsável pelo *live reload* do código)
    - [Swag](https://github.com/swaggo/swag) (responsável por gerar a documentação do Swagger a partir dos comentários no código)
- [Sqlite3](https://www.sqlite.org/download.html) (banco de dados local)

## Como Iniciar o Ambiente de Desenvolvimento

### Com Docker

1. Clone o repositório

```bash
git clone git@github.com:TheDevOpsCorp/redirectify.git
```

2. Vá para a pasta raiz do repositório

```bash
cd redirectify
```

3. Copie a configuração de ambiente

```bash
cp .env.example .env
```

4. Complete as variáveis de ambiente como disposto abaixo

```bash
PORT=8080
APP_ENV=local
DB_URL=./storage/test.db
PEPPER=insira-um-texto-forte-aqui
JWT_SECRET=insira-alguma-coisa-aqui
JWT_REFRESH_SECRET=insira-alguma-coisa-aqui-2
```

5. Roda o Docker

```bash
docker compose up
```

6. Acesse a documentação no [servidor local](http://localhost:8080/docs/index.html) (porta 8080 por padrão, modifique caso você tenha alterado a porta)

### A partir do código-fonte

As instruções abaixo assumem que o desenvolvedor está em um sistema Linux. Se esse não for o caso, utilize comandos equivalentes para o seu sistema.

1. Clone o repositório

```bash
git clone git@github.com:TheDevOpsCorp/redirectify.git
```

2. Vá para a pasta raiz do repositório

```bash
cd redirectify
```

3. Baixe as dependências do projeto

```bash
go mod tidy && go get ./...
```

4. Copie o arquivo com as variáveis de ambiente

```bash
cp .env.example .env
```

5. Complete as variáveis de ambiente como disposto abaixo

```bash
PORT=8080
APP_ENV=local
DB_URL=./storage/test.db
PEPPER=insira-um-texto-forte-aqui
JWT_SECRET=insira-alguma-coisa-aqui
JWT_REFRESH_SECRET=insira-alguma-coisa-aqui-2
```

6. Gere a documentação

```bash
make docs
```

7. Execute o código

- Com *live reload*

```bash
make watch
```

- Sem *live reload*

```bash
make run
```

8. Acesse a documentação no [servidor local](http://localhost:8080/docs/index.html) (porta 8080 por padrão, modifique caso você tenha alterado a porta)

## Comandos `make`

- Rodar todos os comandos, incluindo testes

```bash
make all build
```

- Gerar o executável da aplicação

```bash
make build
```

- Rodar a aplicação

```bash
make run
```

- Gerar a documentação

```bash
make docs
```

- Rodar a aplicação com *live reload*

```bash
make watch
```

- Rodar testes

```bash
make test
```

- Limpar o executável antigo

```bash
make clean
```
