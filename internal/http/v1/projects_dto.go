package v1

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateProjectResponse struct {
	ID          domain.ProjectID `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type GetProjectResponse struct {
	ID          domain.ProjectID `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type GetProjectsListRequest struct {
	SkipLimit

	Q string `query:"q"`
}

func (r *GetProjectsListRequest) toQuery() (service.GetProjectsListQuery, error) {
	skip := 0
	limit := 10

	if r.Skip != nil {
		skip = *r.Skip
	}

	if r.Limit != nil {
		limit = *r.Limit
	}

	if skip < 0 {
		return service.GetProjectsListQuery{}, fiber.NewError(fiber.ErrBadRequest.Code, "skip must be positive")
	}

	if limit < 1 || limit > 50 {
		return service.GetProjectsListQuery{}, fiber.NewError(fiber.ErrBadRequest.Code, "limit must be between 1 and 50")
	}

	return service.GetProjectsListQuery{
		Skip:  skip,
		Limit: limit,
		Q:     strings.TrimSpace(r.Q),
	}, nil
}

type ProjectListItemResponse struct {
	ID          domain.ProjectID `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
}

type GetProjectsListResponse = SuccessListResponse[ProjectListItemResponse]
