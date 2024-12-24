package python_template

const Dockerfile_python_template = `
# Dockerfile for {{.Framework}}

FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY {{.RequirementsFile}} {{.WorkDir}}

RUN {{if eq .Framework "flask"}}pip install -r {{.RequirementsFile}}{{else if eq .Framework "nodejs"}}npm install{{end}}

CMD ["sh", "-c", "{{.RunCommand}}"]
`