package templates

const FlaskDockerfileDevTemplate = `
# Dockerfile for Flask (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

# Install dependencies
COPY {{.RequirementsFile}} /app/
RUN pip install --no-cache-dir -r {{.RequirementsFile}}

# Copy application code
COPY . /app/

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`

const DjangoDockerfileDevTemplate = `
# Dockerfile for Django (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY {{.RequirementsFile}} {{.WorkDir}}

RUN pip install -r {{.RequirementsFile}}

COPY . .

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}}","{{.Port}}"]
`

const NodejsDockerfileDevTemplate = `
# Dockerfile for Node.js (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

RUN npm install

COPY . .

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`

const ReactDockerfileDevTemplate = `
# Dockerfile for React (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

RUN npm install

COPY . .

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`

const GolangDockerfileDevTemplate = `
# Dockerfile for Golang (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY go.mod go.sum {{.WorkDir}}

RUN go mod download

COPY . .

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`

const AngularDockerfileDevTemplate = `
# Dockerfile for Angular (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

RUN npm install

COPY . .

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port {{.Port}}"]
`

const FastAPIDockerfileDevTemplate = `
# Dockerfile for FastAPI (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY {{.RequirementsFile}} {{.WorkDir}}

RUN pip install -r {{.RequirementsFile}}

COPY . .

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`

const ExpressDockerfileDevTemplate = `
# Dockerfile for Express.js (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

RUN npm install

COPY . .

EXPOSE {{.Port}}

CMD ["sh", "-c", "{{.RunCommand}} --port={{.Port}}"]
`
