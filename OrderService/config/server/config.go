package serverconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port    int    `yaml:"port"`
		Host    string `yaml:"host"`
		Network string `yaml:"network"`
	} `yaml:"server"`
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

	//Сериализация данных из config
	var cfg Config
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("ошибка обработки конфига:%w", err)
	}
	return &cfg, nil
}
