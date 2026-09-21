package service

import "github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"

type ProjectService struct {
	ProjectRepository *repository.ProjectRepository
}

func NewProjectService(projectRepository *repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		ProjectRepository: projectRepository,
	}
}
