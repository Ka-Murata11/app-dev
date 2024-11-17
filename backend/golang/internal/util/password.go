package util

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordUtil interface {
	CreateHashPassword(password string) (string, error)
	CompareHashPassword(hashedPassword string, password string) error
}

type passwordUtil struct{}

func NewPasswordUtil() PasswordUtil {
	return &passwordUtil{}
}

func (u *passwordUtil) CreateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (u *passwordUtil) CompareHashPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
