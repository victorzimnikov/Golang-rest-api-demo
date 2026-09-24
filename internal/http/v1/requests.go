package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
)

type SkipLimit struct {
	Skip  *int `query:"skip"`
	Limit *int `query:"limit"`
}

func getProjectIDParam(ctx fiber.Ctx) (domain.ProjectID, error) {
	projectIDRaw := ctx.Params("projectId")
	projectID, err := strconv.ParseInt(projectIDRaw, 10, 64)
	if err != nil || projectID <= 0 {
		return 0, fiber.NewError(
			fiber.StatusBadRequest,
			"invalid project id",
		)
	}

	return domain.ProjectID(projectID), nil
}

func getRequestBody[T any](ctx fiber.Ctx) (*T, error) {
	var request T

	if err := ctx.Bind().Body(&request); err != nil {
		return nil, fiber.NewError(
			fiber.StatusBadRequest,
			"invalid request body",
		)
	}

	return &request, nil
}
