package middleware

import (
	"coffee/internal/model"
	"coffee/internal/model/apperrors"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(jwtServices model.JWTServices) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		
		if token == "" {
			//Temporary Code
			var err error
			token, err = ctx.Cookie("Authorization")
			if err != nil {
				ctx.AbortWithStatusJSON(apperrors.Authorization, apperrors.NewAuthorization("you don't have a token yet"))
				return
			}
			// End Temporary
		}
		
		userID, err := jwtServices.ValidateAccessToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Code, err)
			return
		}
		
		ctx.Set("auth", userID)
		ctx.Next()
	}
}
