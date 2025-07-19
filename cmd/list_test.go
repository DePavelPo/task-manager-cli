package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/DePavelPo/task-manager-cli/internal/storage"
	"github.com/DePavelPo/task-manager-cli/models"
	"github.com/stretchr/testify/mock"
)

func TestList(t *testing.T) {

	cases := map[string]struct {
		storeMockResponse []models.Task
		storeMockError    error
		completed         *bool
		expectedTasks     []models.Task
	}{
		"two tasks": {
			storeMockResponse: []models.Task{
				{ID: 1, Title: "Test Task 1", Completed: false},
				{ID: 2, Title: "Test Task 2", Completed: true},
			},
			storeMockError: nil,
			completed:      nil,
			expectedTasks: []models.Task{
				{Title: "Test Task 1", Completed: false},
				{Title: "Test Task 2", Completed: true},
			},
		},
		"one completed task": {
			storeMockResponse: []models.Task{
				{ID: 1, Title: "Test Task", Completed: true},
			},
			storeMockError: nil,
			completed:      boolPtr(true),
			expectedTasks: []models.Task{
				{Title: "Test Task", Completed: true},
			},
		},
		"one pending task": {
			storeMockResponse: []models.Task{
				{ID: 1, Title: "Test Task 2", Completed: false},
			},
			storeMockError: nil,
			completed:      boolPtr(false),
			expectedTasks: []models.Task{
				{Title: "Test Task 2", Completed: false},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mockStore := new(storage.MockStorage)
			storeFactory = func() (storage.Storage, error) { return mockStore, nil }
			defer func() {
				storeFactory = func() (storage.Storage, error) {
					store, err := storage.NewSQLiteStore("./task-manager.db")
					return store, err
				}
			}()

			mockStore.On("LoadTasks", mock.Anything).Return(tc.storeMockResponse, tc.storeMockError)
			mockStore.On("CloseDB").Return(nil)

			cmd := listCmd
			cmd.Flags().Set("completed", "false")
			cmd.Flags().Set("pending", "false")

			if tc.completed != nil {
				if *tc.completed {
					cmd.Flags().Set("completed", "true")
				} else {
					cmd.Flags().Set("pending", "true")
				}
			}

			r, w, _ := os.Pipe()
			origStdout := os.Stdout
			os.Stdout = w

			// Call the Run function directly
			cmd.Run(cmd, []string{""})

			w.Close()
			os.Stdout = origStdout

			var outputBuf bytes.Buffer
			_, _ = outputBuf.ReadFrom(r)
			output := outputBuf.String()

			// Assert
			mockStore.AssertExpectations(t)

			// Get the actual calls and verify the arguments
			calls := mockStore.Calls
			if len(calls) > 0 {
				actualCompleted := calls[0].Arguments.Get(0).(*bool)
				if tc.completed == nil {
					if actualCompleted != nil {
						t.Errorf("expected nil, got %v", *actualCompleted)
					}
				} else {
					if actualCompleted == nil {
						t.Errorf("expected %v, got nil", *tc.completed)
					} else if *actualCompleted != *tc.completed {
						t.Errorf("expected %v, got %v", *tc.completed, *actualCompleted)
					}
				}
			}
			// check if the output has expected values
			for _, expectedTask := range tc.expectedTasks {
				if !bytes.Contains([]byte(output), []byte(expectedTask.Title)) {
					t.Errorf("expected task: %s, got: %s", expectedTask.Title, output)
				}
			}
		})
	}
}
