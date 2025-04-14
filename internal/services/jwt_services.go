package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var (
	access_key = []byte(os.Getenv("ACCESS_KEY"))
	refresh_key = []byte(os.Getenv("REFRESH_KEY"))
) 

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTServices struct {
	access_key []byte
	refresh_key []byte
}

func NewJWTServices (viper *viper.Viper) *JWTServices {
	return &JWTServices{
		access_key: []byte(viper.GetString("jwt.key.access")),
		refresh_key: []byte(viper.GetString("jwt.key.refresh")),
	}
}


func (sJWT *JWTServices) GenerateAccessToken(userID string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(sJWT.access_key)
}

func  (sJWT *JWTServices) GenerateRefreshToken(userID string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(sJWT.refresh_key)
}

func  (sJWT *JWTServices)  ValidateAccessToken(tokenString string) (string, error) {
	
	claims := new(JWTClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return access_key, nil
	})

	if err != nil {
		return "", nil
	}


	if !token.Valid {
		return "", nil
	}

	
	return claims.UserID, nil
}

func  (sJWT *JWTServices) ValidateRefreshToken(tokenString string) (string, error) {
	claims := new(JWTClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return refresh_key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", nil
		}
		
		return "", nil
	}

	if !token.Valid {
		return "", nil
	}
	
	return claims.UserID, nil
}