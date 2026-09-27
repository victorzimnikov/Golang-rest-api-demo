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
	Skip  int64
	Limit int32
	Q     string
}

func NewCommentsService(commentsRepository *repository.CommentsRepository) *CommentsService {
	return &CommentsService{
		commentsRepository: commentsRepository,
	}
}

func (s *CommentsService) CreateComment(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	return s.commentsRepository.CreateComment(ctx, comment)
}
