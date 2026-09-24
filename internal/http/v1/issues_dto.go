package v1

import (
	"time"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
)

type CreateIssueRequest struct {
	Title       string
	Description string
	Status      *domain.IssueStatus
	Priority    *domain.IssuePriority
	DueDate     *time.Time
}

type CreateIssueDataProjectResponse struct {
	ID   domain.ProjectID `json:"id"`
	Name string           `json:"name"`
} //	@name	ProjectShort

type CreateIssueDataResponse struct {
	ID          domain.IssueID                 `json:"id"`
	Title       string                         `json:"title"`
	Description string                         `json:"description"`
	Status      domain.IssueStatus             `json:"status"`
	Priority    domain.IssuePriority           `json:"priority"`
	DueDate     *time.Time                     `json:"dueDate"`
	CreatedAt   time.Time                      `json:"createdAt"`
	UpdatedAt   time.Time                      `json:"updatedAt"`
	Project     CreateIssueDataProjectResponse `json:"project"`
} //	@name	Issue

type CreateIssueResponse = SuccessResponse[CreateIssueDataResponse] //	@name	CreateIssueResponse
