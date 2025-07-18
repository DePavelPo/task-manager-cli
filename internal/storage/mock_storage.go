package storage

import (
	"github.com/DePavelPo/task-manager-cli/models"

	"github.com/stretchr/testify/mock"
)

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) LoadTasks(completed *bool) ([]models.Task, error) {
	args := m.Called(completed)
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockStorage) SaveTask(title string) error {
	args := m.Called(title)
	return args.Error(0)
}

func (m *MockStorage) UpdateTask(id uint64, completed bool) error {
	args := m.Called(id, completed)
	return args.Error(0)
}

func (m *MockStorage) DeleteTask(id uint64) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockStorage) CloseDB() error {
	args := m.Called()
	return args.Error(0)
}
