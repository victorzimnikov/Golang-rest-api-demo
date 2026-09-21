package service

import "github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"

type Services struct {
	CommentsService *CommentsService
}

func NewServices(repositories *repository.Repositories) *Services {
	commentsService := NewCommentsService(repositories.CommentsRepository)

	return &Services{
		CommentsService: commentsService,
	}
}
