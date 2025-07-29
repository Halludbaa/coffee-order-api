package main

import (
	"coffee/internal/config"
	"context"
	"fmt"
)


func main(){
	viper := config.NewViper()
	log := config.NewLogger(viper)
	db := config.NewDB(viper, log)
	app := config.NewGin()
	mongo := config.NewMongo(viper, log)

	defer func ()  {
		db.Close()
		mongo.Disconnect(context.TODO())
		log.Fatal("App Was Stopped!")
	}()

	log.Info("App Is Running!")
	config.Boostrap(&config.BoostrapConfig{
		App: app,
		Log: log,
		Viper: viper,
		DB: db,
		Mongo: mongo,
	})
	
	err := app.Run(fmt.Sprintf(":%s", viper.GetString("web.port")))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	


}
