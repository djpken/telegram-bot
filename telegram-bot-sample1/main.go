package main

import (
	"telegram-bot/telegram-sample1/app"
	"telegram-bot/telegram-sample1/botApi"
	_ "telegram-bot/telegram-sample1/env"
	_ "telegram-bot/telegram-sample1/gorm"
)

func main() {
	bot := app.App.TelegramBot
	updates := botApi.GetUpdates(bot, botApi.GetUpdateConfig(0, 60))
	botApi.ListenUpdates(updates)
	defer func() {
		app.App.DisConnect()
	}()
}
