package service

import "github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"

type CommentsService struct {
	commentsRepository *repository.CommentsRepository
}

func NewCommentsService(commentsRepository *repository.CommentsRepository) *CommentsService {
	return &CommentsService{
		commentsRepository: commentsRepository,
	}
}
