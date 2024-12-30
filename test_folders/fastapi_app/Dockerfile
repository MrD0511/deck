# Stage 1: Builder
FROM python:3.11-slim AS builder

WORKDIR /app

# Install dependencies only
COPY requirements.txt /app/
RUN pip install --no-cache-dir -r requirements.txt

# Stage 2: Final (Production)
FROM python:3.11-slim

WORKDIR /app

# Copy only necessary artifacts from the builder stage
COPY --from=builder /usr/local/lib/python3.11 /usr/local/lib/python3.11
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /usr/local/include /usr/local/include

# Copy application code
COPY . .

# Expose the application port
EXPOSE 8000

# Command to run the FastAPI app
CMD ["uvicorn", "app.main:app", "--host=0.0.0.0", "--port=8000"]
