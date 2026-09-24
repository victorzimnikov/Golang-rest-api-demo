package domain

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"
)

const MaxIssueTitleLength = 100

var (
	ErrIssueNotFound           = errors.New("issue not found")
	ErrIssueTitleRequired      = errors.New("issue title is required")
	ErrIssueTitleTooLong       = errors.New("issue title is too long")
	ErrIssueTitleAlreadyExists = errors.New("issue title already exists")
	ErrInvalidIssueStatus      = errors.New("invalid issue status")
	ErrInvalidIssuePriority    = errors.New("invalid issue priority")
)

type IssueID int64

type IssueStatus string

const (
	IssueOpenStatus       IssueStatus = "open"
	IssueInProgressStatus IssueStatus = "in_progress"
	IssueDoneStatus       IssueStatus = "done"
)

type IssuePriority string

const (
	IssueLowPriority    IssuePriority = "low"
	IssueMediumPriority IssuePriority = "medium"
	IssueHighPriority   IssuePriority = "high"
)

type Issue struct {
	ID          IssueID
	Title       string
	Description string
	Status      IssueStatus
	Priority    IssuePriority
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Project     ProjectShort
}

func NewIssue(
	projectID ProjectID,
	title string,
	description string,
	pStatus *IssueStatus,
	pPriority *IssuePriority,
	dueDate *time.Time,
) (*Issue, error) {
	status := IssueOpenStatus
	priority := IssueMediumPriority

	normalizedTitle, err := NormalizeIssueTitle(title)
	if err != nil {
		return nil, err
	}

	if pStatus != nil {
		if !ValidateIssueStatus(*pStatus) {
			return nil, ErrInvalidIssueStatus
		}

		status = *pStatus
	}

	if pPriority != nil {
		if !ValidateIssuePriority(*pPriority) {
			return nil, ErrInvalidIssuePriority
		}

		priority = *pPriority
	}

	return &Issue{
		Project: ProjectShort{
			ID: projectID,
		},
		Title:       normalizedTitle,
		Description: description,
		Status:      status,
		Priority:    priority,
		DueDate:     dueDate,
	}, nil
}

func NormalizeIssueTitle(title string) (string, error) {
	normalizedTitle := strings.TrimSpace(title)

	if err := validateIssueTitle(normalizedTitle); err != nil {
		return "", err
	}

	return normalizedTitle, nil
}

func ValidateIssueStatus(status IssueStatus) bool {
	return status == IssueOpenStatus || status == IssueInProgressStatus || status == IssueDoneStatus
}

func ValidateIssuePriority(priority IssuePriority) bool {
	return priority == IssueLowPriority || priority == IssueMediumPriority || priority == IssueHighPriority
}

func validateIssueTitle(name string) error {
	if name == "" {
		return ErrIssueTitleRequired
	}

	if utf8.RuneCountInString(name) > MaxIssueTitleLength {
		return ErrIssueTitleTooLong
	}

	return nil
}
