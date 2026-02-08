package interseptors

import (
	"context"
	"log"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const requestIDKey = "x-request-id"

// Добавления ID запроса
func XRequestID(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {

	var requestID string

	//Metadata представляет собой мапу map[string][]string, где ключ - это наименование заголовка, значения - это список значений заголовка
	md, ok := metadata.FromIncomingContext(ctx) //Получение метаданных из контекста запроса

	if ok { //Если удалось получить метаданные, то проверяем есть ли в нужном заголовке значение. Если значения нет, то создаем ID.
		values := md.Get(requestIDKey)
		if len(values) > 0 {
			requestID = values[0]
			ctx = context.WithValue(ctx, requestIDKey, requestID) //Сохранение в контесте заголовка со значением
		} else {
			requestID = uuid.NewString()
			ctx = context.WithValue(ctx, requestIDKey, requestID) //Сохранение в контесте заголовка со значением
		}
	}
	return handler(ctx, req) //Передача работы следующему обработчику
}

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

// Обработка паники
func UnaryPanicRecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("pacinc recover. Method:%s.Panic:%v.", info.FullMethod, r) //Логируем панику с указанием вызываемого метода, перехваченной паники.
			err = status.Errorf(codes.Internal, "interal servrer error")          //Указание номера ошикби
		}
	}()
	return handler(ctx, req)
}
