
# Dockerfile for Flask (Production)

FROM python:3.11-slim as builder

WORKDIR /app

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies in a virtual environment
RUN python -m venv /venv
ENV PATH="/venv/bin:$PATH"

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

# Copy application code
COPY . .

# Final production image
FROM python:3.11-slim

WORKDIR /app

# Set environment variables
ENV PATH="/venv/bin:$PATH"

# Copy virtual environment and application from the builder stage
COPY --from=builder /venv /venv
COPY --from=builder /app /app

EXPOSE 8000

CMD ["sh", "-c", "gunicorn run:app --bind 0.0.0.0:8000"]
