package repository

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
)

const (
	pgUniqueViolationCode                      = "23505"
	pgIssueTitleUniqueViolationConstraintName  = "issue_title_unique_idx"
	pgProjectNameUniqueViolationConstraintName = "project_name_unique_idx"
)

type Repositories struct {
	CommentsRepository *CommentsRepository
	ProjectRepository  *ProjectRepository
	IssueRepository    *IssueRepository
}

func NewRepositories(queries *db.Queries) *Repositories {
	commentsRepository := NewCommentsRepository(queries)
	projectRepository := NewProjectRepository(queries)
	issueRepository := NewIssueRepository(queries)

	return &Repositories{
		CommentsRepository: commentsRepository,
		ProjectRepository:  projectRepository,
		IssueRepository:    issueRepository,
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

func checkIsNotUniqueName(err error, constraintName string) bool {
	var pgError *pgconn.PgError

	return errors.As(err, &pgError) && pgError.Code == pgUniqueViolationCode && pgError.ConstraintName == constraintName
}

func fromNullableDate(value pgtype.Date) *time.Time {
	if !value.Valid {
		return nil
	}

	return &value.Time
}

func toNullableDate(value *time.Time) pgtype.Date {
	if value == nil {
		return pgtype.Date{}
	}

	return pgtype.Date{
		Time:  *value,
		Valid: true,
	}
}
