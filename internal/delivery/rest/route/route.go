package route

import (
	"api_setup/internal/delivery/rest/middleware"
	"api_setup/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type  RouteConfig struct {
	App 				*gin.Engine
	AuthMiddleware		gin.HandlerFunc
}

func (c *RouteConfig) Setup(){
	c.SetupMiddleware()
	c.SetupPing()
}

func (c *RouteConfig) SetupMiddleware() {
	c.App.Use(gin.Recovery(), gin.Logger(), middleware.CORSMiddleware())
}

func (c *RouteConfig) SetupPing() {
	c.App.GET("/ping", func (ctx *gin.Context)  {
		ctx.JSON(http.StatusOK, model.NewWebResponse("pong", http.StatusOK))
		
	})
}