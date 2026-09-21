package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

type CommentsHandler struct {
	commentsService *service.CommentsService
}

func NewCommentsHandler(commentsService *service.CommentsService) *CommentsHandler {
	return &CommentsHandler{
		commentsService: commentsService,
	}
}

func (h *CommentsHandler) GetComment(ctx fiber.Ctx) error {
	// GetComment
	return ctx.SendStatus(501)
}

func (h *CommentsHandler) UpdateComment(ctx fiber.Ctx) error {
	// UpdateComment
	return ctx.SendStatus(501)
}

func (h *CommentsHandler) DeleteComment(ctx fiber.Ctx) error {
	// DeleteComment
	return ctx.SendStatus(501)
}
