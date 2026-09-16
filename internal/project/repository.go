package project

import (
	"go-task/internal/task"
	"sort"
	"sync"
	"time"
)

type Repository interface {
	// Project
	Create(name string) Project
	List() []Project
	// FindByID(ID int64) bool
	GetByID(ID int64) (Project, bool)
	Update(ID int64, reName string) (Project, bool)
	Delete(ID int64) error

	// Task
	CreateTask(title string, projectID int64) (task.Task, bool)
	GetProjectAllTasks(projectID int64) ([]task.Task, bool)
	GetProjectSingleTask(projectID, taskID int64) (task.Task, bool)
	UpdateTask(projectID, taskID int64, title string, status string) (task.Task, bool)
	DeleteTask(projectID, taskID int64) bool

	// internal helpers
	// findProjectLocked(projectID int64) (Project, bool)
	// findTaskIndexLocked(Project Project, taskID int64) (task.Task, int, bool)
}

type MemoryRepository struct {
	mu       sync.RWMutex
	projects map[int64]Project
	nextID   int64
}

var _ Repository = (*MemoryRepository)(nil)

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		projects: make(map[int64]Project),
		nextID:   1,
	}
}

func (r *MemoryRepository) Create(name string) Project {
	r.mu.Lock()
	defer r.mu.Unlock()

	project := Project{
		ID:        r.nextID,
		Name:      name,
		CreatedAt: time.Now(),
		TaskCount: 0,
		tasks:     []task.Task{},
		taskID:    1,
	}

	r.projects[project.ID] = project
	r.nextID++

	return project
}

func (r *MemoryRepository) List() []Project {
	r.mu.RLock()
	defer r.mu.RUnlock()

	projects := make([]Project, 0, len(r.projects))

	for _, project := range r.projects {
		projects = append(projects, project)
	}
	// for _, project := range r.projects {
	// 	projects[project.ID-1] = project
	// }
	sort.Slice(projects, func(i int, j int) bool {
		return projects[i].ID < projects[j].ID
	})

	return projects
}

// Bad design in concurrency.
// func (r *MemoryRepository) FindByID(ID int64) bool {
// 	r.mu.RLock()
// 	defer r.mu.RUnlock()

// 	if _, exist := r.projects[ID]; !exist {
// 		return false
// 	}

// 	return true
// }

// Get project by ID
func (r *MemoryRepository) GetByID(ID int64) (Project, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// if _, exist := r.projects[ID]; !exist {
	// 	return Project{}, false
	// }
	// return r.projects[ID], true
	project, exist := r.projects[ID]
	return project, exist
}

func (r *MemoryRepository) Update(ID int64, reName string) (Project, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	project, exist := r.projects[ID]
	if !exist {
		return project, exist
	}
	project.Name = reName
	r.projects[ID] = project
	return project, true
}

func (r *MemoryRepository) Delete(ID int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exist := r.projects[ID]
	if !exist {
		return ErrProjectNotFound
	}
	if len(r.projects[ID].tasks) > 0 {
		return ErrProjectHasTasks
	}
	delete(r.projects, ID)
	return nil
}
