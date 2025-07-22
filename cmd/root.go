package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Tasks",
	Short: "A simple CLI todo app",
	Long:  `A command-line application to manage tasks using a CSV file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'goTasks add' to add a task")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
