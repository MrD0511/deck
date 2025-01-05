package createDockerfiles

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
)

func CreateDockerIgnore(ignored_files []string, dir string) error {
	
	dockerIgnoreContent := bytes.Buffer{}
	for _, entry := range ignored_files {
		dockerIgnoreContent.WriteString(entry+ "\n")
	}

	dockerIgnorePath := filepath.Join(dir, ".dockerignore")

	err := os.WriteFile(dockerIgnorePath, dockerIgnoreContent.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("can not write docker ignore: %v", err)
	}

	return nil

}

