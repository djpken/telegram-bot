package main

import (
	"telegram-bot/telegram-sample/app"
	"telegram-bot/telegram-sample/botApi"
	_ "telegram-bot/telegram-sample/env"
	_ "telegram-bot/telegram-sample/gorm"
)

func main() {
	bot := app.App.TelegramBot
	updates := botApi.GetUpdates(bot, botApi.GetUpdateConfig(0, 60))
	botApi.ListenUpdates(updates)
	defer func() {
		app.App.DisConnect()
	}()
}
