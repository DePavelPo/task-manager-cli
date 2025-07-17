package cmd

import (
	"fmt"
	"os"

	"github.com/DePavelPo/task-manager-cli/internal/storage"
	"github.com/DePavelPo/task-manager-cli/models"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of all tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		completed, err := checkCompletedAndPendingFlags(cmd)
		if err != nil {
			logrus.WithError(err).Error("failed to check flags")
			return
		}

		store, err := storage.NewSQLiteStore("./task-manager.db")
		if err != nil {
			logrus.WithError(err).Fatal("init sqlite3 store error")
		}
		defer store.CloseDB()

		tasks, err := store.LoadTasks(completed)
		if err != nil {
			logrus.WithError(err).Error("while LoadTasks")
			return
		}

		if len(tasks) == 0 {
			fmt.Println("Task list is empty")
			return
		}

		responseOutput(tasks)
	},
}

func checkCompletedAndPendingFlags(cmd *cobra.Command) (*bool, error) {
	needCompleted, err := cmd.Flags().GetBool("completed")
	if err != nil {
		return nil, errors.Errorf("cant get --completed flag: %v", err)
	}
	if needCompleted {
		val := true
		return &val, nil
	}

	needPending, err := cmd.Flags().GetBool("pending")
	if err != nil {
		return nil, errors.Errorf("cant get --pending flag: %v", err)
	}
	if needPending {
		val := false
		return &val, nil
	}

	return nil, nil
}

func responseOutput(tasks []models.Task) {
	table := tablewriter.NewTable(os.Stdout)
	table.Header("ID", "Title", "Status")

	greenColor := color.New(color.FgHiGreen).SprintFunc()
	tableData := make([][]any, 0, len(tasks))
	for _, task := range tasks {
		if task.Completed {
			tableData = append(tableData, []any{greenColor(task.ID), greenColor(task.Title), greenColor("✓")})
			continue
		}
		tableData = append(tableData, []any{fmt.Sprint(task.ID), task.Title, " "})
	}
	table.Bulk(tableData)
	table.Render()
}
