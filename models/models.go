package models

import (
	"time"
)

type Task struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
