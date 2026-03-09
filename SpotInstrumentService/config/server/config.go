package serverconfig

import (
	configfile "Academy/gRPCServices/Shared/config"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port    string `yaml:"port"`
		Host    string `yaml:"host"`
		Network string `yaml:"network"`
	} `yaml:"server"`
}

func NewConfig() (*Config, error) {

	file, err := configfile.NewConfigFile("./SpotInstrumentService/config/.env", "SPOT_CONFIG_PATH")
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
