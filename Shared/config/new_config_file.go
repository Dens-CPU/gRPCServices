package configfile

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func NewConfigFile(path_to_env string, envVarible string) (*os.File, error) {
	err := godotenv.Load(path_to_env)
	if err != nil {
		fmt.Println("Файл не найден")
		return nil, err
	}

	//Получение пути к конфигу
	path := os.Getenv(envVarible)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
