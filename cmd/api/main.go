package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	apiReadTimeout         = 5 * time.Second
	apiWriteTimeout        = 5 * time.Second
	apiIdleTimeout         = 30 * time.Second
	apiShutdownTimeout     = 10 * time.Second
)

// @title			REST Demo API
// @version		1.0
// @description	API для управления проектами, задачами и комментариями.
// @accept			json
// @produce		json
// @BasePath		/api/v1
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

	notifyContext, stopSignal := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopSignal()

	connectCtx, cancelConnect := context.WithTimeout(notifyContext, databaseConnectTimeout)
	defer cancelConnect()

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

	return startServer(notifyContext, config.ServerPort, services)
}

func startServer(ctx context.Context, port string, services *service.Services) error {

	app := fiber.New(fiber.Config{
		ErrorHandler: http.ErrorHandler,
		ReadTimeout:  apiReadTimeout,
		WriteTimeout: apiWriteTimeout,
		IdleTimeout:  apiIdleTimeout,
	})

	app.Hooks().OnListen(func(data fiber.ListenData) error {
		log.Printf("server started on %s:%s", data.Host, data.Port)

		return nil
	})

	http.SetupRoutes(app, services)

	if err := app.Listen(fmt.Sprintf(":%s", port), fiber.ListenConfig{
		GracefulContext: ctx,
		ShutdownTimeout: apiShutdownTimeout,
	}); err != nil {
		return fmt.Errorf("listen HTTP server: %w", err)
	}

	return nil
}
