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

// Создания ID запроса
func XRequestID(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		value := md.Get("x-request-id")
		if len(value) > 0 {
			requestID := value[0]
			ctx = context.WithValue(ctx, requestIDKey, requestID)
		} else {
			requestID := uuid.NewString()
			ctx = context.WithValue(ctx, requestIDKey, requestID)
		}
	}
	return handler(ctx, req)
}

// Логирование
func LoggerInterseptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	logger, _ := zap.NewDevelopment() // Создание нового логгера

	requesID := ctx.Value(requestIDKey).(string) //Изъятие ID запроса из контекста

	logger.Info("Запрос сервера", //Логирование входящего запроса
		zap.String("Номер запроса", requesID),
		zap.String("Вызываемы метод", info.FullMethod),
	)

	resp, err := handler(ctx, req) //Вызов следующего обработчика
	if err != nil {
		logger.Error("Ошибка выполнения запроса", //При ошибки работы обработчика выводиться ошибка
			zap.Error(err),
		)
	}
	return resp, err
}

// Обработка паник
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
