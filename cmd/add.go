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
		store, err := storage.NewSQLiteStore("./task-manager.db")
		if err != nil {
			logrus.Fatalf("init sqlite3 store error: %v", err)
		}
		defer store.CloseDB()

		err = store.SaveTask(args[0])
		if err != nil {
			logrus.Errorf("while SaveTask: %v", err)
			return
		}

		fmt.Println("Task was added:", args[0])
	},
}
