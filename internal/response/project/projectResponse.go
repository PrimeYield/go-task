package project

import (
	"time"
)

type ResponseProject struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	TaskCount int64     `json:"task_count"`
	CreatedAt time.Time `json:"created_at"`
}

// type MemoryRepository struct {
// 	mu       sync.RWMutex
// 	projects map[int64]project.Project
// 	nextID   int64
// }

// var _ ProjectRepository = (*MemoryRepository)(nil)

// func NewMemoryRepository() *MemoryRepository {
// 	return &MemoryRepository{
// 		projects: make(map[int64]project.Project),
// 		nextID:   1,
// 	}
// }

// func (r *MemoryRepository) Create(name string) project.Project {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	project := project.Project{
// 		ID:        r.nextID,
// 		Name:      name,
// 		CreatedAt: time.Now(),
// 		TaskCount: 0,
// 		tasks:     []task.Task{},
// 		taskID:    1,
// 	}

// 	r.projects[project.ID] = project
// 	r.nextID++

// 	return project
// }

// func (r *MemoryRepository) List() []project.Project {
// 	r.mu.RLock()
// 	defer r.mu.RUnlock()

// 	projects := make([]project.Project, 0, len(r.projects))

// 	for _, project := range r.projects {
// 		projects = append(projects, project)
// 	}
// 	// for _, project := range r.projects {
// 	// 	projects[project.ID-1] = project
// 	// }
// 	sort.Slice(projects, func(i int, j int) bool {
// 		return projects[i].ID < projects[j].ID
// 	})

// 	return projects
// }

// // Bad design in concurrency.
// // func (r *MemoryRepository) FindByID(ID int64) bool {
// // 	r.mu.RLock()
// // 	defer r.mu.RUnlock()

// // 	if _, exist := r.projects[ID]; !exist {
// // 		return false
// // 	}

// // 	return true
// // }

// // Get project by ID
// func (r *MemoryRepository) GetByID(ID int64) (project.Project, bool) {
// 	r.mu.RLock()
// 	defer r.mu.RUnlock()

// 	// if _, exist := r.projects[ID]; !exist {
// 	// 	return Project{}, false
// 	// }
// 	// return r.projects[ID], true
// 	project, exist := r.projects[ID]
// 	return project, exist
// }

// func (r *MemoryRepository) Update(ID int64, reName string) (project.Project, bool) {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	project, exist := r.projects[ID]
// 	if !exist {
// 		return project, exist
// 	}
// 	project.Name = reName
// 	r.projects[ID] = project
// 	return project, true
// }

// func (r *MemoryRepository) Delete(ID int64) error {
// 	r.mu.Lock()
// 	defer r.mu.Unlock()

// 	_, exist := r.projects[ID]
// 	if !exist {
// 		return project.ErrProjectNotFound
// 	}
// 	if len(r.projects[ID].tasks) > 0 {
// 		return project.ErrProjectHasTasks
// 	}
// 	delete(r.projects, ID)
// 	return nil
// }
