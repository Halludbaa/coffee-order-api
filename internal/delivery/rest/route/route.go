package route

import (
	"coffee/internal/delivery/rest/middleware"
	"coffee/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type  RouteConfig struct {
	Viper				*viper.Viper
	App 				*gin.Engine
	AuthHandler 		model.AuthHandler
	AuthMiddleware		gin.HandlerFunc
	UserLogMiddleware	gin.HandlerFunc
}

func (c *RouteConfig) Setup(){
	c.SetupMiddleware()
	c.SetupPing()
	c.SetupAPIv1()
}

func (c *RouteConfig) SetupMiddleware() {
	c.App.Use(gin.Recovery(), gin.Logger(), middleware.CORSMiddleware(c.Viper))
}

func (c *RouteConfig) SetupPing() {
	c.App.GET("/ping", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, model.NewWebResponse("pong", http.StatusOK))
		
	})
}

func (c *RouteConfig) SetupAPIv1() {
	v1 := c.App.Group("/api/v1")
	{
		c.SetupAuth(v1)
	}

	secure := v1.Group("/secure")
	secure.Use(c.AuthMiddleware)
	{
		secure.GET("/ping", func (ctx *gin.Context)  {
			msg, ok := ctx.Get("auth")
			if !ok {
				msg = "pong"
			}
			ctx.JSON(http.StatusOK, model.NewWebResponse(msg, http.StatusOK))
		})
	}
}

func (c *RouteConfig) SetupAuth(parent *gin.RouterGroup) {
	auth := parent.Group("")
	{
		auth.POST("/_refresh", c.AuthHandler.Refresh)

		guest := auth.Group("")
		guest.Use(middleware.GuestMiddleware())
		{
			guest.POST("/_sign_up", c.AuthHandler.SignUp)
			guest.POST("/_sign_in", c.AuthHandler.SignIn)
		}
		
		guarded := auth.Group("")
		guarded.Use(c.AuthMiddleware)
		{
			guarded.POST("/_logout", c.AuthHandler.Logout)
			guarded.GET("/_info", c.AuthHandler.Info)
		}
	}
}