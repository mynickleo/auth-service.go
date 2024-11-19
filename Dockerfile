FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth-service ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add libc6-compat curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz -C /usr/local/bin

WORKDIR /root/

COPY --from=builder /app/auth-service .

COPY .env .env

COPY ./internal/database/postgres/migrations ./internal/database/postgres/migrations

EXPOSE 3000

CMD ["./auth-service"]
