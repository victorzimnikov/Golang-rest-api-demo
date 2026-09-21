package repository

import "github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"

type CommentsRepository struct {
	queries *db.Queries
}

func NewCommentsRepository(queries *db.Queries) *CommentsRepository {
	return &CommentsRepository{
		queries: queries,
	}
}
