package generate_dockerfile

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/MrD0511/deck/internal/createDockerfiles"
	"github.com/MrD0511/deck/internal/stack"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/MrD0511/deck/templates"
)



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

	
	var selected_option map[string]string
	
	if len(detected_frameworks_report)>1 {
		stack.PrintTechStackReport(detected_frameworks_report)
		
		fmt.Println("More than one framework detected in the given directory. Please select one(working directory).")

		selected_option, err = promptToSelectDir(detected_frameworks_report)
		if err != nil {
			return err
		}
	}else if len(detected_frameworks_report) == 1{
		selected_option = map[string]string{
			"Directory" : string(detected_frameworks_report[0].Directory),
			"Framework" : string(detected_frameworks_report[0].Framework),
		}
	}else {			
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
	
	fmt.Println(selected_option)
	template, exists := templates.Templates[strings.ToLower(selected_option["Framework"])]
	if !exists {
		return fmt.Errorf("framework '%s' not found", selected_option["Framework"])
	}

	template, err = showTemplateByName(template)
	if err != nil {
		return err
	}


	err = createDockerfiles.CreateDockerfileByTemplate(template, selected_option["Directory"])
	if err != nil {
		return err
	}
	
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
		Default: ".",
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

func loadTemplates() (templates.Templates, error) {

	var templates templates.Templates

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

func showTemplateByName(template templates.Template) (templates.Template, error) {

	// Define colors
	title := color.New(color.FgCyan, color.Bold).SprintFunc()
	key := color.New(color.FgBlue).SprintFunc()
	value := color.New(color.FgGreen).SprintFunc()

	// Display the template with color
	fmt.Println("")
	fmt.Printf("%s %s\n", title("Template for:"), value(template.Framework))
	fmt.Printf("%s: %s\n", key("Framework"), value(template.Framework))
	fmt.Printf("%s: %s\n", key("Base Image"), value(template.BaseImage))
	fmt.Printf("%s: %s\n", key("Work Directory"), value(template.WorkDir))
	fmt.Printf("%s: %s\n", key("Requirements File"), value(template.RequirementsFile))
	fmt.Printf("%s: %s\n", key("Run Command"), value(template.RunCommand))

	var customize bool
	survey.AskOne(&survey.Confirm{
		Message: "Do you want to customize this template?",
		Default: false,
	}, &customize)
	fmt.Println("")

	if customize {
		overwriteTemplateOutput(9)
		template = customize_template(template)
		overwriteTemplateOutput(15)
		showTemplateByName(template)
	}

	return template, nil
} 

func customize_template(template templates.Template) templates.Template {

	survey.AskOne(&survey.Input{
		Message: "Framework:",
		Default: template.Framework,
	}, &template.Framework)

	survey.AskOne(&survey.Input{
		Message: "Base Image:",
		Default: template.BaseImage,
	}, &template.BaseImage)

	survey.AskOne(&survey.Input{
		Message: "Working Dir:",
		Default: template.WorkDir,
	}, &template.WorkDir)

	survey.AskOne(&survey.Input{
		Message: "Requirements file name:",
		Default: template.RequirementsFile,
	}, &template.RequirementsFile)

	survey.AskOne(&survey.Input{
		Message: "Run command:",
		Default: template.RunCommand,
	}, &template.RunCommand)

	return template
}	

func overwriteTemplateOutput(lines int) {
	// Overwrite the template part with blank lines
	for i := 0; i < lines; i++ {
		fmt.Print("\033[1A") // Move up 3 lines
		fmt.Print("\033[2K") // Clear the current line
	}
}
