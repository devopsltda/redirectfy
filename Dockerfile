FROM golang:1.22.0-alpine3.19 as env

WORKDIR /app

# Instalar Swag e Air
RUN go install github.com/swaggo/swag/cmd/swag@latest && go install github.com/cosmtrek/air@latest && go install github.com/a-h/templ/cmd/templ@latest

# Atualiza e instala dependências
RUN apk update && apk add -U --no-cache --progress bash curl make nodejs npm sqlite
