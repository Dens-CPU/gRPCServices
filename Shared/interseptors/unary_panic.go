package interseptors

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Обработка паники
func UnaryPanicRecoveryInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Info("panic recover",
					zap.String("Method:", info.FullMethod), //Метод, при котором была вызвана паника
					zap.Any("Panic:", r),                   //Сама паника
				)
			}
		}()
		return handler(ctx, req)
	}
}
