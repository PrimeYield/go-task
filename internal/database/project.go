package database

import (
	"context"
	"database/sql"
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
