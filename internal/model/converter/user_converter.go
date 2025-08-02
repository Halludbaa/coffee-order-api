package converter

import (
	"coffee/internal/entity"
	"coffee/internal/model"
)

func SignUpToUser(signInReq *model.SignUpRequest) *entity.User {
	return &entity.User{
		Username: signInReq.Username,
		Email: signInReq.Email,
		Password: signInReq.Password,
	}
}

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		// Id: user.Id,
		Username: user.Username,
		Fullname: user.Fullname,
		Email: user.Email,
	}
}

