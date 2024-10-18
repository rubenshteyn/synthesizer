#!/bin/bash

# Поднимаем Docker контейнеры
docker-compose up -d

# Ждем, пока база данных запустится
echo "Ожидание запуска PostgreSQL..."
sleep 10

# Создаем таблицы и добавляем пользователя
psql -h localhost -U synthesizer_admin -d synthesizer -c "
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    role VARCHAR(50),
    password VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS monomers (
    name VARCHAR(100) PRIMARY KEY,
    count INT
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    sequence INT,
    user_id INT REFERENCES users(id),
    status VARCHAR(50),
    priority INT
);
"

# Добавляем пользователя
psql -h localhost -U synthesizer_admin -d synthesizer -c "
INSERT INTO users (name, password, role) VALUES ('Admin', '1234', 'admin');
"

# Запускаем фронтенд
echo "Запуск фронтенда..."
(cd client && npm run start &)

# Запускаем бэкенд
echo "Запуск бэкенда..."
(go run cmd/main.go &)

echo "Все сервисы запущены!"
