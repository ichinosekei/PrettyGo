
# README

## Описание проекта

Проект состоит из двух приложений:

1. **HTTP-сервер**: Веб-сервер с тремя запросами:
    - `GET /version`: Возвращает версию API в формате semantic versioning (v1.0.0).
    - `POST /decode`: Принимает строку, закодированную в формате Base64, в формате JSON и декодирует её.
    - `GET /hard-op`: Симулирует длительную операцию (от 10 до 20 секунд) с вероятностью вернуть либо `500 Internal Server Error`, либо `200 OK`.

2. **HTTP-клиент**: Клиент, который последовательно вызывает указанные выше запросы сервера. Последний запрос (`/hard-op`) отменяется, если его выполнение занимает более 15 секунд, используя пакет `context`.

## Структура проекта

```
hw2/
├── client/
│   ├── client.go          // Основной клиентский код
│   ├── decode.go          // Метод для POST запроса на декодирование base64
│   ├── hardop.go          // Метод для GET запроса на выполнение сложной операции
│   ├── version.go         // Метод для GET запроса версии сервера
│
├── server/
│   ├── main.go            // Сервер, предоставляющий эндпоинты
│
├── go.mod                 // Модуль Go
├── README.md              // Описание проекта и инструкции
    
```

## Настройка


### 1. Инициализация модулей Go

Для сервера и клиента необходимо инициализировать модули Go для управления зависимостями:

**Для сервера:**

```bash
cd server
go mod init server
```

**Для клиента:**

```bash
cd ../client
go mod init client
```

Это создаст файлы `go.mod` в обеих директориях.

### 2. Запуск сервера

1. Перейдите в директорию `server`:
   ```bash
   cd server
   ```

2. Запустите сервер:
   ```bash
   go run main.go
   ```

Сервер будет запущен и будет слушать порт 8080. Вы увидите следующее сообщение в терминале:
```
Server is running on port 8080...
```
### 3. Клиент инцилизируется файлом main.go лежащем в корне проекта

### 4. Запуск клиента и работа с ним

1. Откройте новый терминал и перейдите в директорию `client`:
   ```bash
   cd hw2
   ```

2. Запустить инциализацию клиента:
   ```bash
   go run main.go
   ```
Клиент использует структура Client, который инициализируется с параметрами схемы, хоста и порта. Пример инициализации:
```go
client := client.NewClient("http", "localhost", "8080")
```
Клиент поддерживает следующие методы:

GetVersion() — выполняет GET запрос на /version и возвращает версию сервера:
```go
version, err := client.GetVersion()
if err != nil {
    fmt.Println("Error fetching version:", err)
} else {
    fmt.Println("Server version:", version)
}
```
DecodeString() — выполняет POST запрос на /decode, отправляя строку в формате base64 для декодирования:
```go
decoded, err := client.DecodeString("SGVsbG8sIFdvcmxkIQ==")
if err != nil {
    fmt.Println("Error decoding string:", err)
} else {
    fmt.Println("Decoded string:", decoded)
}

```
HardOp() — выполняет GET запрос на /hard-op, эмулируя сложную операцию:
```go
err := client.HardOp()
if err != nil {
    fmt.Println("Error during hard operation:", err)
} else {
    fmt.Println("Hard operation completed successfully")
}

```

- `GET /version`: Получает и выводит версию API.
- `POST /decode`: Декодирует строку Base64 и выводит результат.
- `GET /hard-op`: Симулирует длительную операцию и выводит успех или неудачу. Если запрос выполняется дольше 15 секунд, клиент отменит его и выведет сообщение о таймауте.

### 5. Graceful shutdown

Вы можете остановить сервер, отправив сигнал `SIGTERM` или `SIGINT` (например, нажав `Ctrl+C`). Сервер попытается корректно завершить работу, дав возможность завершить текущие запросы в течение 5 секунд перед остановкой.

## Тестирование запросов сервера

Вы можете вручную протестировать сервер с помощью команды `curl` или любого HTTP-клиента (например, Postman).

### 1. `GET /version`

Получение версии API:

```bash
curl http://localhost:8080/version
```

Ожидаемый результат:
```
v1.0.0
```

### 2. `POST /decode`

Отправляем строку, закодированную в Base64, и получаем оригинальную строку:

```bash
curl -X POST http://localhost:8080/decode \
     -H "Content-Type: application/json" \
     -d '{"inputString":"SGVsbG8sIFdvcmxkIQ=="}'
```

Ожидаемый результат (декодированная строка):
```json
{"outputString": "Hello, World!"}
```

### 3. `GET /hard-op`

Симулирует длительную операцию (10–20 секунд). Ответ может быть либо `200 OK`, либо `500 Internal Server Error`:

```bash
curl http://localhost:8080/hard-op
```

Ожидаемые результаты:
- Если успешно:
  ```
  Operation completed successfully
  ```
- Если произошла ошибка:
  ```
  Internal Server Error
  ```

## Вывод клиента

Когда вы запустите клиента, он выполнит следующие действия:

1. **Получение версии API**: Вывод версии (например, `v1.0.0`).
2. **Декодирование строки Base64**: Вывод декодированной строки (например, `Hello, World!`).
3. **Симуляция длительной операции**: Выводит успех или таймаут для запроса к `/hard-op`.

Пример вывода:
```
API Version: v1.0.0
Decoded String: Hello, World!
Request to /hard-op succeeded, Status Code: 200
```

Если запрос к `/hard-op` был отменён по таймауту:
```
API Version: v1.0.0
Decoded String: Hello, World!
Request to /hard-op timed out
```

## Заключение

Этот проект демонстрирует:
- Простой HTTP-сервер.
- Плавное завершение работы сервера (graceful shutdown).
- Обработку таймаута в клиенте с использованием пакета `context`.
