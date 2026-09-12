package project

import "sync"

type Repository interface {
	Create(name string) Project
	List() []Project
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
		ID:   r.nextID,
		Name: name,
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

	return projects
}
