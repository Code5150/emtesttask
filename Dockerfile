FROM golang:1.26-alpine AS builder

# Установка необходимых зависимостей для сборки
RUN apk add --no-cache git make gcc musl-dev

# Установка инструментов
RUN go install github.com/google/wire/cmd/wire@latest && \
    go install github.com/swaggo/swag/cmd/swag@latest

# Установка рабочей директории
WORKDIR /app

# Копирование файлов зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копирование исходного кода
COPY . .

# Генерация Swagger документации
# Замените ./cmd/api на путь к вашей точке входа
RUN swag init

# Генерация кода Wire
RUN wire

RUN ls -li && cd ./docs && ls -li && cd ..

# Сборка приложения
# CGO_ENABLED=0 для статической линковки
# -ldflags="-w -s" для уменьшения размера бинарного файла
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o /app/main \
    .

FROM alpine:3.24

# Создание непривилегированного пользователя
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Установка рабочей директории
WORKDIR /app

# Копирование скомпилированного бинарного файла
COPY --from=builder /app/main .

COPY --from=builder /app/docs ./docs

RUN ls -li && cd ./docs && ls -li && cd ..

# Копирование конфигурационных файлов (если есть)
# COPY --from=builder /app/configs ./configs
# COPY --from=builder /app/.env .

# Переключение на непривилегированного пользователя
USER appuser

# Открытие порта (измените на свой)
EXPOSE 8080

ENTRYPOINT ["./main"]