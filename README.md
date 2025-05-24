# Quotation API Service
REST API для управления цитатами. Позволяет создавать, получать, удалять цитаты.

## 🛠 Технологии
- Go 1.24 (чистая архитектура)
- Docker (контейнеризация приложения)
- Swagger (документация API)

## Установка и запуск проекта

### 1. Клонирование репозитория
```bash
git clone https://github.com/voronkov44/rest_go_quotation.git
```
### 2. Переход в корневую директорию 
```bash
cd rest_go_quotation
```

### 3. Сборка контейнера и запуск приложения
*Требуется установка [docker](https://www.docker.com/products/docker-desktop/), если не установлен, смотрите [зависимости.](https://github.com/voronkov44/rest_go_quotation?tab=readme-ov-file#%D1%83%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BA%D0%B0-docker)*

Сборка контейнера:
```bash
docker build -t quotation-book:v1 .
```
Запуск контейнера:
```bash
docker run -d -p 8081:8081 --name quotation-book quotation-book:v1
```

Сервер будет доступен на http://localhost:8081

### Общие команды для работы с контейнером

```
# Просмотр запущенных контейнеров
docker ps

#Просмотр всех контейнеров, включая остановленные
docker ps -a

# Остановка контейнера
docker stop <container_id>

# Запуск остановленного контейнера
docker start <container_id>

# Удалание контейнера
docker rm <container_id>

# Удаление образа
docker rmi <image_id>

# Просмотр логов контейнера
docker logs -f <container_name>

# Вход в контейнер
docker exec -it <container_name> sh - sh

docker exec -it <container_name> /bin/bash - bash

# Очистка системы
docker system prune
```

## 📚 Документация API

### После запуска сервера откройте Swagger UI:

```
http://localhost:8081/swagger/index.html
```

### Пример запросов
Запросы:
(POST) Для создания заметки:
```
curl -X POST http://localhost:8081/quotes \
-H "Content-Type: application/json" \
-d '{"author":"Confucius", "quote":"Life is simple, but we insist on making it complicated."}'
```

(GETALL) Для получения всех заметок:
```
curl http://localhost:8081/quotes
```

(RANDOM) Для получения рандомной заметки:
```
curl http://localhost:8081/quotes/random
```

(AUTHOR) Для получения заметки по автору:
(Был создан отдельный эндпоинт, для того чтобы получить заметку по автору, воспользуйтесь http://localhost:8081/quotes/author/{author})
```
curl http://localhost:8081/quotes/author/Confucius
```

(DELETE) Для удаления заметки по ID:
```
curl -X DELETE http://localhost:8081/quotes/1
```
*Цитаты хранятся в оперативной памяти. После перезапуска приложения все данные обнуляются.*

## Зависимости
### Установка docker
Установка пакета [Docker Engine](https://docs.docker.com/engine/install/)
