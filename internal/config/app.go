package config

import (
	"api_setup/internal/delivery/rest/route"
	"api_setup/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BoostrapConfig struct {
	App			*gin.Engine
	Log			*logrus.Logger
	Validate 	*validator.Validate
	Viper		*viper.Viper
}
func Boostrap(config *BoostrapConfig) {
	_ = services.NewJWTServices(config.Viper) // jwt service inital
	
	router := route.RouteConfig{
		App: config.App,
	}

	router.Setup()

}