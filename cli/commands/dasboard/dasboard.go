package dasboard_cli

import (
	"log"

	"github.com/MrD0511/deck/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	// Add "dashboard" as a subcommand of root
	DashboardCmd.AddCommand(startCmd)
}

var DashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Manage the dashboard",
	Long:  "Commands to manage and interact with the custom dashboard server.",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the dashboard server",
	Long:  "Start the dashboard server to serve Kubernetes resources on the web.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting the dashboard server...")
		if err := server.StartServer(); err != nil {
			log.Fatalf("Error starting the server: %v", err)
		}
	},
}