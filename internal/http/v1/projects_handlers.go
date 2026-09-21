package v1

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

func (h *ProjectHandler) GetProjectsList(ctx fiber.Ctx) error {
	// GetProjectsList
	return ctx.SendStatus(501)
}

func (h *ProjectHandler) CreateProject(ctx fiber.Ctx) error {
	var request CreateProjectRequest

	if err := ctx.Bind().Body(&request); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid request body",
		)
	}

	projectBody, err := domain.NewProject(request.Name, request.Description, time.Now())
	if err != nil {
		return err
	}

	_, err = h.projectService.CreateProject(ctx.Context(), projectBody)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (h *ProjectHandler) GetProject(ctx fiber.Ctx) error {
	// GetProject
	return ctx.SendStatus(501)
}

func (h *ProjectHandler) UpdateProject(ctx fiber.Ctx) error {
	// UpdateProject
	return ctx.SendStatus(501)
}

func (h *ProjectHandler) DeleteProject(ctx fiber.Ctx) error {
	// DeleteProject
	return ctx.SendStatus(501)
}

func (h *ProjectHandler) CreateProjectIssue(ctx fiber.Ctx) error {
	// CreateProjectIssue
	return ctx.SendStatus(501)
}

func (h *ProjectHandler) GetProjectIssues(ctx fiber.Ctx) error {
	// GetProjectIssues
	return ctx.SendStatus(501)
}
