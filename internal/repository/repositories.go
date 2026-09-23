package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
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

func toNullableText(value *string) pgtype.Text {
	if value == nil {
		return pgtype.Text{}
	}

	return pgtype.Text{
		String: *value,
		Valid:  true,
	}
}
