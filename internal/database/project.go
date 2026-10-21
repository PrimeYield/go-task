package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"go-task/internal/response/project"
)

type ProjectSQL struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

type ProjectDB struct {
	db *sql.DB
}

func NewProjectDB(db *sql.DB) *ProjectDB {
	return &ProjectDB{
		db: db,
	}
}

func (r *ProjectDB) Create(
	ctx context.Context,
	name string,
) (project.Project, error) {
	const query = `
		INSERT INTO projects (name)
		VALUES ($1)
		RETURNING id, name, created_at
	`

	var row ProjectSQL

	err := r.db.QueryRowContext(ctx, query, name).Scan(
		&row.ID,
		&row.Name,
		&row.CreatedAt,
	)
	if err != nil {
		return project.Project{}, fmt.Errorf(
			"insert project: %w", err,
		)
	}

	return project.Project{
		ID:        row.ID,
		Name:      row.Name,
		TaskCount: 0,
		CreatedAt: row.CreatedAt,
	}, nil
}

func (r *ProjectDB) GetByID(
	ctx context.Context,
	id int64,
) (project.Project, error) {
	const query = `
		SELECT id, name, created_at
		FROM public.projects
		WHERE id = $1
	`

	var row ProjectSQL

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&row.ID,
		&row.Name,
		&row.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return project.Project{}, project.ErrProjectNotFound
	}

	if err != nil {
		return project.Project{}, fmt.Errorf(
			"get project by ID: %w", err,
		)
	}

	return project.Project{
		ID:        row.ID,
		Name:      row.Name,
		TaskCount: 0,
		CreatedAt: row.CreatedAt,
	}, nil
}
