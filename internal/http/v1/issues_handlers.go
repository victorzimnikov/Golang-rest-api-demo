package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

type IssueHandler struct {
	commentsService *service.CommentsService
}

func NewIssueHandler(commentsService *service.CommentsService) *IssueHandler {
	return &IssueHandler{
		commentsService: commentsService,
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
