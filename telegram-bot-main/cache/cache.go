package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"telegram-bot/telegram-bot-main/env"
	"time"
)

type Cacher interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Close() error
}

type RedisCache struct {
	Client *redis.Client
	Ctx    context.Context
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.Client.Get(r.Ctx, key).Result()
}

func (r *RedisCache) Set(key string, value string, expiration time.Duration) error {
	return r.Client.Set(r.Ctx, key, value, expiration).Err()
}

func (r *RedisCache) Close() error {
	return r.Client.Close()
}

func NewRedisCache(ctx context.Context) Cacher {
	config := env.Environment.Cache
	cache := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	// 在這裡檢查連接是否正確，並返回錯誤（如果有的話）
	_, err := cache.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	log.Println("[App] Cache initialized")

	return &RedisCache{
		Client: cache,
		Ctx:    ctx,
	}
}
