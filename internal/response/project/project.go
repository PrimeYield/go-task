package project

import (
	"go-task/internal/response/task"
	"time"
)

type Project struct {
	ID        int64               `json:"id"`
	Name      string              `json:"name"`
	TaskCount int64               `json:"task_count"`
	CreatedAt time.Time           `json:"created_at"`
	tasks     map[int64]task.Task // maximum = 100
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

const LimitTaskCount = 100
