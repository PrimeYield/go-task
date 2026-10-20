package task

import (
	"time"
)

type ResponseTask struct {
	TaskID    int64     `json:"task_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
