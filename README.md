# Тестовое задание на Lead Golang Developer

Вам необходимо написать 2 приложения - gRPC клиент и gRPC сервер.

Сервер должен принимать запрос от клиента `Fetch(url string)`, где `url` - ссылка на API endpoint сервиса, по которому можно получить CSV файл. При получении этого запроса, он должен загрузить по указанной ссылке актуальную информацию в формате .csv файла и выполнить создание \ обновление в БД MongoDB информации о продукции.

Второй gRPC метод `List(<paging params>, <sorting params>)` принимает кастомные структуры с параметрами для сортировки и пагинации и возвращает список продуктов.

## Покрытие тестами

| Package | Coverage |
|---|---|
| Total | 65,0% |
| Config | 61,9% |
| Server | 47,1% |
| Services | 81,8% |

### **Команды тестирования**

Сделать тесты и записать результаты в cover.out

```go
go test -coverprofile=cover.out -v ./...
```

Проверить покрытие тестами

```go
go tool cover -func=cover.out
```

## Генерация grpc

```bash
protoc --go_out=. --go-grpc_out=. transport.proto
```

## Запуск

Запуск клиента

```bash
go run cmd/client/main.go
```

Запуск сервера

```bash
go run cmd/server/main.go
```

## Пример .env файла

```env
# MONGO CONFIG
MONGO_DBNAME=products
MONGO_COLLECTIONNAME=products
MONGO_HOST=localhost
MONGO_PORT=27017

# GRPC CONFIG
GRPC_PORT=9000
```
