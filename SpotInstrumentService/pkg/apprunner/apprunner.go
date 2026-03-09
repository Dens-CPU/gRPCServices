package apprunner

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot"
	"Academy/gRPCServices/Shared/logger"
	"Academy/gRPCServices/Shared/opentelimetry"
	"Academy/gRPCServices/SpotInstrumentService/internal/adapters/memory"
	redisadapter "Academy/gRPCServices/SpotInstrumentService/internal/adapters/redis"
	spothandlers "Academy/gRPCServices/SpotInstrumentService/internal/controllers/grpc_handlers"
	"Academy/gRPCServices/SpotInstrumentService/internal/usecase"
	grpcserver "Academy/gRPCServices/SpotInstrumentService/pkg/spotserver"
	"context"
	"fmt"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

func AppRunner() error {

	//Инициализация zap логгера
	logger, err := logger.NewLogger()
	if err != nil {
		return fmt.Errorf("ошибка инициализации логгера:%w", err)
	}
	defer logger.Sync()

	//Инициализация хранилища
	storage, err := memory.NewStorage(logger)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		//Создание контекста для времени управления состояниями рынков
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		msg := storage.AccessControl(ctx) //Запуск управления рынками
		logger.Info(msg)
	}()

	//Создание контекста для redis
	rctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Инициализация redis
	redis, err := redisadapter.NewRedis(rctx)
	if err != nil {
		logger.Error("ошибка инициализации redis:",
			zap.Error(err))
		return err
	}

	//Инициализация трейсера
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator( //Пропагатор для сериализации traceID из контекста
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	trace, err := opentelimetry.NewTrace(context.Background(), "spotService")
	if err != nil {
		logger.Error("ошибка инициализации трейсера:",
			zap.Error(err))
		return err
	}
	defer trace.Shutdown(context.Background())

	//Инициализая нового сервиса
	service := usecase.NewSpotInstrument(storage, logger, trace)

	//Инициализация обработчиков
	handlers := spothandlers.NewHandlers(service)

	//Инициализация нового grpc сервера
	grpcServer, err := grpcserver.New(redis)
	if err != nil {
		logger.Error("ошибка инициализации grpc-сервера:",
			zap.Error(err),
		)
		return err
	}

	//Регистрация grpc методов в сервере
	spotAPI.RegisterSpotInstrumentServiceServer(grpcServer, handlers)

	//Запуск работы сервера
	fmt.Println("Сервер запущен на порту 8080...")
	err = grpcServer.Serve(grpcServer.Listener)
	if err != nil {
		logger.Error("ошибка работы сервера:",
			zap.Error(err))
		return err
	}
	wg.Wait()
	return nil
}
