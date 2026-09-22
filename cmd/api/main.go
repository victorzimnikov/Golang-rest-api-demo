package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/config"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/database/db"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/http"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/repository"
	"github.com/victorzimnikov/Golang-rest-api-demo/internal/service"
)

const (
	databaseConnectTimeout = 5 * time.Second
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	config, err := config.LoadApiConfig()
	if err != nil {
		return err
	}

	connectCtx, cancel := context.WithTimeout(context.Background(), databaseConnectTimeout)
	defer cancel()

	pool, err := pgxpool.New(connectCtx, config.DatabaseURL)
	if err != nil {
		return fmt.Errorf("create PostgreSQL connection pool: %w", err)
	}
	defer pool.Close()

	if err := pool.Ping(connectCtx); err != nil {
		return fmt.Errorf("ping PostgreSQL: %w", err)
	}

	queries := db.New(pool)

	repositories := repository.NewRepositories(queries)
	services := service.NewServices(repositories)

	return startServer(config.ServerPort, services)
}

func startServer(port string, services *service.Services) error {

	app := fiber.New(fiber.Config{
		ErrorHandler: http.ErrorHandler,
	})

	app.Hooks().OnListen(func(data fiber.ListenData) error {
		log.Printf("server started on %s:%s", data.Host, data.Port)

		return nil
	})

	http.SetupRoutes(app, services)

	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		return fmt.Errorf("listen HTTP server: %w", err)
	}

	return nil
}
