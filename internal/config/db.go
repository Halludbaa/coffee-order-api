package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)


func NewDB(viper *viper.Viper, log *logrus.Logger) *sqlx.DB {
	DB_CLIENT := viper.GetString("db.client");
	DB_HOST := viper.GetString(fmt.Sprintf("db.%s.host", DB_CLIENT))
	// DB_PORT := viper.GetString(fmt.Sprintf("db.%s.port", DB_CLIENT))
	DB_USER := viper.GetString(fmt.Sprintf("db.%s.user", DB_CLIENT))
	DB_PASSWORD := viper.GetString(fmt.Sprintf("db.%s.password", DB_CLIENT))
	DB_NAME := viper.GetString(fmt.Sprintf("db.%s.name", DB_CLIENT))

	ConnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME )

	db, err:= sqlx.Connect(DB_CLIENT, ConnStr)


	if err != nil {
		log.Fatal(err)
	}

	return db
}