package database

import (
	"github.com/drc288/Drc.Microservice.Auth/models"
)

func ValidateEmail(email string) (models.User, bool) {
	db := Db

	var user models.User

	result := db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return user, false
	}

	return user, true
}
