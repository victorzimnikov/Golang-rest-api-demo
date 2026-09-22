package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
)

const (
	pgUniqueViolationCode                      = "23505"
	pgProjectNameUniqueViolationConstraintName = "project_name_unique_idx"
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

	var pgError *pgconn.PgError

	if errors.As(err, &pgError) && pgError.Code == pgUniqueViolationCode && pgError.ConstraintName == pgProjectNameUniqueViolationConstraintName {
		return nil, domain.ErrProjectNameAlreadyExists
	}

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

func (r *ProjectRepository) GetProjectByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	project, err := r.queries.GetProject(ctx, int64(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrProjectNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("get project by ID:%d: %w", id, err)
	}

	return &domain.Project{
		ID:          domain.ProjectID(project.ID),
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}, nil
}

func (r *ProjectRepository) GetProjectsList(ctx context.Context, params db.GetProjectsListParams) (int64, []domain.Project, error) {
	total, err := r.queries.CountProjects(ctx, params.Q)
	if err != nil {
		return 0, nil, fmt.Errorf("count projects: %w", err)
	}

	listRaw, err := r.queries.GetProjectsList(ctx, params)
	if err != nil {
		return 0, nil, fmt.Errorf("get projects list: %w", err)
	}

	list := make([]domain.Project, len(listRaw))

	for idx, item := range listRaw {
		list[idx] = domain.Project{
			ID:          domain.ProjectID(item.ID),
			Name:        item.Name,
			Description: item.Description,
		}
	}

	return total, list, nil
}

func (r *ProjectRepository) DeleteProject(ctx context.Context, id domain.ProjectID) error {
	_, err := r.queries.DeleteProject(ctx, int64(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrProjectNotFound
	}

	if err != nil {
		return fmt.Errorf("delete project by ID:%d: %w", id, err)
	}

	return nil
}
