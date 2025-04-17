# Cooling Service API

🧊 Микросервис для реализации механизма периода охлаждения по кредитным договорам.

## 📦 Структура проекта

```
cooling-service/
├── cmd/server/               # Точка входа (main.go)
├── internal/api/             # Роутинг и HTTP-интерфейс
├── internal/cooling/         # Бизнес-логика периода охлаждения
├── docs/                     # Swagger UI и OpenAPI спецификация
├── Dockerfile                # Docker-образ
├── docker-compose.yml        # Композиция контейнеров
├── go.mod                    # Go-модуль
└── README.md                 # Инструкция
```

## 🚀 Быстрый старт

### 🔧 Запуск в Docker

```bash
docker-compose up --build
```

Сервис будет доступен по адресу: [http://localhost:8080](http://localhost:8080)

### 📚 Swagger UI

Открой `docs/swagger.html` в браузере для просмотра документации.

---

## 📌 Основные эндпоинты

| Метод | Путь                   | Описание |
|-------|------------------------|----------|
| POST  | `/cooling/register`    | Регистрация кредита |
| GET   | `/cooling/validate`    | Проверка периода охлаждения |
| POST  | `/cooling/pay`         | Отметка возврата суммы |
| POST  | `/cooling/withdraw`    | Подача заявления на отказ |
| GET   | `/cooling/status`      | Проверка статуса отказа |
| GET   | `/cooling/report`      | Отчёт по отказам |

---

## 🧪 Тестирование

### 🧱 Unit-тесты
```bash
go test ./internal/cooling
```

### 🔄 Интеграционные тесты
```bash
go test ./internal/api
```

---

## 📄 OpenAPI спецификация

Файл: `docs/openapi.yaml`  
Swagger UI: `docs/swagger.html`

---

## ✍️ Авторы и лицензия

Разработано в рамках архитектурного проектирования микросервисов (2025). Лицензия MIT.