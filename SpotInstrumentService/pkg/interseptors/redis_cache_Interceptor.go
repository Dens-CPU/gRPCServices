package interseptors

import (
	redisadapter "Academy/gRPCServices/SpotInstrumentService/internal/adapters/redis"
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
		reqBytes, _ := json.Marshal(req)
		key := fmt.Sprintf("grpc_cache:%s:%s", info.FullMethod, string(reqBytes))

		//Проверка ключа в кеше
		cached, err := cache.Get(ctx, key)
		if err == nil {
			var resp interface{}
			json.Unmarshal([]byte(cached), &resp)
			return resp, nil
		}

		//Если ключа нет то вызывается следующий обработчик
		resp, err := handler(ctx, req)
		if err != nil {
			return resp, err
		}

		//Сохранение нового лога в кеш
		respBytes, _ := json.Marshal(resp)
		_ = cache.Set(ctx, key, string(respBytes), ttl)

		return resp, nil
	}
}
