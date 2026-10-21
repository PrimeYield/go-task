package handler

import (
	"context"

	"go-task/internal/response/project"
)

type ProjectService interface {
	CreateProject(ctx context.Context, name string) (project.Project, error)
	GetProject(ctx context.Context, id int64) (project.Project, error)
}
