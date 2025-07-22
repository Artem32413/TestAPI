# Warehouse Management API

Микросервис для управления товарами на складах с аналитикой продаж

##  Технологический стек
- **Язык**: Go 1.24
- **База данных**: PostgreSQL 16+
- **Драйвер БД**: pgx/v5
- **Миграции**: go-migrate (CLI)
- **Логирование**: zap logger (JSON формат)
- **Документация**: Swagger/OpenAPI 3.0
- **Идентификаторы**: google/uuid

##  Архитектурные особенности
- Чистая архитектура (Clean Architecture)
- Принципы SOLID
- Graceful shutdown
- Стандартный HTTP-сервер Go

##  Быстрый старт

### Предварительные требования
- Docker 20.10+
- Docker Compose 2.0+
- Go 1.24

### Запуск проекта
```bash
# Сборка и запуск
docker-compose -f build/docker-compose.yml up --build

# Остановка
docker-compose -f build/docker-compose.yml down
```
### Запуск Swagger документации
При сборке образа проекта 
``` 
http://localhost:8080/docs/
```