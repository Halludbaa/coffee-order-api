package config

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo(viper *viper.Viper, log *logrus.Logger) *mongo.Client {
	MONGO_HOST := viper.GetString("db.mongo.host")
	MONGO_PORT := viper.GetString("db.mongo.port")
	MONGO_PASSWORD := viper.GetString("db.mongo.password")
	MONGO_USER := viper.GetString("db.mongo.username")
	
	ConnStr := fmt.Sprintf("mongodb://%s:%s@%s:%s", MONGO_USER, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT)

	clientOption := options.Client().ApplyURI(ConnStr)
	
	client, err :=  mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	return client
}