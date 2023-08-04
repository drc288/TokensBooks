package main

import (
	"log"

	"github.com/drc288/Drc.Microservice.Auth/api/routes"
	"github.com/drc288/Drc.Microservice.Auth/database"
	"github.com/drc288/Drc.Microservice.Auth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Aqui codigo de validacion de base de datos

	// Inicializacion de Fiber
	app := fiber.New()
	app.Use(cors.New())

	db := database.InitDb()
	db.AutoMigrate(&models.User{})

	// Health app - add middleware to validate first the connection to the database
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Send([]byte("Welcome to the application"))
	})

	routes.ServicesSetup(app)

	log.Fatal(app.Listen(":8080"))
}
