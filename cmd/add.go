package cmd

import (
	"fmt"

	"github.com/DePavelPo/task-manager-cli/internal/storage"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
			logrus.WithError(err).Fatal("init sqlite3 store error")
		}
		defer store.CloseDB()

		err = store.SaveTask(args[0])
		if err != nil {
			logrus.WithError(err).Error("while SaveTask")
			return
		}

		fmt.Println("Task was added:", args[0])
	},
}
