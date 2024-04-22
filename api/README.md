# Redirectfy

O Redirectfy é uma solução para redirecionamento de links que abrange tanto uma API para uso programático como uma aplicação Web responsiva para o usuário final.

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
  - [Turso](https://turso.tech/) (banco de dados em nuvem)

## Comandos `make`

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
