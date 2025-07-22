package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/MichaelYoDev/goTasks/tasks"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Task ID must be a number")
			os.Exit(1)
		}

		file, err := tasks.OpenFileForReadWrite()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to open task file:", err)
			os.Exit(1)
		}
		defer tasks.CloseFile(file)

		taskList, err := tasks.LoadTasksFromFile(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to read tasks:", err)
			os.Exit(1)
		}

		updated := make([]tasks.Task, 0, len(taskList))
		found := false
		for _, t := range taskList {
			if t.ID == taskID {
				found = true
				continue // skip this task
			}
			updated = append(updated, t)
		}

		if !found {
			fmt.Fprintln(os.Stderr, "Task not found:", taskID)
			os.Exit(1)
		}

		err = tasks.SaveTasksToFile(file, updated)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to save tasks:", err)
			os.Exit(1)
		}

		fmt.Println("Deleted task", taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
