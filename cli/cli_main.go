package cli

import (
	// "fmt"

	"github.com/spf13/cobra"
	"github.com/MrD0511/deck/cli/commands/generate_dockerfile"
)

func Cli_main(){
	rootCmd := &cobra.Command{
		Use: "deck",
		Short: "A tool to simplify docker container generation and handling.",
	}

	rootCmd.AddCommand(generate_dockerfile.GenerateCommand())
	rootCmd.Execute()
}

