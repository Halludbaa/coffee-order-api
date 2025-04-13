package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(viper *viper.Viper) *logrus.Logger{
	logger := logrus.New()
	
	logger.SetLevel(logrus.Level(viper.GetInt32("log.level")))
	logger.SetFormatter(&logrus.JSONFormatter{})
	// file, _ := os.OpenFile(os.Getenv("LOG_FILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// logger.SetOutput(file)

	return logger
}