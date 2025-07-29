package middleware

import (
	"coffee/internal/model/apperrors"

	"github.com/gin-gonic/gin"
)

func GuestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		
		if token == "" {
			token, _ = ctx.Cookie("Authorization")
			if token != "" {
				ctx.AbortWithStatusJSON(apperrors.Authorization, apperrors.NewAuthorization("you already login!"))
				return
			}
		}
		
		ctx.Next()
	}
} 