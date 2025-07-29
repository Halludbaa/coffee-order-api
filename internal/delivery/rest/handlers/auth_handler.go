package handlers

import (
	"coffee/internal/entity"
	"coffee/internal/model"
	"coffee/internal/model/apperrors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	log			*logrus.Logger
	service		model.AuthServices
}

func NewAuthHandler(log *logrus.Logger, service model.AuthServices) model.AuthHandler {
	return &AuthHandler{
		log, service,
	}
}

func (authH *AuthHandler) SignUp(ctx *gin.Context) {
	request := new(model.SignUpRequest)

	if err := ctx.ShouldBind(request); err != nil {
		errors := apperrors.GetValidateMessage(err)
		ctx.AbortWithStatusJSON(400, apperrors.NewBadRequest("bad request", errors))
		return
	}

	if request.Password != request.ConfirmPassword {
		ctx.AbortWithStatusJSON(400, apperrors.NewBadRequest("bad request", apperrors.PasswordNotMatch()))
		return
	}

	response, err := authH.service.SignUp(ctx, request)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}

	ctx.JSON(http.StatusCreated, model.WebResponse[*model.UserResponse]{
		Code: http.StatusCreated,
		Message: "successfully sign up!",
		Data: response,
	})
	
}

func (authH *AuthHandler) SignIn(ctx *gin.Context) {
	request := new(model.SignInRequest)

	if err := ctx.ShouldBind(request); err != nil {
		errors := apperrors.GetValidateMessage(err)
		ctx.AbortWithStatusJSON(400, apperrors.NewBadRequest("bad request", errors))
		return
	}

	if request.Username == request.Email {
		ctx.AbortWithStatusJSON(apperrors.BadRequest, apperrors.NewBadRequest("username or email can't be null", []apperrors.APIError{}))
		return
	}

	userAgent := ctx.GetHeader("User-Agent")
	response, err := authH.service.SignIn(ctx, request, userAgent)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}

	ctx.SetCookie("Authorization", response.AccessToken, 60 * 30, "", "", true, true)
	ctx.SetCookie("X-Refresh", response.RefreshToken, 60 * 60 * 24 * 7, "", "", true, true)


	ctx.JSON(http.StatusCreated, model.WebResponse[*model.SignInResponse]{
		Code: http.StatusCreated,
		Message: "successfully sign in!",
		Data: response,
	})
}

func (authH *AuthHandler) Logout(ctx *gin.Context) {
	token, err := ctx.Cookie("X-Refresh")
	if err != nil {
		ctx.AbortWithStatusJSON(apperrors.Authorization, apperrors.NewAuthorization("you don't have a token yet"))
		return
	}

	authH.log.Debugf("Token: %v", token)

	session := &entity.Session{
		Token: token,
	}

	if err := authH.service.Logout(ctx, session); err != nil {
		ctx.AbortWithStatusJSON(err.Code, err)
		return
	}

	ctx.SetCookie("Authorization", "", -1, "", "", true, true)
	ctx.SetCookie("X-Refresh", "", -1, "", "", true, true)
	
	ctx.JSON(http.StatusOK, model.NewWebResponse("Successfully logout", http.StatusOK))
}

func bindTokenInBody(ctx *gin.Context) string {
	request := new(model.SignInResponse)
	if err := ctx.ShouldBind(request); err != nil {
		return ""
	}

	return request.RefreshToken
}

func (authH *AuthHandler) Info(ctx *gin.Context) {
	user, ok := ctx.Get("auth")
	if (!ok) {
		ctx.AbortWithStatusJSON(apperrors.Authorization, apperrors.NewAuthorization("you don't have access for this endpoint"))
	}

	response, apperror := authH.service.Info(ctx, user.(string))
	if apperror != nil {
		ctx.AbortWithStatusJSON(apperror.Code, apperror)
		return
	}

	ctx.JSON(http.StatusCreated, model.WebResponse[*model.UserResponse]{
		Code: http.StatusCreated,
		Message: "ok!",
		Data: response,
	})
}

func (authH *AuthHandler) Refresh(ctx *gin.Context) {
	token, err := ctx.Cookie("X-Refresh")
	if err != nil {
		if token = bindTokenInBody(ctx); token == "" {
			ctx.AbortWithStatusJSON(apperrors.Authorization, apperrors.NewAuthorization("you don't have a token yet"))
			return
		}
	}

	authH.log.Debugf("Token: %v", token)

	response, apperror := authH.service.Refresh(ctx, token); 
	if apperror != nil {
		ctx.AbortWithStatusJSON(apperror.Code, apperror)
		return
	}

	ctx.SetCookie("Authorization", response.AccessToken, 60 * 30, "", "", true, true)

	ctx.JSON(http.StatusCreated, model.WebResponse[*model.SignInResponse]{
		Code: http.StatusCreated,
		Message: "you refresh it!",
		Data: response,
	})
}

