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

// GetIssueCommentsList returns a issue comments list.
//
//	@Summary			Get issue comments list
//	@Description	Returns a issue comments list.
//	@Tags					Comments
//	@Produce			json
//	@Param				issueId	path			int			true	"Issue ID"
//	@Param				skip			query			int			false	"Number of comments to skip"	default(0)	minimum(0)
//	@Param				limit			query			int			false	"Maximum number of comments"	default(10)	minimum(1)	maximum(50)
//	@Param				q					query			string	false	"Search query"
//	@Success			200				{object}	GetIssueCommentsListResponse
//	@Router				/issues/{issueId}/comments [get]
func (h *CommentsHandler) GetIssueCommentsList(ctx fiber.Ctx) error {
	issueID, err := getIssueIDParam(ctx)
	if err != nil {
		return err
	}

	query, err := getRequestQuery[service.GetIssueCommentsListCommand, GetIssueCommentsListRequest](ctx)
	if err != nil {
		return err
	}

	total, listIssues, err := h.commentsService.GetIssueCommentsList(ctx.Context(), service.GetIssueCommentsListCommand{
		Q:       query.Q,
		Skip:    query.Skip,
		Limit:   query.Limit,
		IssueID: issueID,
	})
	if err != nil {
		return err
	}

	list := make([]IssueCommentListItemResponse, len(listIssues))

	for idx, item := range listIssues {
		list[idx] = IssueCommentListItemResponse{
			ID:   item.ID,
			Text: item.Text,
		}
	}

	ctx.Status(fiber.StatusOK)

	return ctx.JSON(GetIssueCommentsListResponse{
		List: list,
		Paginator: Paginator{
			Skip:  query.Skip,
			Limit: query.Limit,
			Size:  len(listIssues),
			Total: total,
		},
	})
}
