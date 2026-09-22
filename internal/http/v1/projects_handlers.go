package v1

import (
	"strconv"
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

	project, responseErr := h.projectService.CreateProject(ctx.Context(), projectBody)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusCreated)

	return ctx.JSON(CreateProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	})
}

func (h *ProjectHandler) GetProject(ctx fiber.Ctx) error {
	projectIDRaw := ctx.Params("projectId")
	projectID, err := strconv.ParseInt(projectIDRaw, 10, 64)
	if err != nil || projectID <= 0 {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid project id",
		)
	}

	project, responseErr := h.projectService.GetProjectByID(ctx.Context(), domain.ProjectID(projectID))
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(GetProjectResponse{
		ID:          project.ID,
		Name:        project.Name,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	})
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
