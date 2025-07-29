package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func CORSMiddleware(viper *viper.Viper) gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Header("Access-Control-Allow-Origin", viper.GetString("cors.origin"))
        c.Header("Access-Control-Allow-Credentials", viper.GetString("cors.credentials"))
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST, PATCH, GET, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        
        c.Next()
    }
}