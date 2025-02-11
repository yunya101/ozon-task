# Используем официальный образ Go
FROM golang:1.20-alpine AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы исходного кода в контейнер
COPY . .

# Загружаем зависимости и собираем приложение
RUN go mod tidy
RUN go build -o main ./cmd

# Используем более легкий образ для запуска
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарник из builder
COPY --from=builder /app/main .

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
