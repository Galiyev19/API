package app

import (
	"log"

	"API/internal/api"
	"API/internal/config"
	"API/internal/repository"
	"API/internal/service"
	"API/internal/storage"
)

func Run() {
	cfg := config.NewConfig()

	if err := cfg.InitConfig("config.json", cfg); err != nil {
		log.Fatalf("Error: %v", err)
	}

	db, err := storage.CreateSqlDB(cfg.StoreDriverSqlite, cfg.StorePath, cfg.MigrationPath)
	if err != nil {
		log.Fatalf("%v", err)
	}

	r := repository.NewRepo(db)
	s := service.NewService(r)
	api := api.NewApi(s, cfg)

	api.InitServer()
}
