package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
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

// CreateIssueComment create a issue comment.
//
//	@Summary		Create issue comment
//	@Description	Create a issue comment.
//	@Tags			Comments
//	@Produce		json
//	@Param issueId path int true "Issue ID" minimum(1)
//	@Param			request	body		CreateCommentRequest	true	"Issue data"
//	@Success		201		{object}	CreateCommentResponse
//	@Router			/issues/{issueId}/comments [POST]
func (h *CommentsHandler) CreateIssueComment(ctx fiber.Ctx) error {
	issueID, err := getIssueIDParam(ctx)
	if err != nil {
		return err
	}

	body, err := getRequestBody[CreateCommentRequest](ctx)
	if err != nil {
		return err
	}

	comment, err := domain.NewComment(issueID, body.Text)
	if err != nil {
		return err
	}

	response, responseErr := h.commentsService.CreateComment(ctx.Context(), comment)
	if responseErr != nil {
		return responseErr
	}

	ctx.Status(fiber.StatusCreated)

	return ctx.JSON(
		CreateCommentResponse{
			Data: CreateCommentDataResponse{
				ID:        response.ID,
				Text:      response.Text,
				CreatedAt: response.CreatedAt,
				UpdatedAt: response.UpdatedAt,
				Issue: IssueShort{
					ID:    response.Issue.ID,
					Title: response.Issue.Title,
				},
			},
		},
	)
}

func (h *CommentsHandler) GetIssueComments(ctx fiber.Ctx) error {
	// GetIssueComments
	return ctx.SendStatus(501)
}
