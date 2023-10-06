package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"telegram-bot/telegram-sample1/env"
)

func main() {
	//bot := app.App.TelegramBot
	//updates := botApi.GetUpdates(bot, botApi.GetUpdateConfig(0, 60))
	//botApi.ListenUpdates(updates)
	db, err := gorm.Open(postgres.Open(env.Environment.Dsn), &gorm.Config{})
}
