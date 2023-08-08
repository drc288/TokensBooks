package routes

import (
	"github.com/drc288/Drc.Microservice.Auth/api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ServicesSetup(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/health", handlers.Healt)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/registry", handlers.Registry)
}
