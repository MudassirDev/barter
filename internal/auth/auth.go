package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwordString string) (string, error) {
	err := validatePassword(passwordString)
	if err != nil {
		return "", err
	}
	password := []byte(passwordString)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func validatePassword(password string) error {
	length := len(password)
	if length < 12 {
		return fmt.Errorf("password too short!")
	}
	if length > 60 {
		return fmt.Errorf("password too long!")
	}
	return nil
}
