package database

import "github.com/drc288/Drc.Microservice.Auth/models"

func RegistryUser(user models.User) (int, bool, error) {
	db := Db

	user.Password, _ = EncryptPassowrd(user.Password)
	if err := db.Create(&user).Error; err != nil {
		return 0, false, err
	}

	return int(user.ID), true, nil
}
