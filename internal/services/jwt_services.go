package services

import (
	"coffee/internal/model"
	"coffee/internal/model/apperrors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
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
	log	*logrus.Logger
}

func NewJWTServices (viper *viper.Viper, log *logrus.Logger) model.JWTServices {
	return &JWTServices{
		log: log,
		access_key: []byte(viper.GetString("jwt.key.access")),
		refresh_key: []byte(viper.GetString("jwt.key.refresh")),
	}
}


func (sJWT *JWTServices) GenerateAccessToken(userID string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
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

func  (sJWT *JWTServices)  ValidateAccessToken(tokenString string) (string, *apperrors.Apperrors) {
	
	claims := new(JWTClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return sJWT.access_key, nil
	})

	sJWT.log.Debug(err)
	if err != nil {
		return "", apperrors.NewAuthorization("expired or invalid token")
	}

	if !token.Valid {
		sJWT.log.Debug(token.Valid)
		return "", apperrors.NewAuthorization("expired or invalid token")
	}

	
	return claims.UserID, nil
}

func  (sJWT *JWTServices) ValidateRefreshToken(tokenString string) (string, *apperrors.Apperrors){
	claims := new(JWTClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return sJWT.refresh_key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", apperrors.NewAuthorization("invalid token")
		}
		
		return "", apperrors.NewAuthorization("expired or invalid token")
	}

	if !token.Valid {
		return "",  apperrors.NewAuthorization("expired or invalid token")
	}
	
	return claims.UserID, nil
}