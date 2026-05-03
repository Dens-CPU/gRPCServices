package entryapiservice

const (
	GlobalPathToEnv = "."    //Дериктория, в которой находиться общий файл env
	EnvFile         = ".env" //Название файла .env
	ConfigType      = "yaml" //Тип конфиг файла (yaml)

	PathToLocalEnv = "PATH_APIGETWAY_CONFIG_ENV" //Переменная окружения, в которой лежит путь в локально env файлу
	PathToConfig   = "APIGETWAY_CONFIG_PATH"     //Переменная окружения, в которой лежит путь к конфигу
)
