package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)


func NewDB(viper *viper.Viper, log *logrus.Logger) *sqlx.DB {
	DB_HOST := viper.GetString("db.host")
	// DB_PORT := viper.GetString("db.port")
	DB_USER := viper.GetString("db.user")
	DB_PASSWORD := viper.GetString("db.password")
	DB_NAME := viper.GetString("db.name")

	ConnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME )

	db, err:= sqlx.Connect(viper.GetString("db.app"), ConnStr)


	if err != nil {
		log.Fatal(err)
	}

	return db
}