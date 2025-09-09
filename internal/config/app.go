package config

import (
	"coffee/internal/delivery/rest/handlers"
	"coffee/internal/delivery/rest/middleware"
	"coffee/internal/delivery/rest/route"
	"coffee/internal/repositories/postgres"
	"coffee/internal/services"
	"coffee/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
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
	Redis		*redis.Client
}

func Boostrap(config *BoostrapConfig) {
	tokenUtil := utils.NewTokenUtil(config.Viper, config.Redis)

	authMiddleware := middleware.NewAuthMiddleware(tokenUtil)

	menuRepo := postgres.NewMenuRepository(config.DB, config.Log)

	menuService := services.NewMenuService(menuRepo, config.Log)

	menuHandler := handlers.NewMenuHandler(menuService, config.Log)


	router := route.RouteConfig{
		Viper: config.Viper,
		App: config.App,
		AuthMiddleware: authMiddleware,
		MenuHandler: menuHandler,
	}

	router.Setup()
}