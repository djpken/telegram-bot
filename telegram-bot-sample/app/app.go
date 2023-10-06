package app

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
)

type application struct {
	TelegramBot *telegramBotApi.BotAPI
	Gorm        *gorm.DB
}

var App = new(application)

func (app *application) DisConnect() {
	if app.Gorm == nil {
		return
	}
	db, err := app.Gorm.DB()
	if err != nil {
		panic(err)
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
}
