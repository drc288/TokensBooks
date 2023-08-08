package database

import "github.com/drc288/Drc.Microservice.Auth/models"

func ValidateEmail(email string) (models.User, bool) {
	db := Db

	var user models.User

	if err := db.First(&user, "email = ?", email); err != nil {
		return user, false
	}

	return user, true
}
