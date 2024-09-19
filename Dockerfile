# Используем базовый образ Golang для сборки приложения
FROM golang:1.22.5-alpine AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем все файлы проекта в рабочую директорию
COPY . .

# Устанавливаем зависимости и собираем приложение
RUN go mod download
RUN go build -o crpt ./cmd/crpt

# Используем базовый образ Alpine для запуска приложения
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение из предыдущего этапа
COPY --from=build /app/crpt .

# Копируем файл конфигурации
COPY config.yaml .

# Устанавливаем точку входа
ENTRYPOINT ["./crpt"]
