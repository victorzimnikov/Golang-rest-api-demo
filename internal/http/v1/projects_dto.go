package v1

import (
	"time"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
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

type GetProjectListRequest struct {
	Skip  string
	Limit string
	Q     string
}
