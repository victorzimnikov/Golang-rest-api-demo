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
	case errors.Is(err, domain.ErrProjectNameRequired):
	case errors.Is(err, domain.ErrCommentTextRequired):
	case errors.Is(err, domain.ErrProjectNameTooLong):
		return ctx.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Error: err.Error(),
		})

	case errors.Is(err, domain.ErrProjectNotFound):
		return ctx.Status(fiber.StatusNotFound).JSON(errorResponse{
			Error: err.Error(),
		})
	}

	log.Printf("internal server error: %v", err)

	return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse{
		Error: "internal server error",
	})
}
