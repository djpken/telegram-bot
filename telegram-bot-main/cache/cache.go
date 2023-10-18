package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"telegram-bot/telegram-bot-main/env"
	"time"
)

// Cacher 定義了一個快取的接口
type Cacher interface {
	Get(key string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Close() error
}

type RedisCache struct {
	Client *redis.Client
}

func (r *RedisCache) Get(key string) (string, error) {
	return r.Client.Get(context.Background(), key).Result()
}

func (r *RedisCache) Set(key string, value string, expiration time.Duration) error {
	return r.Client.Set(context.Background(), key, value, expiration).Err()
}

func (r *RedisCache) Close() error {
	return r.Client.Close()
}

// NewRedisCache 創建一個新的 RedisCache 實例
func NewRedisCache() Cacher {
	config := env.Environment.Cache
	cache := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	// 在這裡檢查連接是否正確，並返回錯誤（如果有的話）是個好主意。
	_, err := cache.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err) // 或者您可以選擇返回錯誤而不是使用 log.Fatalf
	}

	log.Println("[App] Cache initialized")
	return &RedisCache{Client: cache}
}

// 在您的主程式或初始化部分，您可以這樣使用：
// app.App.Cache = cache.NewRedisCache()
