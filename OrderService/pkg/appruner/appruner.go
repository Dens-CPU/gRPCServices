package apprunner

import (
	"Academy/gRPCServices/OrderService/internal/adapters/notify"
	"Academy/gRPCServices/OrderService/internal/adapters/postgres"
	spotservice "Academy/gRPCServices/OrderService/internal/adapters/spot_service"
	orderhandlers "Academy/gRPCServices/OrderService/internal/controllers/grpc_handlers"
	"Academy/gRPCServices/OrderService/internal/usecase"
	"Academy/gRPCServices/OrderService/pkg/orderserver"
	orderAPI "Academy/gRPCServices/Protobuf/gen/order"
	"Academy/gRPCServices/Shared/logger"
	"Academy/gRPCServices/Shared/opentelimetry"
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

func AppRunner(ctx context.Context) error {
	logger, err := logger.NewLogger()
	if err != nil {
		return fmt.Errorf("ошибка инициализации логгера:%w", err)
	}
	defer logger.Sync()

	storage, err := postgres.NewDB(ctx, logger) //Инициализаця хранилища postgres
	if err != nil {
		logger.Error("ошибка инициализации хранилища:",
			zap.Error(err),
		)
		return err
	}

	spotClient, err := spotservice.NewClient()
	if err != nil {
		logger.Error("ошибка инициализации хранилища:",
			zap.Error(err),
		)
		return err
	}

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

	service := usecase.NewOrderServ(storage, spotClient, notify, logger, trace)

	handlers := orderhandlers.NewHandlers(service)

	grpcServer, err := orderserver.New()
	if err != nil {
		logger.Error("ошибка инициализации grpc-сервера:",
			zap.Error(err),
		)
		return err
	}

	orderAPI.RegisterOrderServiceServer(grpcServer, handlers)

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
