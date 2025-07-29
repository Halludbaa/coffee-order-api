package helper

import (
	"api_setup/internal/model/apperrors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(request string) (string, *apperrors.Apperrors) {
	password, err := bcrypt.GenerateFromPassword([]byte(request), bcrypt.DefaultCost)
	if err != nil  {
		return "", apperrors.NewInternal()
	}

	return string(password), nil
}

func ValidatePassword(request string, hashed string) (*apperrors.Apperrors) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(request)); err != nil{
		return apperrors.NewBadRequest("failed to sign in", []apperrors.APIError{})
	}
	return nil
}