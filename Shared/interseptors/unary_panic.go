package interseptors

import (
	"context"
	"runtime/debug"

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
				logger.Warn("panic recover",
					zap.String("Method:", info.FullMethod),
					zap.String("Стек:", string(debug.Stack())),
					zap.Any("Panic:", r),
				)
			}
		}()
		return handler(ctx, req)
	}
}
