package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/MichaelYoDev/goTasks/tasks"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task ID]",
	Short: "Mark a task as complete",
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

		found := false
		for i := range taskList {
			if taskList[i].ID == taskID {
				taskList[i].IsComplete = true
				found = true
				break
			}
		}

		if !found {
			fmt.Fprintln(os.Stderr, "Task not found:", taskID)
			os.Exit(1)
		}

		err = tasks.SaveTasksToFile(file, taskList)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to save tasks:", err)
			os.Exit(1)
		}

		fmt.Println("Marked task", taskID, "as complete")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
