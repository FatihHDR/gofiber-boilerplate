package handler

import (
	"github.com/gofiber/fiber/v3"
	"myapp/internal/service"
)

type HealthHandler struct {
	service service.HealthService
}

func NewHealthHandler(service service.HealthService) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) Health(c fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(h.service.Check())
}
