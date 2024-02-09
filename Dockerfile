FROM golang:1.22.0-alpine3.19 as env

WORKDIR /app

# Instalar Swag e Air
RUN go install -v github.com/swaggo/swag/cmd/swag@latest && go install -v github.com/cosmtrek/air@latest

# Atualiza e instala dependências
RUN apk update && apk add -U --no-cache --progress curl make nodejs npm sqlite

# Node e TailwindCSS
RUN npm install
