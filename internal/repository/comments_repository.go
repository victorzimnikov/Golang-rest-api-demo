package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
)

type CommentsRepository struct {
	queries *db.Queries
}

type GetIssueCommentsListParams struct {
	Skip       int64
	LimitCount int32
	Q          string
}

func NewCommentsRepository(queries *db.Queries) *CommentsRepository {
	return &CommentsRepository{
		queries: queries,
	}
}

func (r *CommentsRepository) CreateComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	row, err := r.queries.CreateComment(ctx, db.CreateCommentParams{
		IssueID: int64(comment.Issue.ID),
		Text:    comment.Text,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrIssueNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("create comment: %w", err)
	}

	return &domain.Comment{
		ID:        domain.CommentID(row.ID),
		Text:      row.Text,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		Issue: domain.IssueShort{
			ID:    domain.IssueID(row.IssueID),
			Title: row.IssueTitle,
		},
	}, nil
}
