package router

import (
	"github.com/gofiber/fiber/v2"
	"myapp/internal/handler"
)

func Setup(app *fiber.App, healthHandler *handler.HealthHandler) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", healthHandler.Health)
}
