# Используем минимальный образ Go
FROM golang:1.24-alpine

# Добавляем сертификаты для HTTPS (например, для PostgreSQL)
RUN apk add --no-cache ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Устанавливаем зависимости
RUN go mod download

# Собираем бинарник
RUN go build -o gateway cmd/service/main.go

# Запускаем приложение
CMD ["/app/gateway"]
