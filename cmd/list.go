package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task/taskdb"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := taskdb.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete.")
		} else {
			fmt.Println("You have the following tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i + 1, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}