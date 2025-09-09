package route

import (
	"coffee/internal/delivery/rest/handlers"
	"coffee/internal/delivery/rest/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type  RouteConfig struct {
	Viper				*viper.Viper
	App 				*fiber.App
	AuthMiddleware		fiber.Handler
	MenuHandler			*handlers.MenuHandler
}

func (c *RouteConfig) Setup(){
	c.SetupMiddleware()
	c.SetupMenuRoute()
}

func (c *RouteConfig) SetupMiddleware() {
	c.App.Use(middleware.CORSMiddleware(c.Viper))
}

func (c *RouteConfig) SetupMenuRoute() {
	c.App.Post("/api/menu", c.MenuHandler.CreateMenuItem)
	c.App.Get("/api/menu", c.MenuHandler.GetAllMenuItems)
	c.App.Get("/api/menu/:id", c.MenuHandler.GetMenuItemByID)
	c.App.Put("/api/menu/:id", c.MenuHandler.UpdateMenuItem)
	c.App.Delete("/api/menu/:id", c.MenuHandler.DeleteMenuItem)

	c.App.Post("/api/stores/:storeId/menu", c.MenuHandler.AddToStoreMenu)
	c.App.Get("/api/stores/:storeId/menu", c.MenuHandler.GetStoreMenu)
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
