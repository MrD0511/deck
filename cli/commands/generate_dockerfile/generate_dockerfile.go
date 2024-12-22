package generatedockerfile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/MrD0511/deck/internal/stack"
	"github.com/spf13/cobra"
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

			generate_dockerfile_procedure(dir)
		},
	}
}

func generate_dockerfile_procedure(dir string) {

	fmt.Println("Detecting the framework...")

	detected_frameworks_report, err := stack.DetectFramework(dir)
	
	if err != nil {
		fmt.Println(err)
	}

	stack.PrintTechStackReport(detected_frameworks_report)

	if len(detected_frameworks_report)>1 {
		fmt.Println("More than one framework detected in the given directory. Please select one(working directory).")

		var selectedDir string
		promptToSelectDir(detected_frameworks_report, &selectedDir)


	}

}	

//func to select a dir form given options
func promptToSelectDir(framework_report []stack.TechStackReport, selected *string) error{
	var options []string

	//fetching dir and framework from framework report
	for _, report := range framework_report {

		option := "option" 		//random string for var initialization
		if report.Directory == "." {
			option = fmt.Sprintf("%s (%s)", "current Directory", report.Framework)
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

	err := survey.AskOne(prompt, selected)
	if err != nil {
		return err
	}

	if *selected == "Add custom directory" {
		err := promptToAddCustomDir(selected)
		if err != nil {
			return err
		}
	}

	return nil
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
