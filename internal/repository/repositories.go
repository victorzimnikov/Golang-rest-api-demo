package repository

import (
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
)

type Repositories struct {
	CommentsRepository *CommentsRepository
	ProjectRepository  *ProjectRepository
}

func NewRepositories(queries *db.Queries) *Repositories {
	commentsRepository := NewCommentsRepository(queries)
	projectRepository := NewProjectRepository(queries)

	return &Repositories{
		CommentsRepository: commentsRepository,
		ProjectRepository:  projectRepository,
	}
}
