package redisconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Redis Redis `yaml:"redis"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func NewConfig() (*Config, error) {
	// Загружаем .env
	err := godotenv.Load("./SpotInstrumentService/config/.env")
	if err != nil {
		return nil, fmt.Errorf("cannot load .env: %w", err)
	}

	// Берем путь к yaml
	path := os.Getenv("CONFIG_PATH")

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("файл не найден: %w", err)
	}
	defer file.Close()

	var cfg Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка сериализации конфиг файла: %w", err)
	}

	return &cfg, nil
}
