# Сборка бинарника для Order Service
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копирование зависимостей и загрузка модулей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения с указанием нужной архитектуры (arm64)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o order-service ./cmd/main.go

# Финальный образ на основе Alpine
FROM alpine:latest
COPY --from=builder /app/order-service /order-service

EXPOSE 8081
ENTRYPOINT ["/order-service"]
