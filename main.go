package main

import (
	api "API/pkg/apiserver"
	"API/pkg/handler"
	"API/pkg/repository"
	"API/pkg/service"
	"API/pkg/store"
	"fmt"

	"github.com/sirupsen/logrus"
)

//	@title			API Documentation
//	@version		1.0
//	@description	This is a simple API documentation example
//	@host			localhost:8080

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	config, err := api.InitConfig()
	if err != nil {
		logrus.Fatalf("error initializing config: %s", err)
	}

	st := store.NewStore()
	db, err := st.Open(config.DataBaseURL)
	if err != nil {
		logrus.Fatalf("error initializing config: %s", err)
	}

	fmt.Println(config.DataBaseURL)

	repos := repository.NewRepository(db)    // repos
	services := service.NewService(repos)    // services
	handlers := handler.NewHandler(services) // handlers

	srv := new(api.Server)
	if err := srv.Run(config.BindAddr, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured  while running http server: %s", err.Error())
	}
}
