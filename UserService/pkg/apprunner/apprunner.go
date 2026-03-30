package apprunner

import (
	"context"

	"github.com/DencCPU/gRPCServices/Protobuf/gen/user_service"
	"github.com/DencCPU/gRPCServices/Shared/config"
	entryuserservice "github.com/DencCPU/gRPCServices/Shared/enter_points/entry_user_service"
	"github.com/DencCPU/gRPCServices/Shared/logger"
	userconfig "github.com/DencCPU/gRPCServices/UserService/config"
	"github.com/DencCPU/gRPCServices/UserService/internal/adapters/jwt"
	"github.com/DencCPU/gRPCServices/UserService/internal/adapters/postgres"
	userhandlers "github.com/DencCPU/gRPCServices/UserService/internal/controllers/grpc_handlers"
	"github.com/DencCPU/gRPCServices/UserService/internal/usecase"
	"github.com/DencCPU/gRPCServices/UserService/pkg/userserver"
	"go.uber.org/zap"
)

func AppRunner() error {
	//Создание логгера
	logger, err := logger.NewLogger()
	if err != nil {
		return err
	}

	//Создание нового конфига
	loader := config.NewConfigLoader(
		entryuserservice.GlobalPathToEnv,
		entryuserservice.EnvFile,
		entryuserservice.ConfigType,
		entryuserservice.PathToLocalEnv,
		entryuserservice.PathToConfig,
	)
	cfg, err := config.NewConfig[userconfig.Config](loader)
	storage, err := postgres.NewDB(context.Background(), logger, cfg.Postgres)
	if err != nil {
		logger.Error("ошибка инициализации хранилища:",
			zap.Error(err),
		)
		return err
	}

	//Создание нового JWT-сервиса
	jwt := jwt.NewJWT(cfg.JWT)

	//Создание сервиса обработки
	service := usecase.NewService(storage, logger, jwt)

	//Создание обработчиков
	handler := userhandlers.NewHandler(service)

	server, err := userserver.NewServer(cfg.Server, logger)
	if err != nil {
		logger.Error("ошибка создания нового grpc сервера:",
			zap.Error(err),
		)
		return err
	}
	user_service.RegisterUserServiceServer(server, handler)

	logger.Info("Работа сервера:",
		zap.String("Сервер запущен на порту ", cfg.Server.Port),
	)
	err = server.Serve(server.Listener)
	if err != nil {
		logger.Error("Ошибка работы сервера:",
			zap.Error(err),
		)
	}
	return nil
}
