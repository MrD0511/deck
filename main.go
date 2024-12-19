package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/MrD0511/deck/internal/stack"
)

func main() {
	var directory string
	if len(os.Args) > 1 {
		directory = os.Args[1]
	} else {
		// If no argument is provided, use the current directory
		directory = "."
	}

	// Get the absolute path of the directory
	absDir, err := filepath.Abs(directory)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Search for the framework in the given directory or its parent directories
	framework, err := stack.DetectFramework(absDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(framework) == 0 {
		fmt.Println("No framework detected in the provided directory or its parent directories.")
	} else {
		fmt.Println("Detected framework:", framework)
		stack.PrintTechStackReport(framework)
	}

}
