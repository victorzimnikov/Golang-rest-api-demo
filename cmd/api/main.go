package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/config"
)

func main() {
	config, err := config.LoadApiConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	log.Printf("server started on %s port", config.ServerPort)

	if err := app.Listen(fmt.Sprintf(":%s", config.ServerPort)); err != nil {
		log.Fatal(err)
	}
}
