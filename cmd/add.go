package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/DePavelPo/task-manager-cli/internal/storage"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add new task",
	Args:  cobra.MinimumNArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		db := storage.OpenDB()
		defer storage.CloseDB(db)

		storage.Migrate(db)

		err := storage.InsertTask(args[0], db)
		if err != nil {
			logrus.Errorf("while InsertTask: %v", err)
			return
		}

		fmt.Println("Task was added:", args[0])
	},
}
