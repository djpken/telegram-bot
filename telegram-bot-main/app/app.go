package app

import (
	"fmt"
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"telegram-bot/telegram-bot-main/cache"
)

type Application struct {
	TelegramBot *telegramBotApi.BotAPI
	DB          *gorm.DB
	Cache       cache.Cacher
}

func NewApplication(bot *telegramBotApi.BotAPI, db *gorm.DB, c cache.Cacher) *Application {
	return &Application{
		TelegramBot: bot,
		DB:          db,
		Cache:       c,
	}
}

func (app *Application) Disconnect() error {
	var errs []string

	if app.DB != nil {
		db, err := app.DB.DB()
		if err != nil {
			errs = append(errs, err.Error())
		}
		if err = db.Close(); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if app.Cache != nil {
		if err := app.Cache.Close(); err != nil {
			errs = append(errs, err.Error())
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors while disconnecting: %v", errs)
	}

	return nil
}
