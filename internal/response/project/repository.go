package project

import "go-task/internal/response/task"

// Repository 暫時供尚未遷移的 Handler 使用。
// 新的 POST /projects 使用 Service → ProjectStore。
type Repository interface {
	List() []Project
	GetByID(id int64) (Project, bool)
	Update(id int64, name string) (Project, bool)
	Delete(id int64) error

	CreateTask(title string, projectID int64) (task.Task, error)
	GetProjectAllTasks(projectID int64) ([]task.Task, bool)
	GetProjectSingleTask(projectID, taskID int64) (task.Task, bool)
	UpdateTask(
		projectID, taskID int64,
		title, status string,
	) (task.Task, bool)
	DeleteTask(projectID, taskID int64) bool
}
