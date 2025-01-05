package templates

const FlaskDockerfileProdTemplate = `
# Dockerfile for Flask (Production)

FROM {{.BaseImage}} as builder

WORKDIR {{.WorkDir}}

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies in a virtual environment
RUN python -m venv /venv
ENV PATH="/venv/bin:$PATH"

COPY {{.RequirementsFile}} ./
RUN pip install --no-cache-dir -r {{.RequirementsFile}}

# Copy application code
COPY . .

# Final production image
FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

# Set environment variables
ENV PATH="/venv/bin:$PATH"

# Copy virtual environment and application from the builder stage
COPY --from=builder /venv /venv
COPY --from=builder {{.WorkDir}} {{.WorkDir}}

EXPOSE {{.Port}}

CMD ["sh", "-c", "gunicorn app:app --bind 0.0.0.0:{{.Port}}"]
`

const DjangoDockerfileProdTemplate = `
# Dockerfile for Django (Production)

FROM {{.BaseImage}} as builder

WORKDIR {{.WorkDir}}

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies in a virtual environment
RUN python -m venv /venv
ENV PATH="/venv/bin:$PATH"

COPY {{.RequirementsFile}} ./
RUN pip install --no-cache-dir -r {{.RequirementsFile}}

COPY . .

# Final production image
FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

# Set environment variables
ENV PATH="/venv/bin:$PATH"

# Copy virtual environment and application from the builder stage
COPY --from=builder /venv /venv
COPY --from=builder {{.WorkDir}} {{.WorkDir}}

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} 0.0.0.0:{{.Port}}"]
`

const NodejsDockerfileProdTemplate = `
# Dockerfile for Node.js (Production)

FROM {{.BaseImage}} as builder

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}
RUN npm install --production

COPY . .

# Final production image
FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY --from=builder {{.WorkDir}} {{.WorkDir}}

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`

const ReactDockerfileProdTemplate = `
# Dockerfile for React (Production)

FROM {{.BaseImage}} as builder

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}
RUN npm install --production

COPY . .
RUN npm run build

# Final production image
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY --from=builder /usr/src/app/build .

EXPOSE {{.Port}}

CMD ["nginx", "-g", "daemon off;"]
`

const GolangDockerfileProdTemplate = `
# Dockerfile for Golang (Production)

FROM {{.BaseImage}} as builder

WORKDIR {{.WorkDir}}

COPY go.mod go.sum {{.WorkDir}}
RUN go mod download

COPY . .
RUN go build -o app

# Final production image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder {{.WorkDir}}/app .

EXPOSE {{.Port}}

CMD ["./app", "--port", "{{.Port}}"]
`

const AngularDockerfileProdTemplate = `
# Dockerfile for Angular (Production)

FROM {{.BaseImage}} as builder

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

# Install the Angular CLI globally and production dependencies
RUN npm install -g @angular/cli && npm install 

COPY . .
RUN npm run build --prod -- --output-path=dist/build

# Final production image
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY --from=builder /usr/src/app/dist/build/browser .

EXPOSE {{.Port}}

CMD ["nginx", "-g", "daemon off;"]
`

const FastAPIDockerfileProdTemplate = `
# Dockerfile for FastAPI (Production)

FROM {{.BaseImage}} as builder

WORKDIR {{.WorkDir}}

RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Install Python dependencies in a virtual environment
RUN python -m venv /venv
ENV PATH="/venv/bin:$PATH"

COPY {{.RequirementsFile}} ./
RUN pip install --no-cache-dir -r {{.RequirementsFile}}

COPY . .

# Final production image
FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

# Set environment variables
ENV PATH="/venv/bin:$PATH"

# Copy virtual environment and application from the builder stage
COPY --from=builder /venv /venv
COPY --from=builder {{.WorkDir}} {{.WorkDir}}

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`

const ExpressDockerfileProdTemplate = `
# Dockerfile for Express.js (Production)

FROM {{.BaseImage}} as builder

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}
RUN npm install --production

COPY . .

# Final production image
FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY --from=builder {{.WorkDir}} {{.WorkDir}}

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`

// Mapping of production templates to their corresponding frameworks
var DockerfileProdTemplates = map[string]string{
	"flask":   FlaskDockerfileProdTemplate,
	"django":  DjangoDockerfileProdTemplate,
	"nodejs":  NodejsDockerfileProdTemplate,
	"react":   ReactDockerfileProdTemplate,
	"golang":  GolangDockerfileProdTemplate,
	"angular": AngularDockerfileProdTemplate,
	"fastapi": FastAPIDockerfileProdTemplate,
	"express": ExpressDockerfileProdTemplate,
}
