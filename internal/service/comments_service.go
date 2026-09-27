package service

import (
	"context"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"
)

type CommentsService struct {
	commentsRepository *repository.CommentsRepository
}

type GetIssueCommentsListCommand struct {
	Skip    int64
	Limit   int32
	Q       string
	IssueID domain.IssueID
}

type UpdateCommentCommand struct {
	CommentID domain.CommentID
	Text      *string
}

func NewCommentsService(commentsRepository *repository.CommentsRepository) *CommentsService {
	return &CommentsService{
		commentsRepository: commentsRepository,
	}
}

func (s *CommentsService) GetComment(ctx context.Context, id domain.CommentID) (*domain.Comment, error) {
	return s.commentsRepository.GetCommentByID(ctx, id)
}

func (s *CommentsService) CreateComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	return s.commentsRepository.CreateComment(ctx, comment)
}

func (s *CommentsService) GetIssueCommentsList(ctx context.Context, command GetIssueCommentsListCommand) (int64, []domain.IssueCommentListItem, error) {
	return s.commentsRepository.GetIssueCommentsList(ctx, repository.GetIssueCommentsListParams{
		Q:          command.Q,
		Skip:       command.Skip,
		LimitCount: command.Limit,
		IssueID:    command.IssueID,
	})
}

func (s *CommentsService) DeleteComment(ctx context.Context, id domain.CommentID) error {
	return s.commentsRepository.DeleteComment(ctx, id)
}

func (s *CommentsService) UpdateComment(ctx context.Context, command UpdateCommentCommand) (*domain.Comment, error) {
	return s.commentsRepository.UpdateComment(ctx, repository.UpdateCommentParams(command))
}
