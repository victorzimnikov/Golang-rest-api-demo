package http

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	v1 "github.com/victorzimnikov/Golang-rest-api-demo/internal/http/v1"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

func SetupRoutes(app *fiber.App, services *service.Services) {
	api := app.Group("/api", logger.New())

	v1.SetupV1Routes(api, services)
}
