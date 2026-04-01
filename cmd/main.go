package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"myapp/internal/config"
	"myapp/internal/handler"
	"myapp/internal/repository"
	"myapp/internal/service"
	"myapp/pkg/router"
)

func main() {
	cfg := config.Load()

	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	healthRepo := repository.NewHealthRepository()
	healthService := service.NewHealthService(healthRepo)
	healthHandler := handler.NewHealthHandler(healthService)

	router.Setup(app, healthHandler)

	addr := ":" + cfg.AppPort
	log.Printf("%s running on %s (%s)", cfg.AppName, addr, cfg.AppEnv)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
