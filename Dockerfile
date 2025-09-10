# Этап сборки
FROM golang:1.22.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /cbrstats ./cmd/cbrstats

# Финальный образ
FROM alpine:3.20

WORKDIR /root/

COPY --from=builder /cbrstats .

# Точка входа
ENTRYPOINT ["./cbrstats"]
