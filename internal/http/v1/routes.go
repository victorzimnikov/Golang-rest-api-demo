package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

func SetupV1Routes(api fiber.Router, services *service.Services) {
	router := api.Group("/v1")

	// Projects
	router.Get("/projects", GetProjectsList)
	router.Post("/projects", CreateProject)
	router.Get("/projects/:projectId", GetProject)
	router.Patch("/projects/:projectId", UpdateProject)
	router.Delete("/projects/:projectId", DeleteProject)
	router.Post("/projects/:projectId/issues", CreateProjectIssue)
	router.Get("/projects/:projectId/issues", GetProjectIssues)

	// Issues
	issueHandler := NewIssueHandler(services.CommentsService)

	router.Get("/issues/:issueId", issueHandler.GetIssue)
	router.Patch("/issues/:issueId", issueHandler.UpdateIssue)
	router.Delete("/issues/:issueId", issueHandler.DeleteIssue)
	router.Post("/issues/:issueId/comments", issueHandler.CreateIssueComment)
	router.Get("/issues/:issueId/comments", issueHandler.GetIssueComments)

	// Comments
	commentsHandlers := NewCommentsHandler(services.CommentsService)

	router.Get("/comments/:commentId", commentsHandlers.GetComment)
	router.Patch("/comments/:commentId", commentsHandlers.UpdateComment)
	router.Delete("/comments/:commentId", commentsHandlers.DeleteComment)
}
