package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/MichaelYoDev/goTasks/tasks"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var showAll bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all incomplete tasks (or all tasks with --all)",
	Run: func(cmd *cobra.Command, args []string) {
		file, err := tasks.OpenFileForRead()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to open task file:", err)
			os.Exit(1)
		}
		defer tasks.CloseFile(file)

		allTasks, err := tasks.LoadTasksFromFile(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to load tasks:", err)
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		if showAll {
			fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
		} else {
			fmt.Fprintln(w, "ID\tTask\tCreated")
		}

		for _, t := range allTasks {
			if !showAll && t.IsComplete {
				continue
			}
			createdAgo := timediff.TimeDiff(t.CreatedAt)
			if showAll {
				fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", t.ID, t.Description, createdAgo, t.IsComplete)
			} else {
				fmt.Fprintf(w, "%d\t%s\t%s\n", t.ID, t.Description, createdAgo)
			}
		}

		w.Flush()
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show all tasks including completed ones")
	rootCmd.AddCommand(listCmd)
}
