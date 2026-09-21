package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/api"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/config"
)

func main() {
	config, err := config.LoadApiConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Hooks().OnListen(func(data fiber.ListenData) error {
		log.Printf("server started on %s:%s", data.Host, data.Port)

		return nil
	})

	api.SetupRoutes(app)

	if err := app.Listen(fmt.Sprintf(":%s", config.ServerPort)); err != nil {
		log.Fatal(err)
	}
}
