package http

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
)

type errorResponse struct {
	Error string `json:"error"`
}

func ErrorHandler(ctx fiber.Ctx, err error) error {
	var fiberErr *fiber.Error

	if errors.As(err, &fiberErr) {
		return ctx.Status(fiberErr.Code).JSON(errorResponse{
			Error: fiberErr.Message,
		})
	}

	switch {
	case errors.Is(err, domain.ErrProjectNameRequired),
		errors.Is(err, domain.ErrCommentTextRequired),
		errors.Is(err, domain.ErrProjectNameTooLong),
		errors.Is(err, domain.ErrIssueTitleRequired),
		errors.Is(err, domain.ErrInvalidIssueStatus),
		errors.Is(err, domain.ErrInvalidIssuePriority),
		errors.Is(err, domain.ErrIssueTitleTooLong):
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: err.Error(),
		})

	case errors.Is(err, domain.ErrIssueTitleAlreadyExists),
		errors.Is(err, domain.ErrProjectNameAlreadyExists):
		return ctx.Status(fiber.StatusConflict).JSON(errorResponse{
			Error: err.Error(),
		})

	case errors.Is(err, domain.ErrIssueNotFound),
		errors.Is(err, domain.ErrProjectNotFound):
		return ctx.Status(fiber.StatusNotFound).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("internal server error: %v", err)

	return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse{
		Error: "internal server error",
	})
}
