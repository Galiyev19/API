package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	Enviroment string `json:"enviroment"`
	Version    string `json:"version"`
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
	body, err := ioutil.ReadAll(configJson)
	if err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}

	if err := json.Unmarshal(body, cfg); err != nil {
		return fmt.Errorf("ERROR: %v", err)
	}
	return nil
}
