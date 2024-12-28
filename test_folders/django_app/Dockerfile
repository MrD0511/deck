
# Dockerfile for Django (Development)

FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt /app

RUN pip install -r requirements.txt

COPY . .

EXPOSE 8000

CMD ["sh", "-c", "python hello/manage.py runserver 0.0.0.0:8000"]
