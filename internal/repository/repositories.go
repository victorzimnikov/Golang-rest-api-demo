package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
)

type DBTX interface {
	Exec(
		ctx context.Context,
		sql string,
		arguments ...any,
	) (pgconn.CommandTag, error)
	QueryRow(
		ctx context.Context,
		sql string,
		arguments ...any,
	) pgx.Row
}

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
