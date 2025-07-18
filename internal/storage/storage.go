package storage

import (
	"github.com/DePavelPo/task-manager-cli/models"
	_ "github.com/mattn/go-sqlite3"
)

type Storage interface {
	LoadTasks(completed *bool) ([]models.Task, error)
	SaveTask(title string) error
	UpdateTask(id uint64, completed bool) error
	DeleteTask(id uint64) error
}

func intToBool(i int) bool {
	return i != 0
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func boolPtr(b bool) *bool {
	return &b
}
