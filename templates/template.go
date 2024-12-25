package templates

import "fmt"

type Template struct {
	Framework        string `json:"framework"`
	BaseImage        string `json:"base_image"`
	WorkDir          string `json:"work_dir"`
	RequirementsFile string `json:"requirements_file"`
	RunCommand       string `json:"run_command"`
}

type Templates struct {
	Templates map[string]Template `json:"templates"`
}

// Development Dockerfile templates for various frameworks
const FlaskDockerfileDevTemplate = `
# Dockerfile for Flask (Development)

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY {{.RequirementsFile}} {{.WorkDir}}

RUN pip install -r {{.RequirementsFile}}

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

// Mapping of development templates to their corresponding frameworks
var DockerfileDevTemplates = map[string]string{
	"flask":   FlaskDockerfileDevTemplate,
	"django":  DjangoDockerfileDevTemplate,
	"nodejs":  NodejsDockerfileDevTemplate,
	"react":   ReactDockerfileDevTemplate,
	"golang":  GolangDockerfileDevTemplate,
	"angular": AngularDockerfileDevTemplate,
	"fastapi": FastAPIDockerfileDevTemplate,
	"express": ExpressDockerfileDevTemplate,
}

func GetDockerfileTemplate(framework string) (string, error) {
	templ, exists := DockerfileDevTemplates[framework]
	if !exists {
		return "", fmt.Errorf("framework not found")
	}

	return templ, nil
}