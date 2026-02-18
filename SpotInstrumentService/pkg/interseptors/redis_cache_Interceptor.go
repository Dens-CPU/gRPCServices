package interseptors

import (
	redisadapter "Academy/gRPCServices/SpotInstrumentService/internal/adapters/redis"
	spoterrors "Academy/gRPCServices/SpotInstrumentService/internal/domain/errors"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"google.golang.org/grpc"
)

func RedisCacheInterceptor(cache *redisadapter.RedisDB, ttl time.Duration) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		//Формирование ключа
		var resp interface{}
		key, exist := ctx.Value(requestIDKey).(string) //Получение ID зароса из контекста
		if !exist {
			return resp, spoterrors.Unavailable_request_id
		}

		//Вызов следующего обработчика
		resp, err := handler(ctx, req)
		if err != nil {
			return resp, err
		}

		//Сохранение лога в кэш
		respBytes, _ := json.Marshal(resp)
		err = cache.Set(ctx, key, string(respBytes), ttl)
		if err != nil {
			return nil, fmt.Errorf("ошибка сохранения лога в кэш:%w", err)
		}

		return resp, nil
	}
}
