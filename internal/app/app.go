package app

import (
	"fmt"
	"log"

	"API/internal/config"
	"API/internal/handlers"
	"API/internal/service"
	"API/internal/storage"
)

func Run() {
	cfg := config.NewConfig()

	if err := cfg.InitConfig("config.json", cfg); err != nil {
		log.Fatalf("Error: %v", err)
	}

	connStr := "user=postgres_db password=123456GG dbname=my_db host=localhost port=5432 sslmode=disable"

	db, err := storage.InitDB(cfg.StoreDriver, connStr, cfg.MigrationPath)
	if err != nil {
		log.Fatalf("%v", err)
	}
	s := service.NewService(db)
	h := handlers.NewHandler(s, cfg)

	router := h.InitRoutes()

	router.Run(":8080")
}
