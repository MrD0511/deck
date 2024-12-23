package generate_dockerfile

import (
	"encoding/json"
	"fmt"

	// "html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/MrD0511/deck/internal/stack"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type Template struct {
	Framework  			string `json:"framework"`
	BaseImage			string `json:"base_image"`
	WorkDir				string `json:"work_dir"`
	RequirementsFile	string `json:"requirements_file"`
	RunCommand			string `json:"run_command"`	
}

type Templates struct {
	Templates map[string]Template `json:"templates"`
}

func GeneateCommand() *cobra.Command{
	return &cobra.Command{
		Use: "generate",
		Short: "Generate a dockerfile",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string){

			dir := args[0]
			cwd, err := os.Getwd()

			if err != nil {
				return
			}

			if dir == "." {
				dir = cwd
			}else{
				dir = filepath.Join(cwd, dir)
			}

			info, err := os.Stat(dir)
			if os.IsNotExist(err) {
				fmt.Printf("Directory does not exist: %s\n", dir)
				return
			}

			if !info.IsDir() {
				fmt.Printf("The provided path is not a directory: %s\n", dir)
				return
			}

			err = generate_dockerfile_procedure(dir)
			if err != nil {
				fmt.Println(err)
				return
			}
		},
	}
}

func generate_dockerfile_procedure(dir string) error{

	fmt.Println("Detecting the framework...")

	detected_frameworks_report, err := stack.DetectFramework(dir)
	
	if err != nil {
		fmt.Println(err)
	}

	stack.PrintTechStackReport(detected_frameworks_report)

	var selected_option map[string]string

	if len(detected_frameworks_report)>1 {
		fmt.Println("More than one framework detected in the given directory. Please select one(working directory).")

		selected_option, err = promptToSelectDir(detected_frameworks_report)
		if err != nil {
			return err
		}
	}else{
		fmt.Println("No frameworks found")
		selected_option, err = addCustomeDirNFramework()
		if err != nil {
			return err
		}
	}

	if string(selected_option["Framework"]) == "unknown" {
		var selected_framework string
		promptToSelectFramework(&selected_framework)
		selected_option["Framework"] = strings.ToLower(selected_framework)
	}

	templates, err := loadTemplates()
	if err != nil {
		return err
	}

	showTemplateByName(templates, selected_option["Framework"])

	return nil
}	

//func to select a dir form given options
func promptToSelectDir(framework_report []stack.TechStackReport) (map[string]string, error){
	var options []string

	//fetching dir and framework from framework report
	for _, report := range framework_report {

		option := "option" 		//random string for var initialization
		if report.Directory == "." {
			option = fmt.Sprintf("%s (%s)", "Current directory", report.Framework)
		}else{
			option = fmt.Sprintf("%s (%s)", report.Directory, report.Framework)
		}

		options = append(options, option)
	}


	options = append(options, "Add custom directory")

	//survey pkg used to select a dir
	prompt := &survey.Select{
		Message: "Select a working directory",
		Options: options,
		Default: options[0],
	}

	var selected_option string
	err := survey.AskOne(prompt, &selected_option)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if selected_option == "Add custom directory" {
		return addCustomeDirNFramework()
	}	


	selected_option= strings.Replace(selected_option, "Current directory", ".", 1)

    for _, report := range framework_report {
		option := fmt.Sprintf("%s (%s)", report.Directory, report.Framework)
        if selected_option == option {
            return map[string]string{
                "Directory": report.Directory,
                "Framework": string(report.Framework),
            }, nil
        }
    }

    return nil, fmt.Errorf("invalid selection")
}

func promptToAddCustomDir(selected *string) error{
	prompt := &survey.Input{
		Message: "Add custom working directory:",
	}

	err := survey.AskOne(prompt, selected)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	selectedDir := filepath.Join(cwd, *selected)

	info, err := os.Stat(selectedDir)
	if os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n",selectedDir)
		return err
	}

	if !info.IsDir() {
		fmt.Printf("The provided path is not a directory: %s\n", selectedDir)
		return err
	}

	return nil
}

func promptToSelectFramework(selected *string) error{
	options := [6]string{"React", "Angular", "Flask", "FastAPI", "Django", "Gin"}

	prompt := &survey.Select{
		Message: "Select a Framework",
		Options: options[:],
		Default: options[0],
	}

	err := survey.AskOne(prompt, selected)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func addCustomeDirNFramework() (map[string]string, error) {
	var custom_dir string
	err := promptToAddCustomDir(&custom_dir)
	if err != nil {
		return nil, err
	}

	var custom_framework string
	err = promptToSelectFramework(&custom_framework)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"Directory" : custom_dir,
		"Framework" : custom_framework,
	}, nil
}

func loadTemplates() (Templates, error) {

	var templates Templates

	file, err := ioutil.ReadFile("./templates/template.json")
	if err != nil {
		return templates, fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(file, &templates)
	if err != nil {
		return templates, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return templates, nil
}

func showTemplateByName(templates Templates, frameworkName string) error {
	template, exists := templates.Templates[frameworkName]
	if !exists {
		return fmt.Errorf("framework '%s' not found", frameworkName)
	}

	// Define colors
	title := color.New(color.FgCyan, color.Bold).SprintFunc()
	key := color.New(color.FgBlue).SprintFunc()
	value := color.New(color.FgGreen).SprintFunc()

	// Display the template with color
	fmt.Println("")
	fmt.Printf("%s %s\n", title("Template for:"), value(frameworkName))
	fmt.Printf("%s: %s\n", key("Framework"), value(template.Framework))
	fmt.Printf("%s: %s\n", key("Base Image"), value(template.BaseImage))
	fmt.Printf("%s: %s\n", key("Work Directory"), value(template.WorkDir))
	fmt.Printf("%s: %s\n", key("Requirements File"), value(template.RequirementsFile))
	fmt.Printf("%s: %s\n", key("Run Command"), value(template.RunCommand))
	fmt.Println("")

	return nil
} 
