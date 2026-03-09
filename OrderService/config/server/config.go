package serverconfig

import (
	configfile "Academy/gRPCServices/Shared/config"
	"fmt"

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

	path_to_env := "./OrderService/config/.env" //Путь к файлу .env
	envVarible := "ORDER_CONFIG_PATH"           //Переменная окружения

	file, err := configfile.NewConfigFile(path_to_env, envVarible)
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
