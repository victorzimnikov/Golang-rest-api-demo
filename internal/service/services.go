package service

import "github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"

type Services struct {
	CommentsService *CommentsService
	ProjectService  *ProjectService
}

func NewServices(repositories *repository.Repositories) *Services {
	commentsService := NewCommentsService(repositories.CommentsRepository)
	projectService := NewProjectService(repositories.ProjectRepository)

	return &Services{
		ProjectService:  projectService,
		CommentsService: commentsService,
	}
}
