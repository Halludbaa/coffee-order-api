package utils

import (
	"coffee/internal/model"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type TokenUtil struct {
	SecretKey []byte
	Redis     *redis.Client
}

type TokenClaims struct {
	Id string
	Role string
	jwt.RegisteredClaims
}

func NewTokenUtil(viper *viper.Viper, redis *redis.Client) *TokenUtil{
	return &TokenUtil{
		SecretKey: []byte(viper.GetString("jwt.secret_key")),
		Redis: redis,
	}
}

func (t *TokenUtil) CreateToken(ctx context.Context, auth model.Auth) (string, error) {
	claims := TokenClaims{
		Id: auth.Id,
		Role: auth.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour*24*30)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString(t.SecretKey)
	if err != nil {
		return "", err
	}

	_, err = t.Redis.SetEx(ctx, jwtToken, auth, time.Hour*24*30).Result()
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (t *TokenUtil) ParseToken(ctx context.Context, tokenString string) (*model.Auth, error) {
	claims := new(TokenClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return t.SecretKey, nil
	})

	if err != nil {
		return nil, fiber.ErrUnauthorized
	}

	if claims.ExpiresAt.Before(time.Now()) || !token.Valid {
		return nil, fiber.ErrUnauthorized
	}

	return &model.Auth{
		Id: claims.Id,
		Role: claims.Role,
	}, nil
}


