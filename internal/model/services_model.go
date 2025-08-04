package model

import (
	"coffee/internal/model/apperrors"
)


type JWTServices interface{
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(tokenString string) (string, *apperrors.Apperrors)
	ValidateRefreshToken(tokenString string) (string, *apperrors.Apperrors)
}
