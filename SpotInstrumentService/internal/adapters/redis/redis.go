package redisadapter

import (
	redisconfig "Academy/gRPCServices/SpotInstrumentService/config/redis"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	client *redis.Client
}

func NewRedis(ctx context.Context) (*RedisDB, error) {

	//Получение параметров подключения из конфига
	cfg, err := redisconfig.NewConfig()
	if err != nil {
		return nil, err
	}

	//Формирование строки подключения
	addr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)

	//Инициализация нового Redis клиента
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	//Проверка подключения к БД redis
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisDB{client: rdb}, nil
}

func (r *RedisDB) Get(ctx context.Context, key string) (string, error) {
	resp, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("ошибка получения данных из кэша:%w", err)
	}
	return resp, nil
}

func (r *RedisDB) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	err := r.client.Set(ctx, key, string(value), ttl).Err()
	if err != nil {
		return err
	}
	return nil
}
