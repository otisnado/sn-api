package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/otisnado/sn-api/models"
)

func JWTGenerator(t models.User) (string, error) {

	secretKey := []byte("Contrasena_01")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"website":   t.Web,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
