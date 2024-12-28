
# Dockerfile for Flask (Development)

FROM python:3.11-slim 

WORKDIR /app

# Install dependencies
COPY requirements.txt /app/
RUN pip install --no-cache-dir -r requirements.txt

# Copy application code
COPY . /app/

EXPOSE 8000

CMD ["sh", "-c", "python -m flask run --host=0.0.0.0 --port={{.Port}} --port=8000"]
