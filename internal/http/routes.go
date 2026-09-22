package http

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/timeout"
	v1 "github.com/victorzimnikov/Golang-rest-api-demo/internal/http/v1"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"

	swaggo "github.com/gofiber/contrib/v3/swaggo"

	_ "github.com/victorzimnikov/Golang-rest-api-demo/docs"
)

func SetupRoutes(app *fiber.App, services *service.Services) {
	app.Get("/swagger/*", swaggo.HandlerDefault)

	api := app.Group("/api", logger.New())

	api.Use(timeout.New(func(c fiber.Ctx) error {
		return c.Next()
	}, timeout.Config{
		Timeout: 2 * time.Second,
	}))

	v1.SetupV1Routes(api, services)
}
