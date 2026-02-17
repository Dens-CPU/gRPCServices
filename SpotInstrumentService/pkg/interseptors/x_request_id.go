package interseptors

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

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
