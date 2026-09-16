package project

import (
	"go-task/internal/task"
	"time"
)

func (r *MemoryRepository) CreateTask(title string, projectID int64) (task.Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	project, exist := r.projects[projectID]
	if !exist {
		return task.Task{}, false
	}

	var newTask task.Task
	newTask.CreatedAt = time.Now()
	newTask.Title = title
	newTask.Status = "todo"
	newTask.TaskID = project.taskID
	project.taskID++
	project.tasks = append(project.tasks, newTask)
	project.TaskCount = int64(len(project.tasks))

	r.projects[projectID] = project
	return newTask, true
}

func (r *MemoryRepository) GetProjectAllTasks(projectID int64) ([]task.Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	project, exist := r.projects[projectID]
	if !exist {
		return nil, false
	}
	return project.tasks, true
}

func (r *MemoryRepository) GetProjectSingleTask(ProjectID, taskID int64) (task.Task, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	// resProject, exist := r.projects[ProjectID]
	// if !exist {
	// 	return task.Task{}, false
	// }
	// for i := 0; i < len(resProject.tasks); i++ {
	// 	if resProject.tasks[i].TaskID == taskID {
	// 		return resProject.tasks[i], true
	// 	}
	// }
	resProject, exist := r.findProjectLocked(ProjectID)
	if !exist {
		return task.Task{}, false
	}
	resTask, _, exist := r.findTaskIndexLocked(resProject, taskID)
	if !exist {
		return task.Task{}, false
	}
	return resTask, true
}

func (r *MemoryRepository) UpdateTask(projectID, taskID int64, title string, status string) (task.Task, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	resProject, exist := r.findProjectLocked(projectID)
	if !exist {
		return task.Task{}, false
	}
	resTask, idx, exist := r.findTaskIndexLocked(resProject, taskID)
	if !exist {
		return task.Task{}, false
	}
	resTask.Title = title
	resTask.Status = status
	r.projects[projectID].tasks[idx] = resTask
	return resTask, true
}

func (r *MemoryRepository) DeleteTask(projectID, taskID int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	resProject, exist := r.findProjectLocked(projectID)
	if !exist {
		return false
	}
	_, idx, exist := r.findTaskIndexLocked(resProject, taskID)
	if !exist {
		return false
	}
	tempTasks := resProject.tasks[idx+1:]
	resProject.tasks = resProject.tasks[:idx]
	resProject.tasks = append(resProject.tasks, tempTasks...)
	resProject.TaskCount = int64(len(resProject.tasks))

	r.projects[projectID] = resProject
	return true
}
