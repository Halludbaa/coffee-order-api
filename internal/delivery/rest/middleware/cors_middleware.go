package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
)

func CORSMiddleware(viper *viper.Viper) fiber.Handler {
    return cors.New(cors.Config{
        AllowOrigins: viper.GetString("cors.origin"),
        AllowCredentials: viper.GetBool("cors.credentials"),
        AllowHeaders: viper.GetString("cors.headers"),
        AllowMethods: viper.GetString("cors.methods"),
    })   
}