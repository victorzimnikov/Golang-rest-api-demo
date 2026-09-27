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
	IssueID    domain.IssueID
}

type UpdateCommentParams struct {
	CommentID domain.CommentID
	Text      *string
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

func (r *CommentsRepository) GetIssueCommentsList(ctx context.Context, params GetIssueCommentsListParams) (int64, []domain.IssueCommentListItem, error) {
	total, err := r.queries.CountIssueComments(ctx, db.CountIssueCommentsParams{
		IssueID: int64(params.IssueID),
		Q:       params.Q,
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, nil, domain.ErrIssueNotFound
	}

	if err != nil {
		return 0, nil, fmt.Errorf("count issue comments: %w", err)
	}

	listRaw, err := r.queries.GetIssueCommentsList(ctx, db.GetIssueCommentsListParams{
		IssueID:    int64(params.IssueID),
		Q:          params.Q,
		Skip:       params.Skip,
		LimitCount: params.LimitCount,
	})
	if err != nil {
		return 0, nil, fmt.Errorf("get issue comments list: %w", err)
	}

	list := make([]domain.IssueCommentListItem, len(listRaw))

	for idx, item := range listRaw {
		list[idx] = domain.IssueCommentListItem{
			ID:   domain.CommentID(item.ID),
			Text: item.Text,
		}
	}

	return total, list, nil
}

func (r *CommentsRepository) GetCommentByID(ctx context.Context, id domain.CommentID) (*domain.Comment, error) {
	row, err := r.queries.GetComment(ctx, int64(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrCommentNotFound
	}

	if err != nil {
		return nil, err
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

func (r *CommentsRepository) DeleteComment(ctx context.Context, id domain.CommentID) error {
	_, err := r.queries.DeleteComment(ctx, int64(id))
	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrCommentNotFound
	}

	if err != nil {
		return fmt.Errorf("delete comment by ID:%d: %w", id, err)
	}

	return nil
}

func (r *CommentsRepository) UpdateComment(ctx context.Context, params UpdateCommentParams) (*domain.Comment, error) {
	response, err := r.queries.UpdateComment(ctx, db.UpdateCommentParams{
		Text: toNullableText(params.Text),
		ID:   int64(params.CommentID),
	})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrCommentNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("update comment by ID:%d: %w", params.CommentID, err)
	}

	return &domain.Comment{
		ID:        domain.CommentID(response.ID),
		Text:      response.Text,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
		Issue: domain.IssueShort{
			ID:    domain.IssueID(response.IssueID),
			Title: response.IssueTitle,
		},
	}, nil
}
