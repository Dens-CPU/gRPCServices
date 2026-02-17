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
	cfg, err := redisconfig.NewConfig()
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisDB{client: rdb}, nil
}

func (r *RedisDB) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisDB) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return r.client.Set(ctx, key, value, ttl).Err()
}
