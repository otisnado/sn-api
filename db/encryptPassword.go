package db

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	cost := 8
	byte, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(byte), err
}
