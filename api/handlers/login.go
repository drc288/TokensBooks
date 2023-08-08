package handlers

import (
	"time"

	"github.com/drc288/Drc.Microservice.Auth/database"
	"github.com/drc288/Drc.Microservice.Auth/jwt"
	"github.com/drc288/Drc.Microservice.Auth/models"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var l models.Login

	if err := c.BodyParser(&l); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("need username or password")
	}

	email := l.Email

	if !validEmail(email) {
		return c.Status(fiber.StatusBadRequest).SendString("mallformed email")
	}

	password := l.Password

	user, ok := database.Login(email, password)
	if !ok {
		return c.Status(fiber.StatusBadRequest).SendString("User or password invalid")
	}

	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error generating the jwt token: " + err.Error())
	}

	// Cookie
	expires := time.Now().Add(24 * time.Hour)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expires,
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "Success",
		"message": "Success login",
		"data":    jwtKey,
	})

}
