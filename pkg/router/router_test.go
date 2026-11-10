package router_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"myapp/internal/handler"
	"myapp/internal/model"
	"myapp/internal/repository"
	"myapp/internal/service"
	"myapp/pkg/router"
)

func newTestApp() *fiber.App {
	app := fiber.New()
	healthHandler := handler.NewHealthHandler(
		service.NewHealthService(repository.NewHealthRepository()),
	)
	router.Setup(app, healthHandler)
	return app
}

func TestHealthRoute(t *testing.T) {
	resp, err := newTestApp().Test(httptest.NewRequest(fiber.MethodGet, "/api/v1/health", nil))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}

	var body model.HealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Status != "ok" {
		t.Fatalf("expected status %q, got %q", "ok", body.Status)
	}
}

func TestUnknownRouteReturns404(t *testing.T) {
	resp, err := newTestApp().Test(httptest.NewRequest(fiber.MethodGet, "/api/v1/unknown", nil))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusNotFound {
		t.Fatalf("expected status %d, got %d", fiber.StatusNotFound, resp.StatusCode)
	}
}
