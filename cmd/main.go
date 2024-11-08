package main

import (
	api "API/pkg/apiserver"
	"API/pkg/handler"
	"API/pkg/repository"
	"API/pkg/service"
	"API/pkg/store"
	"log"
)

func main() {
	config, err := api.InitConfig()
	if err != nil {
		log.Fatalf("error initializing config: %s", err)
	}

	st := store.NewStore()
	err = st.Open(config.DataBaseURL)
	if err != nil {
		log.Fatalf("error initializing config: %s", err)
	}

	repos := repository.NewRepository()      // repos
	services := service.NewService(repos)    // services
	handlers := handler.NewHandler(services) // handlers

	srv := new(api.Server)
	if err := srv.Run(config.BindAddr, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured  while running http server: %s", err.Error())
	}
}
