package main

import (
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/botApi"
	_ "telegram-bot/telegram-bot-main/env"
	_ "telegram-bot/telegram-bot-main/gorm"
)

func main() {
	bot := app.App.TelegramBot
	updates := botApi.GetUpdates(bot, botApi.GetUpdateConfig(0, 60))
	botApi.ListenUpdates(updates)
	defer func() {
		app.App.DisConnect()
	}()
}
