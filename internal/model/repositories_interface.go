package model

import (
	"api_setup/internal/entity"
	"context"
)

type UserRepository interface {
	Store(ctx context.Context, request *SignUpRequest)  (*entity.User, error) 
	Remove(ctx context.Context, username string) (error)
	Update(ctx context.Context, userUpdateReq *UserUpdateRequest) (error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	FindById(ctx context.Context, Id string) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}

type SessionRepo interface {
	Store(ctx context.Context, request *entity.Session) (error)
	Remove(ctx context.Context, request *entity.Session) (error)
	FindByUserId(ctx context.Context,  record *entity.Session) (error)
	FindByToken(ctx context.Context,  record *entity.Session) (error)
}