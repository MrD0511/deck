package generatedockerfile

import (
	"fmt"
	"os"

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

			if dir == "." {
				cwd, err := os.Getwd()

				if err != nil {
					return
				}

				dir = cwd
			}
			generate_dockerfile_procedure(dir)
			fmt.Println("Detecting framework..")
		},
	}
}

func generate_dockerfile_procedure(dir string) {
	detected_frameworks_report, err := stack.DetectFramework(dir)
	
	if err != nil {
		fmt.Println(err)
	}

	stack.PrintTechStackReport(detected_frameworks_report)

}	