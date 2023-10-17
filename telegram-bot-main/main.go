package main

import (
	"context"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/bot"
	_ "telegram-bot/telegram-bot-main/bot"
	_ "telegram-bot/telegram-bot-main/cache"
	_ "telegram-bot/telegram-bot-main/db"
	_ "telegram-bot/telegram-bot-main/env"
)

var cache = app.App.Cache
var ctx = context.Background()

func main() {

	telegramBot := app.App.TelegramBot
	updates := bot.GetUpdates(telegramBot, bot.GetUpdateConfig(0, 60))
	bot.ListenUpdates(updates)
	defer func() {
		app.App.DisConnect()
	}()
}
