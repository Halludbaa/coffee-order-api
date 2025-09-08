package middleware

import (
	"coffee/internal/model"
	"coffee/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware(tokenUtil *utils.TokenUtil) fiber.Handler {
	return func(ctx *fiber.Ctx)  error{
		tokenStr := ctx.Get("Authorization", "NOT_FOUND")
		
		auth, err := tokenUtil.ParseToken(ctx.UserContext(), tokenStr)
		if err != nil {
			return err
		}

		ctx.Locals("auth", auth)
		
		return ctx.Next()
	}
} 

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}

