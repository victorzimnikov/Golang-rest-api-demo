package service

import (
	"context"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"
)

type UpdateProjectCommand struct {
	ID          domain.ProjectID
	Name        *string
	Description *string
}

type GetProjectsListCommand struct {
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

func (s *ProjectService) CreateProject(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	return s.ProjectRepository.SaveProject(ctx, db.CreateProjectParams{
		Name:        project.Name,
		Description: project.Description,
	})
}

func (s *ProjectService) GetProjectByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	return s.ProjectRepository.GetProjectByID(ctx, id)
}

func (s *ProjectService) GetProjectsList(ctx context.Context, query GetProjectsListCommand) (int64, []domain.Project, error) {
	return s.ProjectRepository.GetProjectsList(ctx, db.GetProjectsListParams{
		Q:          query.Q,
		Skip:       int64(query.Skip),
		LimitCount: int32(query.Limit),
	})
}

func (s *ProjectService) DeleteProject(ctx context.Context, id domain.ProjectID) error {
	return s.ProjectRepository.DeleteProject(ctx, id)
}

func (s *ProjectService) UpdateProject(
	ctx context.Context,
	command UpdateProjectCommand,
) (*domain.Project, error) {
	var name *string

	if command.Name != nil {
		normalizedName, err := domain.NormalizeProjectName(*command.Name)
		if err != nil {
			return nil, err
		}

		name = &normalizedName
	}

	params := repository.UpdateProjectParams{
		ID:          command.ID,
		Name:        name,
		Description: command.Description,
	}

	return s.ProjectRepository.UpdateProject(ctx, params)
}
