package handlers

import (
	"fmt"
	"net/mail"
	"strconv"

	"github.com/drc288/Drc.Microservice.Auth/database"
	"github.com/drc288/Drc.Microservice.Auth/models"
	"github.com/gofiber/fiber/v2"
)

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func Registry(c *fiber.Ctx) error {
	var u models.User

	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Mallformed user")
	}

	email := u.Email

	if !validEmail(email) {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	password := u.Password
	if len(password) < 6 {
		return c.Status(fiber.StatusBadRequest).SendString("Password is necesari more than 6 characters")
	}

	// Las siguientes lineas validan la duplicacion del email pero desde la estructura del usuario ya se hacen estas validaciones
	// Se deja este codigo para ejemplos futuros
	_, findEmail := database.ValidateEmail(email)

	if findEmail {
		return c.Status(fiber.StatusBadRequest).SendString("Error email in database")
	}

	userId, status, err := database.RegistryUser(u)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error in the insertion of the user " + err.Error())
	}

	if !status {
		return c.Status(fiber.StatusBadRequest).SendString("Error in the insertion of the user in Database")
	}

	return c.Status(fiber.StatusCreated).SendString("User created" + fmt.Sprintf("%s", strconv.Itoa(userId)))

}
