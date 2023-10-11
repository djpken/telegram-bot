package cache

import (
	"github.com/redis/go-redis/v9"
	"log"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/env"
)

func init() {
	config := env.Environment.Cache
	cache := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
	app.App.Cache = cache
	log.Println("[App] Cache initialized")
}
