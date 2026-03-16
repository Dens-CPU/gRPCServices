package appruner

import (
	orderconfig "Academy/gRPCServices/OrderService/config"
	"Academy/gRPCServices/OrderService/internal/adapters/notify"
	"Academy/gRPCServices/OrderService/internal/adapters/postgres"
	spotservice "Academy/gRPCServices/OrderService/internal/adapters/spot_service"
	orderhandlers "Academy/gRPCServices/OrderService/internal/controllers/grpc_handlers"
	"Academy/gRPCServices/OrderService/internal/usecase"
	"Academy/gRPCServices/OrderService/pkg/orderserver"
	orderAPI "Academy/gRPCServices/Protobuf/gen/order"
	"Academy/gRPCServices/Shared/config"
	"Academy/gRPCServices/Shared/logger"
	"Academy/gRPCServices/Shared/opentelimetry"
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

func AppRunner(ctx context.Context) error {
	//Инициализация нового логгера
	logger, err := logger.NewLogger()
	if err != nil {
		return fmt.Errorf("ошибка инициализации логгера:%w", err)
	}
	defer logger.Sync()

	//Получение нового конфига
	loader := config.NewConfigLoader(globalPathToEnv, envFile, configType, pathToLocalEnv, pathToConfig)
	cfg, err := config.NewConfig[orderconfig.Config](loader)
	if err != nil {
		logger.Error("ошибка получения нового конфига:",
			zap.Error(err),
		)
	}

	//Инициализаця хранилища postgres
	storage, err := postgres.NewDB(ctx, logger, cfg.Postgres)
	if err != nil {
		logger.Error("ошибка инициализации хранилища:",
			zap.Error(err),
		)
		return err
	}

	//Инициализация spotClient
	spotClient, err := spotservice.NewClient()
	if err != nil {
		logger.Error("ошибка инициализации хранилища:",
			zap.Error(err),
		)
		return err
	}

	//Инициализация сервиса нотификаций
	notify := notify.NewStatStorage()

	//Инициализация трейсера
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{}))

	trace, err := opentelimetry.NewTrace(ctx, "OrderSrevice")
	if err != nil {
		logger.Error("ошибка запуска трейсера:",
			zap.Error(err),
		)
		return err
	}

	defer trace.Shutdown(ctx)

	//Инициализация сервиса обработки
	service := usecase.NewOrderServ(storage, spotClient, notify, logger, trace)

	//Инициализация обработчиков
	handlers := orderhandlers.NewHandlers(service)

	//Инициализация нового gRPC сервера
	grpcServer, err := orderserver.New(cfg.Server)
	if err != nil {
		logger.Error("ошибка инициализации grpc-сервера:",
			zap.Error(err),
		)
		return err
	}

	//Регистрация методов
	orderAPI.RegisterOrderServiceServer(grpcServer, handlers)

	//Запуск работы сервера
	fmt.Println("Сервер работает на порту 8081...")
	err = grpcServer.Serve(grpcServer.Listener)
	if err != nil {
		logger.Error("ошибка инициализации grpc-сервера:",
			zap.Error(err),
		)
		return err
	}
	return nil
}
