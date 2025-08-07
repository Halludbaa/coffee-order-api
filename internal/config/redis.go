package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedis(viper *viper.Viper) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", viper.GetString("db.redis.host"), viper.GetInt32("db.redis.port")),
		DB: 0,
	})

	return client
}