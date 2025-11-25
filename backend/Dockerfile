# Stage 1: Сборка
FROM golang:1.23 AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./


# Копируем ВСЕ исходники 
COPY . .

# Собираем бинарник из ./app/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o lms-backend ./backend/app/main.go

# Stage 2:
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Копируем бинарник
COPY --from=builder /app/lms-backend .

# Копируем миграции 
COPY --from=builder  migrations/ /migrations/

# Порт по умолчанию
EXPOSE 8081

# Запуск
CMD ["./lms-backend"]