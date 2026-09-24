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

func (h *IssueHandler) GetIssue(ctx fiber.Ctx) error {
	// GetIssue
	return ctx.SendStatus(501)
}

func (h *IssueHandler) UpdateIssue(ctx fiber.Ctx) error {
	// UpdateIssue
	return ctx.SendStatus(501)
}

func (h *IssueHandler) DeleteIssue(ctx fiber.Ctx) error {
	// DeleteIssue
	return ctx.SendStatus(501)
}

func (h *IssueHandler) CreateIssueComment(ctx fiber.Ctx) error {
	// CreateIssueComment
	return ctx.SendStatus(501)
}

func (h *IssueHandler) GetIssueComments(ctx fiber.Ctx) error {
	// GetIssueComments
	return ctx.SendStatus(501)
}

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
			Project: CreateIssueDataProjectResponse{
				ID:   response.Project.ID,
				Name: response.Project.Name,
			},
		},
	})
}

func (h *IssueHandler) GetProjectIssues(ctx fiber.Ctx) error {
	// GetProjectIssues
	return ctx.SendStatus(501)
}
