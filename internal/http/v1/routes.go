package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

func SetupV1Routes(api fiber.Router, services *service.Services) {
	router := api.Group("/v1")

	// Projects
	projectHandler := NewProjectHandler(services.ProjectService)

	router.Get("/projects", projectHandler.GetProjectsList)
	router.Post("/projects", projectHandler.CreateProject)
	router.Get("/projects/:projectId", projectHandler.GetProject)
	router.Patch("/projects/:projectId", projectHandler.UpdateProject)
	router.Delete("/projects/:projectId", projectHandler.DeleteProject)

	// Issues
	issueHandler := NewIssueHandler(services.IssueService)

	router.Get("/issues/:issueId", issueHandler.GetIssue)
	router.Patch("/issues/:issueId", issueHandler.UpdateIssue)
	router.Delete("/issues/:issueId", issueHandler.DeleteIssue)
	router.Post("/issues/:issueId/comments", issueHandler.CreateIssueComment)
	router.Get("/issues/:issueId/comments", issueHandler.GetIssueComments)
	router.Post("/projects/:projectId/issues", issueHandler.CreateProjectIssue)
	router.Get("/projects/:projectId/issues", issueHandler.GetProjectIssuesList)

	// Comments
	commentsHandlers := NewCommentsHandler(services.CommentsService)

	router.Get("/comments/:commentId", commentsHandlers.GetComment)
	router.Patch("/comments/:commentId", commentsHandlers.UpdateComment)
	router.Delete("/comments/:commentId", commentsHandlers.DeleteComment)
}
