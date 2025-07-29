package config

import (
	"coffee/internal/delivery/rest/handlers"
	"coffee/internal/delivery/rest/middleware"
	"coffee/internal/delivery/rest/route"
	repo "coffee/internal/repositories/postgres/v1"
	"coffee/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type BoostrapConfig struct {
	App			*gin.Engine
	Log			*logrus.Logger
	Viper		*viper.Viper
	DB			*sqlx.DB
	Mongo		*mongo.Client
}
func Boostrap(config *BoostrapConfig) {
	userRepo := repo.NewUserRepo(config.DB, config.Log)
	sessionRepo := repo.NewSessionRepo(config.DB, config.Log)


	jwtServices := services.NewJWTServices(config.Viper, config.Log) // jwt service inital
	authServices := services.NewAuthServices(userRepo, config.Log, jwtServices, sessionRepo)

	authHandler := handlers.NewAuthHandler(config.Log, authServices)


	authMiddleware := middleware.NewAuthMiddleware(jwtServices)

	router := route.RouteConfig{
		Viper: config.Viper,
		App: config.App,
		AuthHandler: authHandler,
		AuthMiddleware: authMiddleware,
	}

	router.Setup()

}