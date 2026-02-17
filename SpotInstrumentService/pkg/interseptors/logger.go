package interseptors

import (
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

	// configLog := zap.Config{ //Создание конфига логгера
	// 	Level:         zap.NewAtomicLevelAt(zap.InfoLevel),
	// 	Development:   true,
	// 	Encoding:      "console",
	// 	EncoderConfig: zap.NewDevelopmentEncoderConfig(),
	// 	OutputPaths:   []string{"stdout"},
	// }
	logger, _ := zap.NewDevelopment()
	// logger, _ := configLog.Build() //Создание нового логгера
	defer logger.Sync() //Очистка буффера логгера после работы программы

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
