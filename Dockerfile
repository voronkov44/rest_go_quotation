FROM golang:1.24-alpine AS builder

# Установим рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем модули
RUN go mod download

# Копируем оставшийся код проекта в контейнер
COPY . .

# Собираем бинарник приложения
RUN go build -o quotation-book ./cmd

# Указываем порт
EXPOSE 8081

# Команда для запуска контейнера
CMD ["./quotation-book"]
