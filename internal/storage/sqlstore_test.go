package storage

import (
	"testing"

	"github.com/DePavelPo/task-manager-cli/models"
)

func TestLoadTasks(t *testing.T) {
	mockStore := new(MockStorage)

	// Set up expected calls and return values
	mockStore.On("LoadTasks", (*bool)(nil)).Return([]models.Task{
		{ID: 1, Title: "Test", Completed: false},
	}, nil)

	// Use mockStore in your code under test
	tasks, err := mockStore.LoadTasks(nil)

	// Assertions
	mockStore.AssertExpectations(t)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tasks) != 1 || tasks[0].Title != "Test" {
		t.Fatalf("unexpected tasks: %+v", tasks)
	}
}
