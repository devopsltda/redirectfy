FROM golang:1.22.0-alpine3.19

WORKDIR app

RUN apk add sqlite

COPU

CMD ["/app/bin/main"]
