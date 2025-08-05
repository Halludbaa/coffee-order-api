package middleware

import (
	"coffee/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func NewAuthMiddleware( *utils.TokenUtil) fiber.Handler {
	return func(ctx *fiber.Ctx)  error{
		_ = ctx.Get("Authorization", "NOT_FOUND")

		return ctx.Next()
	}
} 