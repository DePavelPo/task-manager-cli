package cmd

import (
	"fmt"
	"os"

	"github.com/DePavelPo/task-manager-cli/internal/storage"
	"github.com/DePavelPo/task-manager-cli/models"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/pkg/errors"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of all tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		completed, err := checkCompletedAndPendingFlags(cmd)
		if err != nil {
			logrus.Error(err)
		}

		store, err := storage.NewSQLiteStore("./task-manager.db")
		if err != nil {
			logrus.Fatalf("init sqlite3 store error: %v", err)
		}
		defer store.CloseDB()

		tasks, err := store.LoadTasks(completed)
		if err != nil {
			logrus.Errorf("while LoadTasks: %v", err)
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
	var completed *bool
	needCompleted, err := cmd.Flags().GetBool("completed")
	if err != nil {
		return completed, errors.Errorf("cant get --completed flag: %v", err)
	}
	if needCompleted {
		completed = pointer(true)
	} else {
		needPending, err := cmd.Flags().GetBool("pending")
		if err != nil {
			return completed, errors.Errorf("cant get --pending flag: %v", err)
		}
		if needPending {
			completed = pointer(false)
		}
	}

	return completed, nil
}

func responseOutput(tasks []models.Task) {
	table := tablewriter.NewTable(os.Stdout)
	table.Header("ID", "Title", "Status")

	greenColor := color.New(color.FgHiGreen).SprintFunc()
	tableData := [][]any{}
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

func pointer[T any](val T) *T {
	return &val
}
