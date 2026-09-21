package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	v1 "github.com/victorzimnikov/Golang-rest-api-demo/api/v1"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	v1.SetupV1Routes(api)
}
