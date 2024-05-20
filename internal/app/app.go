package app

import (
	"log"

	"API/internal/config"
	"API/internal/handlers"
	"API/internal/service"
)

func Run() {
	cfg := config.NewConfig()

	if err := cfg.InitConfig("config.json", cfg); err != nil {
		log.Fatalf("Error: %v", err)
	}

	s := service.NewService()
	h := handlers.NewHandler(s, cfg)

	router := h.InitRoutes()

	router.Run(":8080")
}
