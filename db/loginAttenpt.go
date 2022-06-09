package db

import (
	"github.com/otisnado/sn-api/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginAttempt(email string, password string) (models.User, bool) {
	use, found, _ := CheckUserExists(email)
	if !found {
		return use, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(use.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return use, false
	}
	return use, true
}
