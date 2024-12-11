package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fatih/color"
)

type Framework string

const (
	React   Framework = "react"
	Angular Framework = "angular"
	Django  Framework = "django"
	FastAPI Framework = "fastapi"
	Flask   Framework = "flask"
	Express Framework = "express"
	Gin     Framework = "gin"
	Fiber   Framework = "fiber"
	Unknown Framework = "unknown"
)

var (
	DirectoryColor = color.New(color.FgBlue)
	FileColor      = color.New(color.FgYellow)
	FrameworkColor = color.New(color.FgGreen)
	UnknownColor   = color.New(color.FgRed)
)

type TechStackReport struct {
	Directory string
	File      string
	Framework Framework
}

func detectFramework() ([]TechStackReport, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	files := map[string]bool{"package.json": true, "requirements.txt": true, "go.mod": true}
	paths, err := searchFileInSubDirs(cwd, files)

	if err != nil {
		return nil, err
	}

	var report []TechStackReport

	var wg sync.WaitGroup

	for path, file := range paths {
		wg.Add(1)

		go func(path, file string) {
			defer wg.Done()

			file = strings.ToLower(file)
			filePath := filepath.Join(path, file)
			var framework Framework

			switch {
			case file == "requirements.txt":
				framework = detectPythonFramework(filepath.Join(cwd, filePath))
			case file == "go.mod":
				framework = detectGoFramework(filepath.Join(cwd, filePath))
			case file == "package.json":
				framework = detectNodeFramework(filepath.Join(cwd, filePath))
			}

			report = append(report, TechStackReport{
				Directory: path,
				File:      file,
				Framework: framework,
			})

		}(path, file)

	}

	wg.Wait()

	return report, nil
}

func searchFileInSubDirs(root string, targets map[string]bool) (map[string]string, error) {
	result := make(map[string]string)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && (info.Name() == "node_modules" || info.Name() == ".git" || info.Name() == ".angular") {
			return filepath.SkipDir
		}

		name := strings.ToLower(info.Name())
		if !info.IsDir() {
			if _, exists := targets[name]; exists {
				relativePath, err := filepath.Rel(root, filepath.Dir(path))
				if err != nil {
					return fmt.Errorf("failed to calculate relative path: %w", err)
				}
				result[relativePath] = info.Name()
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func detectPythonFramework(filePath string) Framework {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Unknown
	}

	lines := strings.Split(string(data), "\n")
	fmt.Println("Reading requirements.txt:", filePath) // Debugging line
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "#") || trimmedLine == "" || trimmedLine == "��" {
			continue
		}

		// Print each line being checked for debugging
		fmt.Print("Checking line:", trimmedLine)

		lower := strings.ToLower(trimmedLine)
		switch {
		case strings.Contains(lower, "flask"):
			fmt.Println("Found")
			return Flask
		case strings.Contains(lower, "fastapi"):
			fmt.Println("Found")
			return FastAPI
		case strings.Contains(lower, "django"):
			fmt.Println("Found")
			return Django
		}
		fmt.Println("not")

	}

	return Unknown
}

func detectGoFramework(filePath string) Framework {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Unknown
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		lower := strings.ToLower(strings.TrimSpace(line))

		if strings.HasPrefix(lower, "//") {
			continue
		}

		switch {
		case strings.Contains(lower, "github.com/gin-gonic/gin"):
			return Gin
		case strings.Contains(lower, "github.com/gofiber/fiber"):
			return Fiber
		}
	}
	return Unknown
}

func detectNodeFramework(filePath string) Framework {
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		return Unknown
	}

	var pkg map[string]interface{}

	if err := json.Unmarshal(data, &pkg); err != nil {
		return Unknown
	}

	dependencies := getDependencies(pkg, "dependencies", "devDependencies")

	if _, exists := dependencies["react"]; exists {
		return React
	} else if _, exists := dependencies["@angular/core"]; exists {
		return Angular
	} else if _, exists := dependencies["express"]; exists {
		return Express
	}

	return Unknown
}

func getDependencies(pkg map[string]interface{}, keys ...string) map[string]bool {
	dependencies := make(map[string]bool)

	for _, key := range keys {
		if deps, ok := pkg[key].(map[string]interface{}); ok {
			for name := range deps {
				dependencies[name] = true
			}
		}
	}
	return dependencies
}

func printTechStackReport(reports []TechStackReport) {
	// Print header with color
	color.Cyan("\nTech Stack Detection Report")
	fmt.Printf("%-35s %-20s %-10s\n", "Directory", "File", "Framework")

	for _, report := range reports {
		// Color the columns differently
		DirectoryColor.Printf("%-35s", report.Directory)
		FileColor.Printf("%-20s", report.File)

		if report.Framework == Unknown {
			UnknownColor.Printf("%-10s\n", report.Framework)
		} else {
			FrameworkColor.Printf("%-10s\n", report.Framework)
		}
	}
}

func main() {
	reports, err := detectFramework()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the report
	printTechStackReport(reports)
}
