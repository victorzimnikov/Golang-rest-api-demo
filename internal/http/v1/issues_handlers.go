package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

type IssueHandler struct {
	issueService *service.IssueService
}

func NewIssueHandler(issueService *service.IssueService) *IssueHandler {
	return &IssueHandler{
		issueService: issueService,
	}
}

// GetIssue returns a issue by ID.
//
//	@Summary			Get issue
//	@Description	Returns a issue by its identifier.
//	@Tags					Issues
//	@Produce			json
//	@Param				issueId						path	int	true	"Issue ID"	minimum(1)
//	@Success			200								{object}	GetIssueResponse
//	@Router				/issues/{issueId} [get]
func (h *IssueHandler) GetIssue(ctx fiber.Ctx) error {
	issueID, err := getIssueIDParam(ctx)
	if err != nil {
		return err
	}

	response, responseErr := h.issueService.GetIssue(ctx.Context(), issueID)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(GetIssueResponse{
		Data: GetIssueDataResponse{
			ID:          response.ID,
			Title:       response.Title,
			Description: response.Description,
			Status:      response.Status,
			Priority:    response.Priority,
			DueDate:     response.DueDate,
			CreatedAt:   response.CreatedAt,
			UpdatedAt:   response.UpdatedAt,
			Project:     ProjectShort(response.Project),
		},
	})
}

func (h *IssueHandler) UpdateIssue(ctx fiber.Ctx) error {
	// UpdateIssue
	return ctx.SendStatus(501)
}

// DeleteIssue a issue by ID.
//
//	@Summary			Delete issue
//	@Description	Delete a issue by its identifier.
//	@Tags					Issues
//	@Produce			json
//	@Param				issueId						path	int	true	"Issue ID"	minimum(1)
//	@Success			200								"Issue deleted"
//	@Router				/issues/{issueId} [delete]
func (h *IssueHandler) DeleteIssue(ctx fiber.Ctx) error {
	issueID, err := getIssueIDParam(ctx)
	if err != nil {
		return err
	}

	responseErr := h.issueService.DeleteIssue(ctx.Context(), issueID)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(SuccessResponse[*domain.Project]{
		Data: nil,
	})
}

func (h *IssueHandler) CreateIssueComment(ctx fiber.Ctx) error {
	// CreateIssueComment
	return ctx.SendStatus(501)
}

func (h *IssueHandler) GetIssueComments(ctx fiber.Ctx) error {
	// GetIssueComments
	return ctx.SendStatus(501)
}

// CreateProjectIssue create a project issue.
//
//	@Summary		Create project issue
//	@Description	Create a project issue.
//	@Tags			Issues
//	@Produce		json
//	@Param projectId path int true "Project ID" minimum(1)
//	@Param			request	body		CreateIssueRequest	true	"Issue data"
//	@Success		201		{object}	CreateIssueResponse
//	@Router			/projects/{projectId}/issues [POST]
func (h *IssueHandler) CreateProjectIssue(ctx fiber.Ctx) error {
	projectID, err := getProjectIDParam(ctx)
	if err != nil {
		return err
	}

	body, err := getRequestBody[CreateIssueRequest](ctx)
	if err != nil {
		return err
	}

	issueBody, err := domain.NewIssue(projectID, body.Title, body.Description, body.Status, body.Priority, body.DueDate)
	if err != nil {
		return err
	}

	response, responseErr := h.issueService.CreateIssue(ctx.Context(), issueBody)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusCreated)

	return ctx.JSON(CreateIssueResponse{
		Data: CreateIssueDataResponse{
			ID:          response.ID,
			Title:       response.Title,
			Description: response.Description,
			Status:      response.Status,
			Priority:    response.Priority,
			DueDate:     response.DueDate,
			CreatedAt:   response.CreatedAt,
			UpdatedAt:   response.UpdatedAt,
			Project: ProjectShort{
				ID:   response.Project.ID,
				Name: response.Project.Name,
			},
		},
	})
}

// GetProjectIssuesList returns a project issues list.
//
//	@Summary			Get project issues list
//	@Description	Returns a project issues list.
//	@Tags					Issues
//	@Produce			json
//	@Param				projectId	path			int			true	"Project ID"
//	@Param				skip			query			int			false	"Number of projects to skip"	default(0)	minimum(0)
//	@Param				limit			query			int			false	"Maximum number of projects"	default(10)	minimum(1)	maximum(50)
//	@Param				status		query			string	false	"Issue status" enums(open, in_progress, done)
//	@Param				priority	query			string	false	"Issue priority" enums(low, medium, high)
//	@Param				q					query			string	false	"Search query"
//	@Success			200				{object}	GetProjectIssuesListResponse
//	@Router				/projects/{projectId}/issues [get]
func (h *IssueHandler) GetProjectIssuesList(ctx fiber.Ctx) error {
	projectID, err := getProjectIDParam(ctx)
	if err != nil {
		return err
	}

	query, err := getRequestQuery[service.GetProjectIssuesListCommand, GetProjectIssuesListRequest](ctx)
	if err != nil {
		return err
	}

	total, listIssues, err := h.issueService.GetProjectIssuesList(ctx.Context(), service.GetProjectIssuesListCommand{
		Q:         query.Q,
		Skip:      query.Skip,
		Limit:     query.Limit,
		ProjectID: projectID,
		Status:    query.Status,
		Priority:  query.Priority,
	})
	if err != nil {
		return err
	}

	list := make([]ProjectIssueListItemResponse, len(listIssues))

	for idx, item := range listIssues {
		list[idx] = ProjectIssueListItemResponse(item)
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(GetProjectIssuesListResponse{
		List: list,
		Paginator: Paginator{
			Skip:  query.Skip,
			Limit: query.Limit,
			Size:  len(listIssues),
			Total: total,
		},
	})
}
