package v1

import (
	"strings"
	"time"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

type CreateIssueRequest struct {
	Title       string
	Description string
	Status      *domain.IssueStatus
	Priority    *domain.IssuePriority
	DueDate     *time.Time
}

type ProjectShort struct {
	ID   domain.ProjectID `json:"id"`
	Name string           `json:"name"`
} //	@name	ProjectShort

type CreateIssueDataResponse struct {
	ID          domain.IssueID       `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Status      domain.IssueStatus   `json:"status"`
	Priority    domain.IssuePriority `json:"priority"`
	DueDate     *time.Time           `json:"dueDate"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	Project     ProjectShort         `json:"project"`
} //	@name	Issue

type CreateIssueResponse = SuccessResponse[CreateIssueDataResponse] //	@name	CreateIssueResponse

type GetIssueDataResponse struct {
	ID          domain.IssueID       `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Status      domain.IssueStatus   `json:"status"`
	Priority    domain.IssuePriority `json:"priority"`
	DueDate     *time.Time           `json:"dueDate"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
	Project     ProjectShort         `json:"project"`
} //	@name	Issue

type GetIssueResponse = SuccessResponse[GetIssueDataResponse] //	@name	GetIssueResponse

type GetProjectIssuesListRequest struct {
	SkipLimit

	Q        string               `query:"q"`
	Status   domain.IssueStatus   `query:"status"`
	Priority domain.IssuePriority `query:"priority"`
}

func (r GetProjectIssuesListRequest) toQuery() (service.GetProjectIssuesListCommand, error) {
	skip, limit, err := normalizeSkipLimit(r.Skip, r.Limit)
	if err != nil {
		return service.GetProjectIssuesListCommand{}, err
	}
	if r.Status != "" && !domain.ValidateIssueStatus(r.Status) {
		return service.GetProjectIssuesListCommand{}, domain.ErrInvalidIssueStatus
	}
	if r.Priority != "" && !domain.ValidateIssuePriority(r.Priority) {
		return service.GetProjectIssuesListCommand{}, domain.ErrInvalidIssuePriority
	}

	return service.GetProjectIssuesListCommand{
		Skip:     skip,
		Limit:    limit,
		Q:        strings.TrimSpace(r.Q),
		Status:   r.Status,
		Priority: r.Priority,
	}, nil
}

type ProjectIssueListItemResponse struct {
	ID       domain.IssueID       `json:"id"`
	Title    string               `json:"title"`
	Status   domain.IssueStatus   `json:"status"`
	Priority domain.IssuePriority `json:"priority"`
} //	@name	ListProjectIssue

type GetProjectIssuesListResponse = SuccessListResponse[ProjectIssueListItemResponse] //	@name	GetProjectIssuesListResponse
