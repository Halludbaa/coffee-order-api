package config

import (
	"coffee/internal/delivery/rest/middleware"
	"coffee/internal/delivery/rest/route"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoostrapConfig struct {
	App			*fiber.App
	Log			*logrus.Logger
	Viper		*viper.Viper
	DB			*sqlx.DB
	Mongo		*mongo.Client
}

func Boostrap(config *BoostrapConfig) {
	authMiddleware := middleware.NewAuthMiddleware()

	router := route.RouteConfig{
		Viper: config.Viper,
		App: config.App,
		AuthMiddleware: authMiddleware,
	}

	router.Setup()
}