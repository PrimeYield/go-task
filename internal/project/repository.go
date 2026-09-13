package project

import (
	"sort"
	"sync"
	"time"
)

type Repository interface {
	Create(name string) Project
	List() []Project
	// FindByID(ID int64) bool
	GetByID(ID int64) (Project, bool)
	Update(ID int64, Name string) (Project, bool)
	Delete(ID int64) bool
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
		ID:       r.nextID,
		Name:     name,
		CreateAt: time.Now(),
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

func (r *MemoryRepository) Update(ID int64, Name string) (Project, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	project, exist := r.projects[ID]
	if !exist {
		return project, exist
	}
	project.Name = Name
	r.projects[ID] = project
	return project, true
}

func (r *MemoryRepository) Delete(ID int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exist := r.projects[ID]
	if !exist {
		return false
	}
	delete(r.projects, ID)
	return true
}
