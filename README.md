# API do Redirect Max 

API para manipulação do sistema Redirect Max

## Requisitos

- Make
    - Linux (não precisa instalar porque já vem de fábrica)
    - [Windows](https://gnuwin32.sourceforge.net/packages/make.htm) (utilizado para rodar os scripts da aplicação)
- [Go](https://go.dev/doc/install) (>=1.21)
    - [Air](https://github.com/cosmtrek/air) (responsável pelo *live reload* do código)
    - [Swag](https://github.com/swaggo/swag) (responsável por gerar a documentação do Swagger a partir dos comentários no código)
- [Sqlite3](https://www.sqlite.org/download.html) (banco de dados)

## Como Iniciar o Ambiente de Desenvolvimento

As instruções abaixo assumem que o desenvolvedor está em um sistema Linux. Se esse não for o caso, utilize comandos equivalentes para o seu sistema.

1. Clone o repositório

```bash
git clone git@github.com:TheDevOpsCorp/redirect.git
```

2. Vá para a pasta raiz do repositório

```bash
cd redirect
```

3. Mude para a branch correta

```bash
git checkout main-but-it-works
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
JWT_SECRET=         # Insira uma chave segura aqui
JWT_REFRESH_SECRET= # Insira uma outra chave segura aqui
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

8. Acesse a documentação no [servidor local](http://localhost:8080/api/swagger/index.html) (porta 8080 por padrão, modifique caso você tenha alterado a porta)

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
