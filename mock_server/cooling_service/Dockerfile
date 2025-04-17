# Используем минимальный Python-образ
FROM python:3.10-slim

# Установка зависимостей
RUN pip install flask flask-swagger-ui

# Копируем файлы приложения
WORKDIR /app
COPY mock_cooling_service_with_ui.py ./mock_cooling_service.py
COPY CoolingService_API_with_errors.yaml ./openapi.yaml

# Запуск сервера
CMD ["python", "mock_cooling_service.py"]