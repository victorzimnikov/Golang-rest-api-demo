package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
)

type CreateIssueParams struct {
	ProjectID   domain.ProjectID
	Title       string
	Description string
	Status      domain.IssueStatus
	Priority    domain.IssuePriority
	DueDate     *time.Time
}

type IssueRepository struct {
	queries *db.Queries
}

func NewIssueRepository(queries *db.Queries) *IssueRepository {
	return &IssueRepository{
		queries: queries,
	}
}

func (r *IssueRepository) CreateIssue(ctx context.Context, params CreateIssueParams) (*domain.Issue, error) {
	row, err := r.queries.CreateIssue(ctx, db.CreateIssueParams{
		ProjectID:   int64(params.ProjectID),
		Title:       params.Title,
		Description: params.Description,
		Status:      string(params.Status),
		Priority:    string(params.Priority),
		DueDate:     toNullableDate(params.DueDate),
	})

	if checkIsNotUniqueName(err, pgIssueTitleUniqueViolationConstraintName) {
		return nil, domain.ErrIssueTitleAlreadyExists
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrProjectNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("create issue: %w", err)
	}

	return &domain.Issue{
		ID:          domain.IssueID(row.ID),
		Title:       row.Title,
		Description: row.Description,
		Status:      domain.IssueStatus(row.Status),
		Priority:    domain.IssuePriority(row.Priority),
		DueDate:     fromNullableDate(row.DueDate),
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		Project: domain.ProjectShort{
			ID:   domain.ProjectID(row.ProjectID),
			Name: row.ProjectName,
		},
	}, nil
}

func (r *IssueRepository) GetIssue(ctx context.Context) error {
	return nil
}

func (r *IssueRepository) GetIssuesList(ctx context.Context) error {
	return nil
}

func (r *IssueRepository) DeleteIssue(ctx context.Context) error {
	return nil
}

func (r *IssueRepository) UpdateIssue(ctx context.Context) error {
	return nil
}
