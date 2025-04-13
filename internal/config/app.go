package config

import (
	"api_setup/internal/delivery/rest/route"

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
	router := route.RouteConfig{
		App: config.App,
	}
	router.Setup()

}