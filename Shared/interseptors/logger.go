package interseptors

import (
	"Academy/gRPCServices/Shared/logger"
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Логирование запросов
func LoggerInterseptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	logger, _ := logger.NewLogger()

	//Изъятие ID-запроса из контекста
	requestID, _ := ctx.Value(requestIDKey).(string)

	//Логирование входящего запроса
	logger.Info("Запрос сервиса",
		zap.String("Номер запроса", requestID),
		zap.String("Вызываемый метод", info.FullMethod),
	)

	//Вызов следующего обработчика
	resp, err := handler(ctx, req)
	if err != nil {
		logger.Error("Ошибка выполнения запроса",
			zap.Error(err),
		)
	}
	return resp, err
}
