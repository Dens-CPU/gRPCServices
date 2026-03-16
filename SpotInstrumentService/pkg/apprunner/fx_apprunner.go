package apprunner

import (
	spotAPI "Academy/gRPCServices/Protobuf/gen/spot"
	"Academy/gRPCServices/Shared/config"
	"Academy/gRPCServices/Shared/logger"
	"Academy/gRPCServices/Shared/opentelimetry"
	spotconfig "Academy/gRPCServices/SpotInstrumentService/config"
	"Academy/gRPCServices/SpotInstrumentService/internal/adapters/memory"
	redisadapter "Academy/gRPCServices/SpotInstrumentService/internal/adapters/redis"
	spothandlers "Academy/gRPCServices/SpotInstrumentService/internal/controllers/grpc_handlers"
	"Academy/gRPCServices/SpotInstrumentService/internal/usecase"
	grpcserver "Academy/gRPCServices/SpotInstrumentService/pkg/spotserver"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Подкючение логгера
func LoggerModul() fx.Option {
	return fx.Options(
		fx.Provide(
			func() (*zap.Logger, error) {
				logger, err := logger.NewLogger()
				if err != nil {
					return nil, fmt.Errorf("ошибка инициализации логгера:%w", err)
				}
				return logger, nil
			},
		),
		fx.Invoke(
			func(lc fx.Lifecycle, logger *zap.Logger) {
				lc.Append(fx.Hook{
					OnStop: func(ctx context.Context) error {
						_ = logger.Sync()
						return nil
					},
				})
			},
		),
	)
}

// Подключение In-memory хранилища
func StorageModul() fx.Option {
	var markets = []string{"Yandex Market", "OZON", "Wildberis", "AliExpress"}

	return fx.Options(
		fx.Provide(
			func(logger *zap.Logger) (*memory.Storage, error) {
				storage, err := memory.NewStorage(logger, markets)
				if err != nil {
					return nil, err
				}
				return storage, nil
			},
		),
		fx.Invoke(
			func(lc fx.Lifecycle, logger *zap.Logger, storage *memory.Storage) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							//Создание контекста для времени управления состояниями рынков
							ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
							defer cancel()
							msg := storage.AccessControl(ctx) //Запуск управления рынками
							logger.Info(msg)
						}()
						return nil
					},
				})
			},
		),
	)
}

// Получение конфига
func NewConfigModul() fx.Option {
	return fx.Provide(
		func() *config.ConfigLoader {
			loader := config.NewConfigLoader(globalPathToEnv, envFile, configType, pathToLocalEnv, pathToConfig)
			return loader
		},
		func(loader *config.ConfigLoader, logger *zap.Logger) (*spotconfig.Config, error) {
			config, err := config.NewConfig[spotconfig.Config](loader)
			if err != nil {
				logger.Error("ошибка получения конфига:",
					zap.Error(err),
				)
				return nil, err
			}
			return config, nil
		},
	)
}

// Подключение redis-кэш
func RedisModule() fx.Option {
	return fx.Provide(
		func(config *spotconfig.Config, logger *zap.Logger) (*redisadapter.RedisDB, error) {
			rctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			redis, err := redisadapter.NewRedis(rctx, config.Redis)
			if err != nil {
				logger.Error("ошибка инициализации redis:",
					zap.Error(err))
				return nil, err
			}
			return redis, nil
		},
	)
}

// Подключение трейсера
func TracingModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(logger *zap.Logger) (*trace.TracerProvider, error) {
				otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
					propagation.TraceContext{},
					propagation.Baggage{},
				))
				trace, err := opentelimetry.NewTrace(context.Background(), "spotService")
				if err != nil {
					logger.Error("ошибка инициализации трейсера:",
						zap.Error(err))
					return nil, err
				}
				return trace, nil
			},
		),
		fx.Invoke(
			func(lc fx.Lifecycle, trace *trace.TracerProvider) {
				lc.Append(fx.Hook{
					OnStop: func(ctx context.Context) error {
						return trace.Shutdown(ctx)
					},
				})
			},
		),
	)
}

// Подключение метрик
func MetricModul() fx.Option {
	return fx.Options(
		fx.Provide(
			func(logger *zap.Logger) (*sdkmetric.MeterProvider, error) {
				provider, err := opentelimetry.NewMetricPrometeus(context.Background(), "SpotService")
				if err != nil {
					logger.Error("ошибка инициализации метрик:",
						zap.Error(err))
					return nil, err
				}

				return provider, err
			},
		),
		fx.Invoke(
			func(lc fx.Lifecycle, logger *zap.Logger, provider *sdkmetric.MeterProvider) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							http.Handle("/metrics", promhttp.Handler())
							logger.Info("Запущен сервер для метрик на порту 9464...")
							err := http.ListenAndServe(":9464", nil)
							if err != nil {
								logger.Error("ошибка работы сервера для метрик:",
									zap.Error(err),
								)
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {

						return provider.Shutdown(ctx)
					},
				})
			},
		),
	)
}

// Подключение сервиса обработки
func ServiceModule() fx.Option {
	return fx.Provide(
		func(storage *memory.Storage, logger *zap.Logger, trace *trace.TracerProvider) *usecase.SpotService {
			return usecase.NewSpotInstrument(storage, logger, trace)
		},
	)
}

// Подключение обработчиков
func HandlersModule() fx.Option {
	return fx.Provide(
		func(service *usecase.SpotService) *spothandlers.Handlers {
			return spothandlers.NewHandlers(service)
		},
	)
}

// Создание gRPC сервера
func GrpcModule() fx.Option {
	return fx.Options(
		fx.Provide(
			func(redis *redisadapter.RedisDB, config *spotconfig.Config, logger *zap.Logger) (*grpcserver.Server, error) {
				grpcServer, err := grpcserver.New(redis, config.Server)
				if err != nil {
					logger.Error("ошибка инициализации grpc-сервера:",
						zap.Error(err),
					)
					return nil, err
				}
				return grpcServer, nil
			},
		),

		fx.Invoke(
			func(lc fx.Lifecycle, server *grpcserver.Server, handlers *spothandlers.Handlers, logger *zap.Logger) {
				//Регистрация grpc методов в сервере
				spotAPI.RegisterSpotInstrumentServiceServer(server, handlers)

				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go func() {
							logger.Info("Сервер запущен на порту 8080...")
							err := server.Serve(server.Listener)
							if err != nil {
								logger.Error("ошибка работы сервера:",
									zap.Error(err))
							}
						}()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						done := make(chan struct{})
						go func() {
							logger.Info("Остановка работы сервера...")
							server.GracefulStop()
							close(done)
						}()
						select {
						case <-done:
							logger.Info("Сервер остановлен")
							return nil
						case <-ctx.Done():
							logger.Info("Остановка серврера по таймауту")
							server.Stop()
							return ctx.Err()
						}

					},
				})
			},
		),
	)
}

func FxAppRunner() (*fx.App, error) {
	app := fx.New(
		LoggerModul(),
		StorageModul(),
		NewConfigModul(),
		RedisModule(),
		TracingModule(),
		MetricModul(),
		ServiceModule(),
		HandlersModule(),
		GrpcModule(),
	)

	if err := app.Err(); err != nil {
		return nil, fmt.Errorf("ошибка инициализации графа зависимостей:%w", err)
	}

	return app, nil
}
