# syntax=docker/dockerfile:1

FROM golang:1.22

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копируем модули
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Собираем Go-приложение
RUN go build -o main .

# Запускаем бинарник
CMD ["./main"]
