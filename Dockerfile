# Build stage for API
FROM golang:1.22.1-bookworm as api-builder

ENV CGO_ENABLED=1

WORKDIR /app

COPY api/ .
COPY .env .

RUN apt-get -y update && \
    apt-get install -y --no-install-recommends git make sqlite3 && \
    rm -rf /var/lib/apt/lists/* && \
    go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init --generalInfo routes.go --parseDependency --parseInternal --dir internal/server && \
    go build --ldflags "-s -w" -o /go/bin/main ./cmd/api/main.go

# Build stage for Frontend
FROM node:20 as frontend-builder

WORKDIR /app

COPY frontend/package*.json .

RUN npm install

COPY frontend/ .

RUN npm run staging

# Final stage
FROM nginx:bookworm as exec

WORKDIR /app

# Copy API binary
COPY .env .
COPY --from=api-builder /go/bin/main .

# Copy frontend build
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=frontend-builder /app/dist/redirectfy-frontend/* /usr/share/nginx/html/

# Install nginx
RUN apt-get update && apt-get install -y ca-certificates

# Expose ports                                   
EXPOSE 8080 80

# # Start nginx and API
ENTRYPOINT ["sh", "-c", "/app/main & nginx -g 'daemon off;'"]
