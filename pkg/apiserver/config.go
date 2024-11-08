package api

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	BindAddr    string `yaml:"port"`
	DataBaseURL string `yaml:"data_base_url"`
}

func NewConfig() *Config {
	return &Config{}
}

func InitConfig() (*Config, error) {

	// Открываем файл конфигурации
	file, err := os.Open("./configs/config.yaml") // путь к конфигу
	if err != nil {
		return nil, fmt.Errorf("unable to open config file: %w", err)
	}
	defer file.Close()

	// Создаем структуру для конфигурации
	var cfg Config

	// Декодируем YAML в структуру
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode YAML: %w", err)
	}

	// Логируем, что конфиг был успешно загружен
	log.Println("Config file loaded successfully!")

	return &cfg, nil
}
