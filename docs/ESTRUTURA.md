# Estrutura do Projeto <!-- ignore in toc -->

<!--toc:start-->
- [Backend](#backend)
    - [Tecnologias](#tecnologias)
    - [Diretórios Mais Relevantes](#diretórios-mais-relevantes)
    - [Arquivos de Configuração](#arquivos-de-configuração)
    - [Diretórios Autogerados](#diretórios-autogerados)
- [Frontend](#frontend)
<!--toc:end-->

## Backend

O backend foi construído inicialmente como o servidor principal que operaria
como front e back, e por isso sua estrutura pode diferir um pouco de outros
backends mais tradicionais.

É esperado que o desenvolvedor já tenha uma experiência mínima em [Go][1] e
uso de ferramentas como [make][5] e [Docker][6], sendo capaz de:

- Entender a sintaxe da linguagem;
- Conseguir deduzir ou saber como funcionam as principais bibliotecas;
- Conseguir executar scripts simples;

Esses conhecimentos, caso não estejam afiados, podem ser facilmente adquiridos
abaixo:

- [Tour of Go][7] (Fornece uma visão geral da linguagem)
- [Makefile By Example][8] (Fornece uma visão geral sobre o uso do `make`)
- [Docker Curriculum][9] (Fornece uma visão geral sobre o uso do `docker`)

> [!IMPORTANT]
> Algumas coisas importantes para quem nunca mexeu com [Go][1]:
>
> - Para exportar uma função, tipo, estrutura, etc. É necessário que esta seja
> iniciada com letra maiúscula;
> - Para que um arquivo seja executado pelo comando `go test`, ele deve
> terminar em `_test.go`;
> - Não podem haver dependências circulares no código (um pacote que depende
> de um pacote que depende dele).

### Tecnologias

- [Go][1] (1.22.1): Linguagem de programação do projeto
- [Make][5] (4.4.1): Ferramenta de script usada no projeto
- [Docker][6] (26.0.0): Ferramenta de conteinerização usada no projeto
- [Turso][10] (0.92.0): Banco de dados baseado em [SQLite][14] usado no projeto
- [Mailtrap][11]: Serviço de email usado no desenvolvimento do projeto
- [Kirvano][12]: Serviço de pagamentos usado no projeto
- [Fly][13]: Serviço de hospedagem usado no projeto

### Diretórios Mais Relevantes

- O diretório `/api/cmd/api/` contém o arquivo `main.go`, responsável pelo
bootstrap da API. Ele inicia o servidor e, em caso de desligamento, o finaliza
*graciosamente*;
- O diretório `api/internal/auth/` contém a lógica relativa à autenticação de
usuários para o uso da API, que contempla a geração e checagem dos tokens de
acesso e atualização. Além disso, ele também contempla a checagem de rotas para
verificar se há necessidade de autenticação nelas;
- O diretório `api/internal/models/` contém as operações de disponíveis para o
uso no código. Seus arquivos acompanham as tabelas do banco, e maioria das
operações envolvem o banco de dados;
- O diretório `api/internal/server/` contém as rotas da API, a validação dos
parâmetros passados pelo usuário e as chamadas das operações disponíveis no
diretório `api/internal/models/`;
no código. Seus arquivos acompanham as tabelas do banco, e maioria das
operações envolvem o banco de dados;
- O diretório `api/internal/services/` contém o código de integração com
serviços externos, como o banco de dados (`api/internal/services/database/`) e
o servidor de email (`api/internal/services/email/`);
    - É interessante notar que os arquivos de script SQL para criação das
    tabelas, índices e triggers do banco, além da seed do mesmo, estão
    disponíveis na pasta `api/internal/services/database/source/` e no arquivo
    `api/internal/services/database/seed.sql`, respectivamente;
- O diretório `api/internal/utils/` contém variáveis e funções de uso geral por
toda o código, como  variáveis de ambiente e funções de log, de erro e de
checagem;
- O diretório `/api/tests/` contém todos os testes unitários da API, que podem
ser executados usando o comando `make test`.

### Arquivos de Configuração

- O arquivo `.env` é o **arquivo de configuração mais importante**, pois detém
os dados de acesso ao banco de dados e ao serviço de email, além dos segredos
para tokens JWT, pepper para as senhas criadas, porta da API e em qual ambiente
a API está sendo executada;
- O arquivo `fly.toml` contém toda a configuração do serviço de hospedagem
[fly][13], responsável por hospedar os ambientes de desenvolvimento,
homologação e produção;
- O arquivo `.air.toml` contém toda a configuração da ferramenta [air][4],
responsável pelo *hot reload* da API em caso de mudança;
- O arquivo `Makefile` contém todos os scripts da API, como os responsáveis por
criar a documentação, por construir o binário, etc;
- O diretório `.github` contém todos os workflows de CI da API, como o de
execução automática dos testes unitários e do teste de build.

### Diretórios Autogerados

- O diretório `api/docs/` é gerado ao executar o comando `make docs`, com base
nos comentários dos arquivos contidos no diretório `api/internal/server/`, e
diz respeito à documentação do [Swagger][2] contida na rota
`/api/docs/index.html`. Os comentários seguem as especificações da ferramenta
[swag][3], que é responsável por gerar a documentação;
- O diretório `tmp/` é gerado ao executar o comando `make watch`, e é utilizado
para armazenar arquivos de log relevantes para a ferramenta [air][4],
responsável pelo *hot reload* da API.


[1]: https://go.dev/
[2]: https://swagger.io/
[3]: https://github.com/swaggo/swag
[4]: https://github.com/cosmtrek/air
[5]: https://www.gnu.org/software/make/
[6]: https://www.docker.com/
[7]: https://go.dev/tour/
[8]: https://makefiletutorial.com/
[9]: https://docker-curriculum.com/
[10]: https://turso.tech/
[11]: https://mailtrap.io/
[12]: https://www.kirvano.com/
[13]: https://fly.io/
[14]: https://www.sqlite.org/

## Frontend

<!-- Pablo, por favor descreve aqui a estrutura do frontend -->

