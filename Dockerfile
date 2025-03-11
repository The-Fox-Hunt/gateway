# 🔹 1. Стадия сборки (builder)
FROM golang:1.24-alpine AS builder

# Добавляем сертификаты для HTTPS (например, для PostgreSQL)
RUN apk add --no-cache ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей и устанавливаем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код проекта
COPY . .

# Собираем бинарник
RUN go build -o /app/gateway cmd/service/main.go

# 🔹 2. Финальный контейнер (чистый и минимальный)
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарник из builder-стадии (без лишних файлов)
COPY --from=builder /app/gateway /app/gateway

# Делаем бинарник исполняемым
RUN chmod +x /app/gateway

# Запускаем приложение
CMD ["/app/gateway"]
