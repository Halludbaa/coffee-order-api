package utils

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(request string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(request), bcrypt.DefaultCost)
	if err != nil  {
		return "", fiber.ErrUnauthorized

	}

	return string(password), nil
}

func ValidatePassword(request string, hashed string) (error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(request)); err != nil{
		return fiber.ErrBadRequest
	}
	return nil
}