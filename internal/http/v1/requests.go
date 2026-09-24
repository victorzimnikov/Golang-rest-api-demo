package v1

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
)

type SkipLimit struct {
	Skip  *int64 `query:"skip"`
	Limit *int32 `query:"limit"`
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

type queryRequest[Q any] interface {
	toQuery() (Q, error)
}

func getRequestQuery[Q any, T queryRequest[Q]](ctx fiber.Ctx) (Q, error) {
	var (
		request T
		zero    Q
	)

	if err := ctx.Bind().Query(&request); err != nil {
		return zero, fiber.NewErrorf(fiber.ErrBadRequest.Code, "invalid query parameters: %w", err)
	}

	return request.toQuery()
}

func normalizeSkipLimit(pSkip *int64, pLimit *int32) (int64, int32, error) {
	var skip int64 = 0
	var limit int32 = 10

	if pSkip != nil {
		skip = *pSkip
	}

	if pLimit != nil {
		limit = *pLimit
	}

	if skip < 0 {
		return 0, 0, fiber.NewError(fiber.ErrBadRequest.Code, "skip must be positive")
	}

	if limit < 1 || limit > 50 {
		return 0, 0, fiber.NewError(fiber.ErrBadRequest.Code, "limit must be between 1 and 50")
	}

	return skip, limit, nil
}
