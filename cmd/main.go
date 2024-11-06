package main

import (
	api "API"
	"API/pkg/handler"
	"API/pkg/repository"
	"API/pkg/service"
	"log"

	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()      // repos
	services := service.NewService(repos)    // services
	handlers := handler.NewHandler(services) // handlers

	srv := new(api.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured  while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
