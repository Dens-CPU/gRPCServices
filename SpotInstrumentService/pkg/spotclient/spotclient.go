package spotclient

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot"
	"Academy/gRPCServices/Shared/config"
	spotconfig "Academy/gRPCServices/SpotInstrumentService/config"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type Client struct {
	spotAPI.SpotInstrumentServiceClient
}

const (
	globalPathToEnv = "."    //Дериктория, в которой находиться общий файл env
	envFile         = ".env" //Название файла .env
	configType      = "yaml" //Тип конфиг файла (yaml)

	pathToLocalEnv = "PATH_SPOTSERVICE_CONFIG_ENV" //Переменная окружения, в которой лежит путь в локально env файлу
	pathToConfig   = "SPOT_CONFIG_PATH"            //Переменная окружения, в которой лежит путь к конфигу
)

func NewClient() (*Client, error) {
	loader := config.NewConfigLoader(globalPathToEnv, envFile, configType, pathToLocalEnv, pathToConfig)
	cfg, err := config.NewConfig[spotconfig.Config](loader)
	if err != nil {
		return nil, err
	}
	conn, err := grpc.NewClient(cfg.Server.Host+cfg.Server.Port, grpc.WithInsecure(), grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	if err != nil {
		return nil, err
	}

	client := spotAPI.NewSpotInstrumentServiceClient(conn)
	return &Client{client}, nil
}
