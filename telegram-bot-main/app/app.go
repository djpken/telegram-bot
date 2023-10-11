package app

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type application struct {
	TelegramBot *telegramBotApi.BotAPI
	DB          *gorm.DB
	Cache       *redis.Client
}

var App = new(application)

func (app *application) DisConnect() {
	if app.DB == nil {
		return
	}
	db, err := app.DB.DB()
	if err != nil {
		panic(err)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
	err = app.Cache.Close()
	if err != nil {
		panic(err)
	}
}
