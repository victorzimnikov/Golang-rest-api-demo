package domain

import (
	"errors"
	"strings"
	"time"
)

var ErrCommentTextRequired = errors.New("comment text is required")

type CommentID int64

type Comment struct {
	ID        CommentID
	IssueID   IssueID
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewComment(issueID IssueID, text string, now time.Time) (*Comment, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, ErrCommentTextRequired
	}

	return &Comment{
		IssueID:   issueID,
		Text:      text,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (c *Comment) ChangeText(text string, now time.Time) error {
	text = strings.TrimSpace(text)
	if text == "" {
		return ErrCommentTextRequired
	}

	c.Text = text
	c.UpdatedAt = now

	return nil
}
