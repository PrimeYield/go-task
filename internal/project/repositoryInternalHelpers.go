package project

import (
	"go-task/internal/task"
	"log"
)

// pre-check path validate
func (r *MemoryRepository) findProjectLocked(projectID int64) (Project, bool) {
	resProject, exist := r.projects[projectID]
	if !exist {
		log.Printf("project %d is not exist.", projectID)
	}
	return resProject, exist
}

func (r *MemoryRepository) findTaskIndexLocked(project Project, taskID int64) (task.Task, int, bool) {
	for idx, task := range project.tasks {
		if task.TaskID == taskID {
			return task, idx, true
		}
	}
	log.Printf("task %d is not exist.", taskID)
	return task.Task{}, -1, false
}
