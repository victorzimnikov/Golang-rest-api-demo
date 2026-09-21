package repository

type CommentsRepository struct {
	db DBTX
}

func NewCommentsRepository(db DBTX) *CommentsRepository {
	return &CommentsRepository{
		db: db,
	}
}
