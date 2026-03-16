package appruner

const (
	globalPathToEnv = "."    //Дериктория, в которой находиться общий файл env
	envFile         = ".env" //Название файла .env
	configType      = "yaml" //Тип конфиг файла (yaml)

	pathToLocalEnv = "PATH_ORDERSERVICE_CONFIG_ENV" //Переменная окружения, в которой лежит путь в локально env файлу
	pathToConfig   = "ORDER_CONFIG_PATH"            //Переменная окружения, в которой лежит путь к конфигу
)
