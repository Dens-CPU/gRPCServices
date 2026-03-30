package entryuserservice

const (
	GlobalPathToEnv = "."    //Дериктория, в которой находиться общий файл env
	EnvFile         = ".env" //Название файла .env
	ConfigType      = "yaml" //Тип конфиг файла (yaml)

	PathToLocalEnv = "PATH_USERSERVICE_CONFIG_ENV" //Переменная окружения, в которой лежит путь в локально env файлу
	PathToConfig   = "USERSERVICE_CONFIG_PATH"     //Переменная окружения, в которой лежит путь к конфигу
)
