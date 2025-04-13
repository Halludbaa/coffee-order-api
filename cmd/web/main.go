package main

import (
	"api_setup/internal/config"
	"fmt"

	"github.com/go-playground/validator/v10"
)


func main(){
	viper := config.NewViper()
	app := config.NewGin()
	validate := validator.New()
	log := config.NewLogger(viper)

	defer log.Fatal("App Was Stopped!")

	log.Info("App Is Running!")
	config.Boostrap(&config.BoostrapConfig{
		App: app,
		Validate: validate,
		Log: log,
		Viper: viper,
	})
	
	err := app.Run(fmt.Sprintf(":%s", viper.GetString("web.port")))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	


}
