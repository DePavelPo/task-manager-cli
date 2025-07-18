package cmd

import (
	"fmt"
	"strconv"

	"github.com/DePavelPo/task-manager-cli/internal/storage"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark task as done",
	Args:  cobra.MinimumNArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		store, err := storage.NewSQLiteStore("./task-manager.db")
		if err != nil {
			logrus.WithError(err).Fatal("init sqlite3 store error")
		}
		defer store.CloseDB()

		argUint64, err := strToUint64(args[0])
		if err != nil {
			logrus.WithError(err).Error("cant use argument")
			return
		}

		err = store.UpdateTask(argUint64, true)
		if err != nil {
			logrus.WithError(err).Error("while UpdateTask")
			return
		}

		fmt.Printf("Task %d was marked as done\n", argUint64)
	},
}

func strToUint64(s string) (uint64, error) {
	uint64Value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64Value, nil
}
