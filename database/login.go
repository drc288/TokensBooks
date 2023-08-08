package database

import (
	"github.com/drc288/Drc.Microservice.Auth/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	u, isActive := ValidateEmail(email)

	if !isActive {
		return u, false
	}

	passwordByte := []byte(password)
	passwordDB := []byte(u.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordByte)
	if err != nil {
		return u, false
	}

	return u, true
}
