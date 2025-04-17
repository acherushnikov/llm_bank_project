# Используем официальный Go образ
FROM golang:1.20-alpine

# Установка рабочей директории
WORKDIR /app

# Копируем весь проект
COPY . .

# Сборка приложения
RUN go build -o cooling-service ./cmd/server

# Открываем порт
EXPOSE 8080

# Запуск приложения
CMD ["./cooling-service"]