package stack

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"unicode/utf16"
	"unicode/utf8"

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

func DetectFramework(dir string) ([]TechStackReport, error) {

	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	files := map[string]bool{"package.json": true, "requirements.txt": true, "go.mod": true}
	paths, err := searchFileInSubDirs(dir, files)

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
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	walkErr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && (info.Name() == "node_modules" || info.Name() == ".git" || info.Name() == ".angular") {
			return filepath.SkipDir
		}

		name := strings.ToLower(info.Name())
		if !info.IsDir() {
			if _, exists := targets[name]; exists {
				relativePath, err := filepath.Rel(cwd, filepath.Dir(path))
				if err != nil {
					return fmt.Errorf("failed to calculate relative path: %w", err)
				}
				result[relativePath] = info.Name()
			}
		}

		return nil
	})
	if walkErr != nil {
		return nil, walkErr
	}
	return result, nil
}

func utf16ToUtf8(data []byte) (string, error) {
	if len(data)%2 != 0 {
		return "", fmt.Errorf("invalid UTF-16LE data length")
	}

	// Decode UTF-16LE
	utf16Data := make([]uint16, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		utf16Data[i/2] = uint16(data[i]) | uint16(data[i+1])<<8
	}

	// Convert UTF-16 to UTF-8
	utf8Buf := new(bytes.Buffer)
	for _, r := range utf16.Decode(utf16Data) {
		if r == utf8.RuneError {
			return "", fmt.Errorf("invalid UTF-16 data")
		}
		utf8Buf.WriteRune(r)
	}

	return utf8Buf.String(), nil
}

// Detect Python Framework from requirements.txt
func detectPythonFramework(filePath string) Framework {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return Unknown
	}
	defer file.Close()

	// Read file as UTF-16LE
	reader := bufio.NewReader(file)
	rawData, err := reader.Peek(2) // Check for BOM
	if err != nil {
		fmt.Println("Error reading file header:", err)
		return Unknown
	}

	// Detect encoding
	var data []byte
	if rawData[0] == 0xFF && rawData[1] == 0xFE { // UTF-16LE BOM
		buf := new(bytes.Buffer)
		reader.Discard(2) // Skip BOM
		for {
			chunk := make([]byte, 1024)
			n, err := reader.Read(chunk)
			buf.Write(chunk[:n])
			if err != nil {
				break
			}
		}
		utf8Data, _ := utf16ToUtf8(buf.Bytes()) // Convert to UTF-8 string
		data = []byte(utf8Data)                 // Convert string to []byte
	} else {
		// Assume UTF-8 by default
		data, _ = io.ReadAll(reader) // Already []byte
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if strings.HasPrefix(trimmedLine, "#") || trimmedLine == "" {
			continue
		}

		lower := strings.ToLower(trimmedLine)
		switch {
		case strings.Contains(lower, "flask"):
			return Flask
		case strings.Contains(lower, "fastapi"):
			return FastAPI
		case strings.Contains(lower, "django"):
			return Django
		}
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

func PrintTechStackReport(reports []TechStackReport) {
	// Print header with color
	color.Cyan("\nTech Stack Detection Report")
	fmt.Printf("%-45s %-20s %-10s\n", "Directory", "File", "Framework")

	for _, report := range reports {
		// Color the columns differently
		DirectoryColor.Printf("%-45s", report.Directory)
		FileColor.Printf("%-20s", report.File)

		if report.Framework == Unknown {
			UnknownColor.Printf("%-10s\n", report.Framework)
		} else {
			FrameworkColor.Printf("%-10s\n", report.Framework)
		}
	}
	fmt.Println("")
}
