package services

import (
	"api_setup/internal/entity"
	"api_setup/internal/helper"
	"api_setup/internal/model"
	"api_setup/internal/model/apperrors"
	"api_setup/internal/model/converter"
	"context"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type AuthServices struct {
	repo 		model.UserRepository
	log			*logrus.Logger
	jwtServices model.JWTServices
	session  	model.SessionRepo
}

func NewAuthServices(repo model.UserRepository, log *logrus.Logger, jwtServices model.JWTServices, session model.SessionRepo) model.AuthServices {
	return &AuthServices{
		repo, log, jwtServices, session,
	}
}


func userConflictHandler(pgErr *pq.Error, request *model.SignUpRequest) *apperrors.Apperrors {
	var field, value string
			switch pgErr.Constraint {
			case "users_username_key":
				value = request.Username
				field = "username"
			case "users_email_key":
				value = request.Email
				field = "email"
			}
	return apperrors.NewConflict(field, value)
}


func (authS *AuthServices) SignUp(ctx context.Context, request *model.SignUpRequest) (out *model.UserResponse, err *apperrors.Apperrors) {
	request.Password, err = helper.HashPassword(request.Password)
	if err != nil {
		authS.log.Warn("password validation error")

		return &model.UserResponse{}, err
	}

	user, dbErr := authS.repo.Store(ctx, request)
	if  dbErr != nil {
		if pgErr, ok := dbErr.(*pq.Error); ok && pgErr.Code == "23505" {
			
			return nil,userConflictHandler(pgErr, request)
		}
		authS.log.Warn("Database Error!")
		return nil, apperrors.NewInternal()
	}

	return converter.UserToResponse(user), nil
}

func (authS *AuthServices) SignIn(ctx context.Context, request *model.SignInRequest, userAgent string) (out *model.SignInResponse, apperror *apperrors.Apperrors) {
	findBy := "username"
	if request.Username == "" {
		findBy = "email"
	}

	var user *entity.User
	var dbErr interface{}

	switch findBy{
	case "username":
		user, dbErr = authS.repo.FindByUsername(ctx, request.Username)
	case "email":
		user, dbErr = authS.repo.FindByEmail(ctx, request.Email)
	}

	if dbErr != nil {
		return nil, apperrors.NewInternal()
	}

	if err := helper.ValidatePassword(request.Password, user.Password); err != nil {
		return nil, err	
	}

	access_token, err := authS.jwtServices.GenerateAccessToken(user.Username)
	if err != nil {
		return nil, apperrors.NewInternal()
	}

	refresh_token, err := authS.jwtServices.GenerateRefreshToken(user.Username)
	if err != nil {
		return nil, apperrors.NewInternal()
	}

	session := &entity.Session{
		UserID: user.Id,
		Username: user.Username,
		UserAgent: userAgent,
		Token: refresh_token,
	}
	if dbErr = authS.session.Store(ctx, session); dbErr != nil{
		return nil, apperrors.NewInternal()
	}

	response := &model.SignInResponse{
		User: converter.UserToResponse(user),
		AccessToken: access_token,
		RefreshToken: refresh_token,
	}

	return response, nil
}

func (authS *AuthServices) Refresh(ctx context.Context, token string) (*model.SignInResponse, *apperrors.Apperrors) {
	session := &entity.Session{
		Token: token,
	}
	if err := authS.session.FindByToken(ctx, session); err != nil {
		return nil, err.(*apperrors.Apperrors)
	}

	access_token, err := authS.jwtServices.GenerateAccessToken(session.Username)
	if err != nil {
		return nil, apperrors.NewInternal()
	}

	response := &model.SignInResponse{
		AccessToken: access_token,
	}


	return response, nil
}

func (authS *AuthServices) Logout(ctx context.Context, request *entity.Session) *apperrors.Apperrors {
	if dbErr := authS.session.Remove(ctx, request); dbErr != nil {
		return dbErr.(*apperrors.Apperrors)
	}

	return nil
}

func (authS *AuthServices) Info(ctx context.Context, request string) (*model.UserResponse, *apperrors.Apperrors) {
	user, dbErr := authS.repo.FindByUsername(ctx, request)
	if dbErr != nil {
		return nil, dbErr.(*apperrors.Apperrors)
	}

	return converter.UserToResponse(user), nil
}