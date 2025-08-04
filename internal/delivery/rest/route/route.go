package route

import (
	"coffee/internal/delivery/rest/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type  RouteConfig struct {
	Viper				*viper.Viper
	App 				*fiber.App
	AuthMiddleware		fiber.Handler
}

func (c *RouteConfig) Setup(){
	c.SetupMiddleware()
}

func (c *RouteConfig) SetupMiddleware() {
	c.App.Use(middleware.CORSMiddleware(c.Viper))
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Get("/ping", func (ctx *fiber.Ctx) error  {
		return ctx.JSON(fiber.Map{
			"data": "pong",
		})
	})

}

func (c *RouteConfig) SetupAuthRoute() {
	auth := c.App.Group("/api")
	auth.Use(c.AuthMiddleware)

}
