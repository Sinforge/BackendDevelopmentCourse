# Используем базовый образ Python
FROM python:3.9

# Устанавливаем переменные окружения
ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

# Устанавливаем директорию приложения в контейнере
WORKDIR /app

# Копируем зависимости проекта
COPY requirements.txt /app/

# Устанавливаем зависимости проекта
RUN pip install --no-cache-dir -r requirements.txt

# Копируем исходный код проекта
COPY . /app/
EXPOSE 8000
# Запускаем сервер Django
CMD ["python", "manage.py", "runserver", "0.0.0.0:8000"]