package domain

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	ErrCommentTextRequired = errors.New("comment text is required")
	ErrCommentNotFound     = errors.New("comment not found")
	ErrCommentTextTooLong  = errors.New("comment text is too long")
)

const MaxCommentTextLength = 500

type CommentID int64

type Comment struct {
	ID        CommentID
	Issue     IssueShort
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewComment(issueID IssueID, text string) (*Comment, error) {
	text, err := NormalizeCommentText(text)
	if err != nil {
		return nil, err
	}

	return &Comment{
		Issue: IssueShort{ID: issueID},
		Text:  text,
	}, nil
}

func (c *Comment) ChangeText(text string, now time.Time) error {
	text, err := NormalizeCommentText(text)
	if err != nil {
		return err
	}

	c.Text = text
	c.UpdatedAt = now

	return nil
}

func NormalizeCommentText(text string) (string, error) {
	text = strings.TrimSpace(text)

	if err := validateCommentText(text); err != nil {
		return "", err
	}

	return text, nil
}

func validateCommentText(text string) error {
	if text == "" {
		return ErrCommentTextRequired
	}

	if utf8.RuneCountInString(text) > MaxCommentTextLength {
		return ErrCommentTextTooLong
	}

	return nil
}
