package storage

import (
	"testing"

	"github.com/DePavelPo/task-manager-cli/models"

	"github.com/pkg/errors"
)

func TestLoadTasks(t *testing.T) {

	cases := map[string]struct {
		mockResponse  []models.Task
		mockError     error
		completed     *bool
		expectedTasks []models.Task
		expectedError error
	}{
		"no tasks": {
			mockResponse:  []models.Task{},
			mockError:     nil,
			completed:     nil,
			expectedTasks: []models.Task{},
			expectedError: nil,
		},
		"completed tasks": {
			mockResponse: []models.Task{
				{ID: 1, Title: "Test", Completed: true},
			},
			mockError: nil,
			completed: boolPtr(true),
			expectedTasks: []models.Task{
				{ID: 1, Title: "Test", Completed: true},
			},
			expectedError: nil,
		},
		"pending tasks": {
			mockResponse: []models.Task{
				{ID: 1, Title: "Test", Completed: false},
			},
			mockError: nil,
			completed: boolPtr(false),
			expectedTasks: []models.Task{
				{ID: 1, Title: "Test", Completed: false},
			},
			expectedError: nil,
		},
		"error": {
			mockResponse:  []models.Task{},
			mockError:     errors.New("test error"),
			completed:     nil,
			expectedTasks: []models.Task{},
			expectedError: errors.New("test error"),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mockStore := new(MockStorage)
			mockStore.On("LoadTasks", tc.completed).Return(tc.mockResponse, tc.mockError)

			tasks, err := mockStore.LoadTasks(tc.completed)

			// Assertions
			mockStore.AssertExpectations(t)
			if err != nil {
				if tc.expectedError != nil {
					if err.Error() != tc.expectedError.Error() {
						t.Errorf("expected error: %v, got: %v", tc.expectedError, err)
						return
					}
				} else {
					t.Errorf("unexpected error: %v", err)
					return
				}
			}
			if len(tasks) != len(tc.expectedTasks) {
				t.Errorf("expected tasks: %+v, got: %+v", tc.expectedTasks, tasks)
				return
			}
			for i, task := range tasks {
				if task != tc.expectedTasks[i] {
					t.Errorf("expected task: %+v, got: %+v", tc.expectedTasks[i], task)
				}
			}
		})
	}
}
