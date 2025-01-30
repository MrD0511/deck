package templates

import (
	"encoding/json"
	"fmt"
)

func Get_frameworks_template() (Templates, error) {
	var templates Templates
	
	jsonData, err := json.Marshal(Frameworks_template)
	if err != nil {
		return templates, err
	}
	
	err = json.Unmarshal([]byte(string(jsonData)), &templates)
	if err != nil {
		return templates, err
	}

	return templates, nil
} 

func GetDockerfileTemplate(framework string, isDev bool) (string, error) {

	var templ string
	var exists bool

	if isDev {
		templ, exists = DockerfileDevTemplates[framework]
		if !exists {
			return "", fmt.Errorf("framework not found")
		}
	}else {
		templ, exists = DockerfileProdTemplates[framework]
		if !exists {
			return "", fmt.Errorf("framework not found")
		}
	}

	return templ, nil
}