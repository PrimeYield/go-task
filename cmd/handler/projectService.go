package handler

import (
	"context"

	"go-task/internal/response/project"
)

type ProjectService interface {
	CreateProject(ctx context.Context, name string) (project.Project, error)
}
