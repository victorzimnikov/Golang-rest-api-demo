package v1

import (
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

// GetProjectsList returns a projects list.
//
//	@Summary		Get projects list
//	@Description	Returns a projects list.
//	@Tags			Projects
//	@Produce		json
//	@Param			skip	query		int		false	"Number of projects to skip"	default(0)	minimum(0)
//	@Param			limit	query		int		false	"Maximum number of projects"	default(10)	minimum(1)	maximum(50)
//	@Param			q		query		string	false	"Search query"
//	@Success		200		{object}	GetProjectsListResponse
//	@Router			/projects [get]
func (h *ProjectHandler) GetProjectsList(ctx fiber.Ctx) error {
	query, err := getRequestQuery[service.GetProjectsListCommand, GetProjectsListRequest](ctx)
	if err != nil {
		return err
	}

	total, listProjects, err := h.projectService.GetProjectsList(ctx.Context(), query)
	if err != nil {
		return err
	}

	list := make([]ProjectListItemResponse, len(listProjects))

	for idx, item := range listProjects {
		list[idx] = ProjectListItemResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
		}
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(GetProjectsListResponse{
		List: list,
		Paginator: Paginator{
			Skip:  query.Skip,
			Limit: query.Limit,
			Size:  len(list),
			Total: total,
		},
	})
}

// CreateProject create a project.
//
//	@Summary		Create project
//	@Description	Create a project.
//	@Tags			Projects
//	@Produce		json
//	@Param			request	body		CreateProjectRequest	true	"Project data"
//	@Success		201		{object}	CreateProjectResponse
//	@Router			/projects/{projectId} [POST]
func (h *ProjectHandler) CreateProject(ctx fiber.Ctx) error {
	body, err := getRequestBody[CreateProjectRequest](ctx)
	if err != nil {
		return err
	}

	projectBody, err := domain.NewProject(body.Name, body.Description)
	if err != nil {
		return err
	}

	project, responseErr := h.projectService.CreateProject(ctx.Context(), projectBody)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusCreated)

	return ctx.JSON(CreateProjectResponse{
		Data: CreateProjectDataResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		},
	})
}

// GetProject returns a project by ID.
//
//	@Summary		Get project
//	@Description	Returns a project by its identifier.
//	@Tags			Projects
//	@Produce		json
//	@Param			projectId	path		int	true	"Project ID"	minimum(1)
//	@Success		200			{object}	GetProjectResponse
//	@Router			/projects/{projectId} [get]
func (h *ProjectHandler) GetProject(ctx fiber.Ctx) error {
	projectID, err := getProjectIDParam(ctx)
	if err != nil {
		return err
	}

	project, responseErr := h.projectService.GetProjectByID(ctx.Context(), projectID)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(GetProjectResponse{
		Data: GetProjectDataResponse{
			ID:          project.ID,
			Name:        project.Name,
			Description: project.Description,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		},
	})
}

// UpdateProject update a project.
//
//	@Summary		  Update project
//	@Description	Update a project.
//	@Tags			    Projects
//	@Produce		  json
//	@Param        projectId path int true "Project ID"
//	@Param			  request	body		UpdateProjectRequest	true	"Project data"
//	@Success		  200		{object}	UpdateProjectResponse
//	@Router			  /projects/{projectId} [PATCH]
func (h *ProjectHandler) UpdateProject(ctx fiber.Ctx) error {
	projectID, err := getProjectIDParam(ctx)
	if err != nil {
		return err
	}

	body, err := getRequestBody[UpdateProjectRequest](ctx)
	if err != nil {
		return err
	}

	if body.Description == nil && body.Name == nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid request body",
		)
	}

	command := service.UpdateProjectCommand{
		ID:          projectID,
		Name:        body.Name,
		Description: body.Description,
	}

	response, responseErr := h.projectService.UpdateProject(ctx.Context(), command)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(UpdateProjectResponse{
		Data: UpdateProjectDataResponse(*response),
	})
}

// DeleteProject a project by ID.
//
//	@Summary		Delete project
//	@Description	Delete a project by its identifier.
//	@Tags			Projects
//	@Produce		json
//	@Param			projectId	path	int	true	"Project ID"	minimum(1)
//	@Success		204			"Project deleted"
//	@Router			/projects/{projectId} [delete]
func (h *ProjectHandler) DeleteProject(ctx fiber.Ctx) error {
	projectID, err := getProjectIDParam(ctx)
	if err != nil {
		return err
	}

	responseErr := h.projectService.DeleteProject(ctx.Context(), projectID)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(SuccessResponse[*domain.Project]{
		Data: nil,
	})
}
