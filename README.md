# Arithmetic Expression Web Service

## Описание проекта

Этот проект представляет собой веб-сервис для вычисления арифметических выражений. Пользователь отправляет арифметическое выражение в теле HTTP POST-запроса и получает в ответ результат вычисления или сообщение об ошибке.

### Точка доступа (Endpoint)
- **URL**: `/api/v1/calculate`
- **Метод**: POST
- **Content-Type**: `application/json`

#### Формат запроса
```json
{
  "expression": "выражение, которое ввёл пользователь"
}
```

#### Формат успешного ответа (код 200):
```json
{
  "result": "результат выражения"
}
```

#### Формат ответа при некорректном выражении (код 422):
```json
{
  "error": "Expression is not valid"
}
```

#### Формат ответа при внутренней ошибке сервера (код 500):
```json
{
  "error": "Internal server error"
}
```

---

## Примеры использования

### 1. Успешный запрос
**Запрос**:
```bash
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

**Ответ**:
```json
{
  "result": "6"
}
```

### 2. Некорректное выражение (код 422)
**Запрос**:
```bash
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*a"
}'
```

**Ответ**:
```json
{
  "error": "Expression is not valid"
}
```

### 3. Внутренняя ошибка сервера (код 500)
**Запрос**:
```bash
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "somethingDiff": "what?"
}'
```

**Ответ**:
```json
{
  "error": "Internal server error"
}
```

---

## Запуск проекта

Для запуска веб-сервиса выполните следующую команду:
```bash
go run ./cmd/calc_service/...
```

После запуска сервис будет доступен на `http://localhost:8080` по указанному эндпоинту `/api/v1/calculate`.

---

## Технологии
- **Язык**: Go
- **Фреймворк**: `net/http`
- **Формат данных**: JSON

---

## Структура проекта
```
calc_go/
├── cmd/
│   └── calc_service/
│       └── main.go
├── internal/
│   └── application/ 
│       └── application.go 
├── pkg/
│   └── calc/
│       ├── calc_test.go
│       └── calc.go
└── README.md 
```

---

## Дополнительная информация
Проект принимает арифметические выражения, содержащие следующие символы:
- **Цифры**: `0-9`
- **Операторы**: `+`, `-`, `*`, `/`
- **Скобки**: `(` и `)`

Любые другие символы будут приводить к ошибке 422 с сообщением об ошибке.

---

## Автор
vizurth
