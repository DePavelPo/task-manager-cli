package cmd

import (
	"fmt"

	"github.com/DePavelPo/task-manager-cli/internal/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of all tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var completed *bool
		needCompleted, err := cmd.Flags().GetBool("completed")
		if err != nil {
			logrus.Errorf("cant get --completed flag: %v", err)
			return
		}
		if needCompleted {
			completed = pointer(true)
		} else {
			needPending, err := cmd.Flags().GetBool("pending")
			if err != nil {
				logrus.Errorf("cant get --pending flag: %v", err)
				return
			}
			if needPending {
				completed = pointer(false)
			}
		}

		db := storage.OpenDB()
		defer storage.CloseDB(db)

		storage.Migrate(db)

		tasks, err := storage.SelectTasks(completed, db)
		if err != nil {
			logrus.Errorf("while SelectTasks: %v", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("Task list is empty")
			return
		}

		for _, task := range tasks {
			status := " "
			if task.Completed {
				status = "✓"
			}
			fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Title)
		}
	},
}

func pointer[T any](val T) *T {
	return &val
}
