package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
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

func NewRepositories(db DBTX) *Repositories {
	commentsRepository := NewCommentsRepository(db)
	projectRepository := NewProjectRepository(db)

	return &Repositories{
		CommentsRepository: commentsRepository,
		ProjectRepository:  projectRepository,
	}
}
