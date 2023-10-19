package main

import (
	"context"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/bot"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/db"
	_ "telegram-bot/telegram-bot-main/env"
)

func main() {

	sys := app.NewApplication(bot.NewBot(), db.NewDataBase(), cache.NewRedisCache(context.Background()))

	telegramBot := sys.TelegramBot
	updates := bot.GetUpdates(telegramBot, bot.GetUpdateConfig(0, 60))
	bot.ListenUpdates(sys, updates)
	defer func() {
		_ = sys.Disconnect()
	}()
}
