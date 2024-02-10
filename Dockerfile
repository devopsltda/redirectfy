FROM golang:1.22.0-alpine3.19 as env

ENV CGO_ENABLED=1

WORKDIR /app

# Atualiza e instala dependências
RUN apk update && apk add -U --no-cache --progress bash curl gcc gcompat git make musl-dev nodejs npm sqlite

# Instalar Swag e Air
RUN go install github.com/swaggo/swag/cmd/swag@latest && go install github.com/cosmtrek/air@latest && go install github.com/a-h/templ/cmd/templ@latest
