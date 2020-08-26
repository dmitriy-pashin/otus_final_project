# Анти-брутфорс

## Установка
Запуск БД в контейнере c сохранением состояния базы
```bash
docker run -d --name pgsql \
-e POSTGRES_PASSWORD=postgres \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v !!!path_to_your_folder!!!:/var/lib/postgresql/data \
-p 5432:5432 \
postgres
```

Созадние базы и юзера

```bash
CREATE USER otus WITH ENCRYPTED PASSWORD 'otus';

CREATE DATABASE antibf;

GRANT ALL PRIVILEGES ON DATABASE antibf TO otus;
```

 Миграции

```bash
make migrate-up
```

* build

```bash
make build
```

## Запуск

Запуск веб приложения

```bash
make run-web
```
Запуск тестов

```bash
make tests
```

## Folders

```
.
├── bin (результаты компиляции). Под gitignore
├── config - настройки приложения
└── src
    ├── component
    │   ├── app  -виды приложений для запуска. Реализован только Web, Console не успел
    │   ├── config  - парсинг настроек прилоежения
    │   ├── db  - компонент работы с БД
    │   ├── di  - настройка зависимостей для проекта
    │   │   └── definitions
    ├── controller 
    ├── repository
    ├── service


```