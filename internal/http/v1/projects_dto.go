package v1

import (
	"strings"
	"time"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProjectDataResponse struct {
	ID          domain.ProjectID `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
} //	@name	Project

type CreateProjectResponse = SuccessResponse[CreateProjectDataResponse] //	@name	CreateProjectResponse

type UpdateProjectRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type UpdateProjectDataResponse struct {
	ID          domain.ProjectID `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
} //	@name	Project

type UpdateProjectResponse = SuccessResponse[UpdateProjectDataResponse] //	@name	UpdateProjectResponse

type GetProjectDataResponse struct {
	ID          domain.ProjectID `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
} //	@name	Project

type GetProjectResponse = SuccessResponse[GetProjectDataResponse] //	@name	GetProjectResponse

type GetProjectsListRequest struct {
	SkipLimit

	Q string `query:"q"`
}

func (r GetProjectsListRequest) toQuery() (service.GetProjectsListCommand, error) {
	skip, limit, err := normalizeSkipLimit(r.Skip, r.Limit)
	if err != nil {
		return service.GetProjectsListCommand{}, err
	}

	return service.GetProjectsListCommand{
		Skip:  skip,
		Limit: limit,
		Q:     strings.TrimSpace(r.Q),
	}, nil
}

type ProjectListItemResponse struct {
	ID          domain.ProjectID `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
} //	@name	ListProject

type GetProjectsListResponse = SuccessListResponse[ProjectListItemResponse] //	@name	GetProjectsListResponse
