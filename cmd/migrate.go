package cmd

import "github.com/spf13/cobra"

var startCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Start the migration",
	Long:  "Start the app HTTP server and also consuming ride events",
	Run:   start,
}

func start(_ *cobra.Command, _ []string) {}
