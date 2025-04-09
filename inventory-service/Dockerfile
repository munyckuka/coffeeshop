# Сборка бинарника для Inventory Service
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копирование зависимостей и загрузка модулей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o inventory-service ./cmd/main.go

# Финальный образ
FROM alpine:latest
COPY --from=builder /app/inventory-service /inventory-service

EXPOSE 8080
ENTRYPOINT ["/inventory-service"]
