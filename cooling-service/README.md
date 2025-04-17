# Cooling Service API

🧊 Микросервис для реализации механизма периода охлаждения по кредитным договорам.

## 📦 Структура проекта

```
cooling-service/
├── cmd/server/               # Точка входа (main.go)
├── internal/api/             # Роутинг и HTTP-интерфейс
├── internal/cooling/         # Бизнес-логика
├── internal/storage/         # PostgreSQL-хранилище
├── db/migrations/            # SQL-миграции
├── docs/                     # Swagger/OpenAPI
├── helm/                     # Helm Chart для Kubernetes
├── .kube/                    # Пример Kube config
├── .gitlab/                  # GitLab CI/CD шаблоны
├── docker-compose.yml        # Compose c PostgreSQL
├── Dockerfile                # Docker-образ
├── README.md                 # Инструкция
```

---

## 🚀 Быстрый старт

### 🔧 Запуск с PostgreSQL

```bash
docker-compose up --build
```

- API: [http://localhost:8080](http://localhost:8080)
- PostgreSQL: `localhost:5432` (user: coolinguser / pass: coolingpass)

---

## 📚 Swagger UI

Открой файл `docs/swagger.html` в браузере или подключи OpenAPI из `docs/openapi.yaml`.

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

## 📄 OpenAPI

- Файл спецификации: `docs/openapi.yaml`
- Swagger UI: `docs/swagger.html`

---

## ⚙️ CI/CD

GitLab Pipeline (.gitlab-ci.yml):

- prepare → test → analyze → build → deploy
- интеграция с: Nexus, SonarQube, Kubernetes (Dev/Test/Prod)

---

## 📂 База данных

PostgreSQL таблица: `cooling_periods`  
Создаётся автоматически при старте сервиса (см. `RunMigrations()`)

---

## 🧭 Kubernetes и Helm

- Helm Chart: `helm/cooling-service/`
- Манифесты: `k8s/`
- Поддержка ingress + кастомизация через values.yaml

---

## 📋 Документация

- BRD: Бизнес-требования
- FSD: Функциональные требования с BPMN-диаграммой