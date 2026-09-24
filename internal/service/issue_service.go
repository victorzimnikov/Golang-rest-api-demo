package service

import (
	"context"
	"time"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"
)

type CreateIssueCommand struct {
	ProjectID   domain.ProjectID
	Title       string
	Description string
	Status      domain.IssueStatus
	Priority    domain.IssuePriority
	DueDate     *time.Time
}

type IssueService struct {
	issueRepository *repository.IssueRepository
}

func NewIssueService(issueRepository *repository.IssueRepository) *IssueService {
	return &IssueService{
		issueRepository: issueRepository,
	}
}

func (s *IssueService) CreateIssue(ctx context.Context, issue *domain.Issue) (*domain.Issue, error) {
	params := repository.CreateIssueParams{
		ProjectID:   issue.Project.ID,
		Title:       issue.Title,
		Description: issue.Description,
		Status:      issue.Status,
		Priority:    issue.Priority,
		DueDate:     issue.DueDate,
	}

	return s.issueRepository.CreateIssue(ctx, params)
}

func (s *IssueService) GetIssue(ctx context.Context) error {
	return s.issueRepository.GetIssue(ctx)
}

func (s *IssueService) GetIssuesList(ctx context.Context) error {
	return s.issueRepository.GetIssuesList(ctx)
}

func (s *IssueService) DeleteIssue(ctx context.Context) error {
	return s.issueRepository.DeleteIssue(ctx)
}

func (s *IssueService) UpdateIssue(ctx context.Context) error {
	return s.issueRepository.UpdateIssue(ctx)
}
