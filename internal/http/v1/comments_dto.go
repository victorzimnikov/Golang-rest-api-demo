package v1

import (
	"strings"
	"time"

	"github.com/victorzimnikov/Golang-rest-api-demo/internal/domain"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

type CreateCommentRequest struct {
	Text string
}

type IssueShort struct {
	ID    domain.IssueID `json:"id"`
	Title string         `json:"title"`
} // @name IssueShort

type CreateCommentDataResponse struct {
	ID        domain.CommentID `json:"id"`
	Text      string           `json:"text"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Issue     IssueShort       `json:"issue"`
} //	@name	Comment

type CreateCommentResponse = SuccessResponse[CreateCommentDataResponse] //	@name	CreateCommentResponse

type UpdateCommentRequest struct {
	Text *string `json:"text"`
}

type UpdateCommentDataResponse struct {
	ID        domain.CommentID `json:"id"`
	Text      string           `json:"text"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Issue     IssueShort       `json:"issue"`
} //	@name	Comment

func (r UpdateCommentRequest) toCommand(commentID domain.CommentID) (service.UpdateCommentCommand, error) {
	var text *string

	if r.Text != nil {
		normalizedText, err := domain.NormalizeCommentText(*r.Text)
		if err != nil {
			return service.UpdateCommentCommand{}, err
		}

		text = &normalizedText
	}

	return service.UpdateCommentCommand{
		CommentID: commentID,
		Text:      text,
	}, nil
}

type UpdateCommentResponse = SuccessResponse[UpdateCommentDataResponse] //	@name	UpdateCommentResponse

type GetIssueCommentsListRequest struct {
	SkipLimit

	Q string `query:"q"`
}

func (r GetIssueCommentsListRequest) toQuery() (service.GetIssueCommentsListCommand, error) {
	skip, limit, err := normalizeSkipLimit(r.Skip, r.Limit)
	if err != nil {
		return service.GetIssueCommentsListCommand{}, err
	}

	return service.GetIssueCommentsListCommand{
		Skip:  skip,
		Limit: limit,
		Q:     strings.TrimSpace(r.Q),
	}, nil
}

type IssueCommentListItemResponse struct {
	ID   domain.CommentID `json:"id"`
	Text string           `json:"text"`
} //	@name	ListIssueComment

type GetIssueCommentsListResponse = SuccessListResponse[IssueCommentListItemResponse] //	@name	GetIssueCommentsListResponse

type GetCommentDataResponse struct {
	ID        domain.CommentID `json:"id"`
	Text      string           `json:"text"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	Issue     IssueShort       `json:"issue"`
} //	@name	Comment

type GetCommentResponse = SuccessResponse[GetCommentDataResponse] //	@name	GetCommentResponse
