package helpers

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(value string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func GenerateId() uuid.UUID {
	return uuid.New()
}
