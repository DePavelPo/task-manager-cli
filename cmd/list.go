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
		db := storage.OpenDB()
		defer storage.CloseDB(db)

		storage.Migrate(db)

		tasks, err := storage.SelectTasks(db)
		if err != nil {
			logrus.Errorf("while SelectTasks: %v", err)
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
