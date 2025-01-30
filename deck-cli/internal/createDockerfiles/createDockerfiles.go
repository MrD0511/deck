package createDockerfiles

import (
	// "fmt"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/MrD0511/deck/templates"
)

func CreateDockerfileByTemplate (dockerfile_values_template templates.Template, dir string, isDev bool) error {

	dockerfile_template, err := templates.GetDockerfileTemplate(dockerfile_values_template.Framework, isDev)
	if err != nil {
		return err
	}

	tmpl, err := template.New("Dockerfile").Parse(dockerfile_template)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	var renderedDockerfile bytes.Buffer
	err = tmpl.Execute(&renderedDockerfile, dockerfile_values_template)
	if err != nil {
		return fmt.Errorf("failed to render Dockerfile: %v", err)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}
	
	// Write the rendered Dockerfile to the target directory
	dockerfilePath := filepath.Join(dir, "/Dockerfile")
	err = os.WriteFile(dockerfilePath, renderedDockerfile.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write Dockerfile: %v", err)
	}

	fmt.Printf("Dockerfile successfully created at: %s\n", dockerfilePath)
	return nil
}

