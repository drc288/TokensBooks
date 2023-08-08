package jwt

import (
	"time"

	"github.com/drc288/Drc.Microservice.Auth/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(u models.User) (string, error) {
	jwtPassword := []byte("GO_CONTABLE")

	payload := jwt.MapClaims{
		"email":    u.Email,
		"username": u.Username,
		"_id":      u.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(jwtPassword)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
