TO RUN USE: "go run ./cmd/blog-api"

TO CHECKING HANDLERS USE folder "./rest-client"

Задача: Разработать API блога с использованием Go, Chi router, SQLite (modernc.org/sqlite) и миграций.

Требования:

1.Использовать Chi для маршрутизации
2.Использовать modernc.org/sqlite для работы с SQLite
3.Реализовать миграции с помощью golang-migrate
4.Реализовать CRUD операции для постов блога:
  - Получение всех постов (GET /posts)
  - Получение конкретного поста (GET /posts/{id})
  - Создание нового поста (POST /posts)
  - Обновление поста (PUT /posts/{id})
  - Удаление поста (DELETE /posts/{id})


5.Использовать JSON для обмена данными
6.Обрабатывать ошибки и возвращать соответствующие HTTP-статусы
7.Реализовать базовое логирование запросов

Дополнительные задачи (по желанию):

1.Добавить валидацию входных данных
2.Реализовать пагинацию для списка постов
3.Добавить поиск по заголовку или автору

структура posts

CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    author TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
