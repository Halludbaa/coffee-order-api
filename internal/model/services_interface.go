package model

import (
	"api_setup/internal/entity"
	"api_setup/internal/model/apperrors"
	"context"
)

type AuthServices interface {
	SignUp(ctx context.Context, request *SignUpRequest) (*UserResponse, *apperrors.Apperrors)
	SignIn(ctx context.Context, request *SignInRequest, userAgent string) (*SignInResponse, *apperrors.Apperrors)
	Refresh(ctx context.Context, token string) (*SignInResponse, *apperrors.Apperrors)
	Logout(ctx context.Context, request *entity.Session) *apperrors.Apperrors
	Info(ctx context.Context, request string) (*UserResponse, *apperrors.Apperrors)
}


type JWTServices interface{
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(tokenString string) (string, *apperrors.Apperrors)
	ValidateRefreshToken(tokenString string) (string, *apperrors.Apperrors)
}
