### Эндпоинты API

#### `GET api/counter`

Возвращает текущее значение счетчика.

#### `DELETE api/clear`

Очищает счетчик.

### 1. Сборка образа Docker

В корневом каталоге проекта выполните:

```bash
make build
or
docker build -t myapp .
```

### 2. Запуск контейнера

```
make run
or
docker run -d -p 8080:8080 --name myapp-container myapp
```

### 3. Доступ к Swagger UI

```
localhost:8080/swagger/index.html
```

### Остановка, Удаление контейнера

```
make stop
or
docker stop myapp-container



make clear
or
docker rm myapp-container





```
