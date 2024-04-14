# Сервис баннеров

# Общие вводные
**Баннер** — это документ, описывающий какой-либо элемент пользовательского интерфейса. Технически баннер представляет собой  JSON-документ неопределенной структуры. 
**Тег** — это сущность для обозначения группы пользователей; представляет собой число (ID тега). 
**Фича** — это домен или функциональность; представляет собой число (ID фичи).  
1. Один баннер может быть связан только с одной фичей и несколькими тегами
2. При этом один тег, как и одна фича, могут принадлежать разным баннерам одновременно
3. Фича и тег однозначно определяют баннер

# Getting Started
- Добавить .env файл в директорию с проектом и указать DB_PASSWORD
- Опционально настроить config под себя

# Usage

Запустить сервис можно с помощью команды `make build && make run`
Если приложение запускается впервые, необходимо применить миграции к базе данных `make migrate`


## Examples

Некоторые примеры запросов
**P.S** могут быть неточности/опечатки из-за того, что спешил к дедлайну
- [Регистрация](#sign-up)
- [Аутентификация](#sign-in)
- [Создание тега](#tag)
- [Создание фичи](#feature)
- [Получение пользовательского баннера](#userBanner)
- [Получение баннеров с фильтрацией](#getBanner)
- [Обновление баннера](#updateBanner)
- [Удаление баннера](#deleteBanner)
- [Создание баннера](#createBanner)

### Регистрация <a name="sign-up"></a>

Регистрация:
```curl
curl --location --request POST 'http://localhost:8000/auth/sign-up' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Zanzhit,
    "username":"zanzhit",
    "password":"12345",
    "role": "admin"
}'
```

Пример ответа:
```json
{
    "id": 1
}
```

### Аутентификация <a name="sign-in"></a>

Аутентификация для получения токена доступа:
```curl
curl --location --request POST 'http://localhost:8000/auth/sign-in' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"zanzhit",
    "password":"12345"
}'
```
Пример ответа:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzNTk3NzMsImlhdCI6MTcxMzEwMDU3Mywicm9sZSI6ImFkbWluIn0.jA3wBbYGGSfARtvrZJ7yoUeSFmOfQOOETG5_wVczLcw"
}
```

### Создание тега <a name="tag"></a>

Создание тега с указанным айди:
```curl
curl --location --request POST 'http://localhost:8000/entities/tag' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjY2MTUxMjEsImlhdCI6MTY2NjYwNzkyMSwiVXNlcklkIjoxfQ.c4jMWdmyXePtjTo_qrN6m9n-LQtHk_Q99OuzcpriYs4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1
}'
```

Пример ответа: 200

### Создание фичи <a name="feature"></a>

Создание фичи с указанным айди:
```curl
curl --location --request POST 'http://localhost:8000/entities/feature' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjY2MTUxMjEsImlhdCI6MTY2NjYwNzkyMSwiVXNlcklkIjoxfQ.c4jMWdmyXePtjTo_qrN6m9n-LQtHk_Q99OuzcpriYs4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1
}'
```

Пример ответа: 200

### Получение пользовательского баннера <a name="userBanner"></a>

Получение баннера по указанному тегу и фиче:
```curl
curl -X GET 'http://localhost:8000/userBanner?tag_id=1&feature_id=1' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzNTk3NzMsImlhdCI6MTcxMzEwMDU3Mywicm9sZSI6ImFkbWluIn0.jA3wBbYGGSfARtvrZJ7yoUeSFmOfQOOETG5_wVczLcw'
```

```json
{
    "title": "2Новый баннер",
    "image_url": "https://example.com/image.jpg",
    "description": "Описание баннера"
}
```

### Получение баннеров с фильтрацией по тегу и/или фиче <a name="getBanner"></a>

Получение баннеров по указанным тегу и/или фиче:
```curl
curl -X GET 'http://localhost:8000/entities/banner?tag_id=1&limit=1' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzNTk3NzMsImlhdCI6MTcxMzEwMDU3Mywicm9sZSI6ImFkbWluIn0.jA3wBbYGGSfARtvrZJ7yoUeSFmOfQOOETG5_wVczLcw'
```

Пример ответа:

```json
[
    {
        "id": 2,
        "content": {
            "title": "2Новый баннер",
            "image_url": "https://example.com/image.jpg",
            "description": "Описание баннера"
        },
        "tag": [
            1
        ],
        "feature": 3,
        "is_active": true,
        "created_at": "2024-04-14T13:20:24.008516Z",
        "updated_at": "2024-04-14T13:20:24.008516Z"
    }
]
```

### Обновление баннера <a name="updateBanner"></a>

Обновление баннера:
```curl
curl -X PATCH 'localhost:8000/entities/banner/2' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzNTk3NzMsImlhdCI6MTcxMzEwMDU3Mywicm9sZSI6ImFkbWluIn0.jA3wBbYGGSfARtvrZJ7yoUeSFmOfQOOETG5_wVczLcw'
--header 'Content-Type: application/json' \
--data-raw '{
    "Tag": [1,2]
}'
```

Пример ответа: 200

### Удаление баннера <a name="deleteBanner"></a>

Удаление баннера:
```curl
curl -X DELETE 'localhost:8000/entities/banner/2' \
-H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzNTk3NzMsImlhdCI6MTcxMzEwMDU3Mywicm9sZSI6ImFkbWluIn0.jA3wBbYGGSfARtvrZJ7yoUeSFmOfQOOETG5_wVczLcw'
```

Пример ответа: 200

### Создание баннера <a name="createBanner"></a>

Создание баннера:
```curl
curl --location --request POST 'localhost:8000/entities/banner' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": {
		"title": "2Новый баннер",
		"description": "Описание баннера",
		"image_url": "https://example.com/image.jpg"
	},
    "tag": [1, 2],
    "feature": 1,
    "is_active": true
}'
```

Пример ответа:
```json
{
    "id": 1
}
```
