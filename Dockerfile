# Используем официальный образ PostgreSQL
FROM postgres:latest

# Устанавливаем переменные окружения для PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password123
ENV POSTGRES_DB=restapi_dev

# Копируем файлы миграций в контейнер
COPY migrations/ /docker-entrypoint-initdb.d/

# Устанавливаем рабочую директорию (не обязательно, но это может быть полезно для вашего проекта)
WORKDIR /app

# Докер-образ PostgreSQL автоматически выполнит все SQL-скрипты в папке /docker-entrypoint-initdb.d/
# при первом запуске контейнера, в том числе ваш скрипт для создания таблиц и т.д.