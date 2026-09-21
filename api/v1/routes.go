package v1

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupV1Routes(api fiber.Router) {
	router := api.Group("/v1", logger.New())

	// Projects
	router.Get("/projects", GetProjectsList)
	router.Post("/projects", CreateProject)
	router.Get("/projects/:projectId", GetProject)
	router.Patch("/projects/:projectId", UpdateProject)
	router.Delete("/projects/:projectId", DeleteProject)
	router.Post("/projects/:projectId/issues", CreateProjectIssue)
	router.Get("/projects/:projectId/issues", GetProjectIssues)

	// Issues
	router.Get("/issues/:issueId", GetIssue)
	router.Patch("/issues/:issueId", UpdateIssue)
	router.Delete("/issues/:issueId", DeleteIssue)
	router.Post("/issues/:issueId/comments", CreateIssueComment)
	router.Get("/issues/:issueId/comments", GetIssueComments)

	// Comments
	router.Get("/comments/:commentId", GetComment)
	router.Patch("/comments/:commentId", UpdateComment)
	router.Delete("/comments/:commentId", DeleteComment)
}
