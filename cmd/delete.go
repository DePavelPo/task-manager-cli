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
		db := storage.OpenDB()
		defer storage.CloseDB(db)

		storage.Migrate(db)

		argUint64, err := strToUint64(args[0])
		if err != nil {
			logrus.Errorf("cant use argument: %v", err)
			return
		}

		err = storage.DeleteTask(argUint64, db)
		if err != nil {
			logrus.Errorf("while DeleteTask: %v", err)
			return
		}

		fmt.Printf("Task %d was deleted\n", argUint64)
	},
}
