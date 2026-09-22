package service

import (
	"context"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"
)

type GetProjectsListQuery struct {
	Skip  int
	Limit int
	Q     string
}

type ProjectService struct {
	ProjectRepository *repository.ProjectRepository
}

func NewProjectService(projectRepository *repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		ProjectRepository: projectRepository,
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, data *domain.Project) (*domain.Project, error) {
	return s.ProjectRepository.SaveProject(ctx, data)
}

func (s *ProjectService) GetProjectByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	return s.ProjectRepository.GetProjectByID(ctx, id)
}

func (s *ProjectService) GetProjectsList(ctx context.Context, query GetProjectsListQuery) (int64, []domain.Project, error) {

	return s.ProjectRepository.GetProjectsList(ctx, db.GetProjectsListParams{
		Q:          query.Q,
		Skip:       int64(query.Skip),
		LimitCount: int32(query.Limit),
	})
}
