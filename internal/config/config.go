package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	StoreDriverPostgres string `json:"store_driver_postgres"`
	StoreDriverSqlite   string `json:"store_driver_sqlite"`
	MigrationPath       string `json:"migration_path"`
	StorePath           string `json:"store_path"`
	Enviroment          string `json:"enviroment"`
	Version             string `json:"version"`
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) InitConfig(cfgPath string, config *Config) error {
	configJson, err := os.Open(cfgPath)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	defer configJson.Close()
	body, err := io.ReadAll(configJson)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	if err := json.Unmarshal(body, cfg); err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}
	return nil
}
