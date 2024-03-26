FROM golang:1.22.1 as env

ENV CGO_ENABLED=1

WORKDIR /app

# Atualiza e instala dependências
RUN apt-get -y update && apt-get install -y --no-install-recommends git make nodejs npm sqlite3 && rm -rf /var/lib/apt/lists/*

# Instalar Swag e Air
RUN go install github.com/swaggo/swag/cmd/swag@latest && go install github.com/cosmtrek/air@latest && go install github.com/a-h/templ/cmd/templ@latest
