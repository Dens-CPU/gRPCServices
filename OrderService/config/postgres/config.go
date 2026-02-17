package postgresconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Postgres Database `yaml:"postgres"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Sslmode  string `yaml:"sslmode"`
}

func NewConfig() (*Config, error) {
	//Загрузка файла .env
	err := godotenv.Load("./OrderService/config/.env")
	if err != nil {
		return nil, err
	}
	//Получение пути к конфигу
	path := os.Getenv("ORDER_CONFIG_PATH")

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("ошибка обработки конфига:%w", err)
	}
	return &cfg, nil
}
