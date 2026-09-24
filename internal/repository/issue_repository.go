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

type GetProjectIssuesListParams struct {
	Q          string
	Skip       int64
	LimitCount int32
	ProjectID  domain.ProjectID
	Status     domain.IssueStatus
	Priority   domain.IssuePriority
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

func (r *IssueRepository) GetProjectIssuesList(ctx context.Context, params GetProjectIssuesListParams) (int64, []domain.ProjectIssueListItem, error) {
	total, err := r.queries.CountProjectIssues(ctx, db.CountProjectIssuesParams{
		Q:         params.Q,
		ProjectID: int64(params.ProjectID),
		Status:    string(params.Status),
		Priority:  string(params.Priority),
	})
	if err != nil {
		return 0, nil, fmt.Errorf("count project issues: %w", err)
	}

	listRaw, err := r.queries.GetProjectIssuesList(ctx, db.GetProjectIssuesListParams{
		Q:          params.Q,
		ProjectID:  int64(params.ProjectID),
		Status:     string(params.Status),
		Priority:   string(params.Priority),
		Skip:       params.Skip,
		LimitCount: params.LimitCount,
	})
	if err != nil {
		return 0, nil, fmt.Errorf("get project issues list: %w", err)
	}

	list := make([]domain.ProjectIssueListItem, len(listRaw))

	for idx, item := range listRaw {
		list[idx] = domain.ProjectIssueListItem{
			ID:       domain.IssueID(item.ID),
			Title:    item.Title,
			Status:   domain.IssueStatus(item.Status),
			Priority: domain.IssuePriority(item.Priority),
		}
	}

	return total, list, nil
}

func (r *IssueRepository) DeleteIssue(ctx context.Context) error {
	return nil
}

func (r *IssueRepository) UpdateIssue(ctx context.Context) error {
	return nil
}
