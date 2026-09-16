package project

import (
	"go-task/internal/task"
	"time"
)

type Project struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	TaskCount int64     `json:"task_count"`
	tasks     []task.Task
	taskID    int64
}

type CreateProjectRequest struct {
	Name string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UploadProjectRequest struct {
	Name string `json:"name"`
}
