package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MichaelYoDev/goTasks/tasks"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := strings.Join(args, " ")
		task := tasks.Task{
			Description: description,
			CreatedAt:   time.Now(),
			IsComplete:  false,
		}

		err := tasks.AddTask(task)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to add task:", err)
			os.Exit(1)
		}

		fmt.Println("Task added:", description)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
