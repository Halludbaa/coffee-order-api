package model

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Refresh(ctx *gin.Context)
	Info(ctx *gin.Context)
}

type VtuberHandler interface {
	GetAll(ctx *gin.Context)
}