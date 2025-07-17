package cmd

import (
	"fmt"

	"github.com/DePavelPo/task-manager-cli/internal/storage"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete task by id",
	Args:  cobra.MaximumNArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		store, err := storage.NewSQLiteStore("./task-manager.db")
		if err != nil {
			logrus.Fatalf("init sqlite3 store error: %v", err)
		}
		defer store.CloseDB()

		argUint64, err := strToUint64(args[0])
		if err != nil {
			logrus.Errorf("cant use argument: %v", err)
			return
		}

		err = store.DeleteTask(argUint64)
		if err != nil {
			logrus.Errorf("while DeleteTask: %v", err)
			return
		}

		fmt.Printf("Task %d was deleted\n", argUint64)
	},
}
