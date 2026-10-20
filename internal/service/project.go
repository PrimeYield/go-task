package service

import (
	"context"
	"strings"

	"go-task/internal/response/project"
)

// ProjectStore 定義 Service 需要的資料存取能力。
type ProjectStore interface {
	Create(ctx context.Context, name string) (project.Project, error)
}

// ProjectService 是真正執行業務規則的物件。
type ProjectService struct {
	store ProjectStore
}

// NewProjectService 建立 Service，並交給它一個 Store。
func NewProjectService(store ProjectStore) *ProjectService {
	return &ProjectService{
		store: store,
	}
}

// CreateProject 實作建立專案的業務流程。
func (s *ProjectService) CreateProject(
	ctx context.Context,
	name string,
) (project.Project, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return project.Project{}, project.ErrInvalidProjectName
	}

	return s.store.Create(ctx, name)
}
