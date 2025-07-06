
# Dockerfile for FastAPI (Production)

FROM python:3.11-slim as builder

WORKDIR /app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . .

# Final production image
FROM python:3.11-slim

WORKDIR /app

# Set environment variables
ENV PATH="/venv/bin:$PATH"

# Copy virtual environment and application from the builder stage
COPY --from=builder /usr/local/lib/python3.11/site-packages /usr/local/lib/python3.11/site-packages
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /app /app

EXPOSE 8000

CMD ["sh", "-c", "uvicorn app:app --host=0.0.0.0 --port=8000"]
