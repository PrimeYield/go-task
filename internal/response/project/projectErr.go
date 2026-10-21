package project

import "errors"

var (
	ErrProjectNotFound    = errors.New("project not found")
	ErrProjectHasTasks    = errors.New("project has tasks")
	ErrInvalidProjectName = errors.New("project name cannot be empty")
	ErrInvalidProjectID   = errors.New("project ID must be positive")
)
