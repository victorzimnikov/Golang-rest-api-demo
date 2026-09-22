package service

import (
	"context"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"
)

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
