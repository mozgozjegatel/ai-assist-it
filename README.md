# Todo API Service - REST API для управления пользователями и их задачами на Go с автоматической Swagger-документацией.
## Описание

Проект представляет собой REST API сервис для управления пользователями и их задачами (Todo list). Разработан на Go с использованием роутера `chi` и полностью документирован через Swagger/OpenAPI.

## Структура проекта

```
api-doc-example/
├── main.go                     # Точка входа, инициализация сервера и роутов
├── go.mod                      # Зависимости проекта
├── go.sum                      # Контрольные суммы зависимостей
├── docs/                       # Сгенерированная документация Swagger
│   ├── docs.go                 # Вспомогательный файл
│   ├── swagger.json            # Основная документация в JSON
│   └── swagger.yaml            # Документация в YAML
├── internal/
│   ├── handlers/               # HTTP-обработчики (контроллеры)
│   │   ├── user.go             # Обработчики для пользователей
│   │   └── todo.go             # Обработчики для задач
│   └── models/                 # Модели данных и DTO
│       ├── user.go             # Модели пользователя
│       └── todo.go             # Модели задачи
└── README.md                   # Документация проекта
```

## Эндпоинты для пользователей
```
Метод	Эндпоинт	Описание
GET     /users	    Получить всех пользователей
POST	/users	    Создать нового пользователя
GET     /users/{id}	Получить пользователя по ID
PUT     /users/{id}	Обновить данные пользователя
DELETE	/users/{id}	Удалить пользователя
```

## Эндпоинты для задач
```
Метод       Эндпоинт        Описание
GET         /todos          Получить все задачи
POST        /todos          Создать новую задачу
GET         /todos/{id}     Получить задачу по ID
PUT         /todos/{id}     Обновить задачу
DELETE      /todos/{id}     Удалить задачу
```

### Основные возможности

* 👤 **Управление пользователями**: CRUD операции для пользователей
* 📝 **Управление задачами**: Создание, чтение, обновление и удаление задач
* 📖 **Swagger UI**: Интерактивная документация API
* 🧩 **Чистая архитектура**: Разделение на слои (handlers, models)
* 💾 **In-Memory хранилище**: Быстрое прототипирование (легко заменить на БД)
* 🔒 **Валидация**: Базовая валидация входных данных

## Установка и запуск
1.  **Клонируйте репозиторий:**
    ```bash
    git clone https://github.com/mozgozjegatel/ai-assist-it.git
    cd ai-assist-it
    go mod download
    go install github.com/swaggo/swag/cmd/swag@latest
    # Генерации swagger документации
    swag init -g ./main.go -o ./docs
    # Запуск сервиса
    go run main.go

### Требования
- Go версии 1.23 или выше
- Git (для клонирования)
- chi v5.0.10 или выше - Маршрутизатор HTTP-запросов
- http-swagger v2.0.2 или выше - Middleware для Swagger UI
- swaggo/swag v1.16.2 или выше Генератор Swagger-документации
- swaggo/files v1.0.1 или выше Встраивание статических файлов Swagger

## Примеры запросов и ответов

1. **Создание пользователя (POST /users)**

    Запрос:

    ```bash
    curl -X POST http://localhost:8080/api/v1/users \
    -H "Content-Type: application/json" \
    -d '{"email":"ivan@example.com","name":"Иван Петров"}'
    ```

    Успешный ответ (201 Created):

    ```json
    {
    "success": true,
    "data": {
        "id": 1,
        "email": "ivan@example.com",
        "name": "Иван Петров",
        "created_at": "2024-03-20T12:00:00Z"
    }
    }
    ```

2. **Создание задачи (POST /todos)**

    Запрос:

    ```bash
    curl -X POST http://localhost:8080/api/v1/todos \
    -H "Content-Type: application/json" \
    -d '{"user_id":1,"title":"Купить продукты","description":"Молоко, хлеб, яйца"}'
    ```

    Успешный ответ (201 Created):

    ```json
    {
    "success": true,
    "data": {
        "id": 1,
        "user_id": 1,
        "title": "Купить продукты",
        "description": "Молоко, хлеб, яйца",
        "done": false,
        "created_at": "2024-03-20T12:05:00Z",
        "updated_at": "2024-03-20T12:05:00Z"
    }
    }
    ```
3. **Получение списка задач (GET /todos)**

    Запрос:

    ```bash
    curl http://localhost:8080/api/v1/todos
    ```

    Успешный ответ (200 OK):

    ```json
    {
    "success": true,
    "data": [
        {
        "id": 1,
        "user_id": 1,
        "title": "Купить продукты",
        "description": "Молоко, хлеб, яйца",
        "done": false,
        "created_at": "2024-03-20T12:05:00Z",
        "updated_at": "2024-03-20T12:05:00Z"
        }
    ]
    }
    ```
4. **Обновление задачи (PUT /todos/1)**

    Запрос:

    ```bash
    curl -X PUT http://localhost:8080/api/v1/todos/1 \
    -H "Content-Type: application/json" \
    -d '{"done": true}'
    ```
    Успешный ответ (200 OK):

    ```json
    {
    "success": true,
    "data": {
        "id": 1,
        "user_id": 1,
        "title": "Купить продукты",
        "description": "Молоко, хлеб, яйца",
        "done": true,
        "created_at": "2024-03-20T12:05:00Z",
        "updated_at": "2024-03-20T12:10:00Z"
    }
    }
    ```
5. **Ошибка (404 Not Found)**

    Ответ:

    ```json
    {
    "error": "not_found",
    "message": "Задача не найдена"
    }
