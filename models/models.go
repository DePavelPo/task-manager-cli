package models

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

type Task struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Task) Scan(value interface{}) error {
	return jsoniter.Unmarshal([]byte(value.(string)), t)
}
