package project

import "sync"

type Repository interface {
	Create(name string) Project
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
