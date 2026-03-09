package interseptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
