package cli

import (
	// "fmt"

	dasboard_cli "github.com/MrD0511/deck/deck-cli/cli/commands/dasboard"
	"github.com/MrD0511/deck/deck-cli/cli/commands/generate_dockerfile"
	"github.com/spf13/cobra"
)

func Cli_main(){
	rootCmd := &cobra.Command{
		Use: "deck",
		Short: "A tool to simplify docker container generation and handling.",
	}

	rootCmd.AddCommand(generate_dockerfile.GenerateCommand())
	rootCmd.AddCommand(dasboard_cli.DashboardCmd)
	rootCmd.Execute()
}

