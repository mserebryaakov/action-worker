# Сервис-коннектор к потоку RabbitMQ

## Docker

    docker build . -t action-worker:latest

    docker run --env-file ./.env action-worker:latest

## Переменные окружения (* - обязательное)

    DEBUG - режим логирования (true/false), если отсутствует - false

    WORKERFREQUENCY* - частота работы worker (в секундах)

    RABBITURL* - url к REST потоку RabbitMQ

    RABBITROUTECODE* - логин пользователя к REST потоку RabbitMQ

    RABBITROUTEPASS* - пароль пользователя к REST потоку RabbitMQ

    ELMAURL* - url к стенду ELMA

    ELMATOKEN - токен к стенду ELMA
