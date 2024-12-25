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

CMD ["sh", "-c", "{{.RunCommand}}"]
`

const DjangoDockerfileDevTemplate = `
# Dockerfile for Django (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY {{.RequirementsFile}} {{.WorkDir}}

RUN pip install -r {{.RequirementsFile}}

COPY . .

CMD ["sh", "-c", "{{.RunCommand}}"]
`

const NodejsDockerfileDevTemplate = `
# Dockerfile for Node.js (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

RUN npm install

COPY . .

CMD ["sh", "-c", "{{.RunCommand}}"]
`

const ReactDockerfileDevTemplate = `
# Dockerfile for React (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

RUN npm install

COPY . .

CMD ["sh", "-c", "{{.RunCommand}}"]
`

const GolangDockerfileDevTemplate = `
# Dockerfile for Golang (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY go.mod go.sum {{.WorkDir}}

RUN go mod download

COPY . .

CMD ["sh", "-c", "{{.RunCommand}}"]
`


const AngularDockerfileDevTemplate = `
# Dockerfile for Angular (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

RUN npm install

COPY . .

CMD ["sh", "-c", "{{.RunCommand}}"]`


const FastAPIDockerfileDevTemplate = `
# Dockerfile for FastAPI (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY {{.RequirementsFile}} {{.WorkDir}}

RUN pip install -r {{.RequirementsFile}}

COPY . .

CMD ["sh", "-c", "{{.RunCommand}}"]
`

const ExpressDockerfileDevTemplate = `
# Dockerfile for Express.js (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY package.json package-lock.json {{.WorkDir}}

RUN npm install

COPY . .

CMD ["sh", "-c", "{{.RunCommand}}"]
`

