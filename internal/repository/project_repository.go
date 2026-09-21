package repository

import (
	"context"
	"fmt"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
)

type ProjectRepository struct {
	queries *db.Queries
}

func NewProjectRepository(queries *db.Queries) *ProjectRepository {
	return &ProjectRepository{
		queries: queries,
	}
}

func (r *ProjectRepository) SaveProject(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	row, err := r.queries.CreateProject(ctx, db.CreateProjectParams{
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	})
	if err != nil {
		return nil, fmt.Errorf("create project: %w", err)
	}

	return &domain.Project{
		ID:          domain.ProjectID(row.ID),
		Name:        row.Name,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}, nil
}
