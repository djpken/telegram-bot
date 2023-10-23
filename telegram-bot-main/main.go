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
	defer func() {
		_ = sys.Disconnect()
	}()
	updates := bot.GetUpdates(sys.TelegramBot, bot.GetUpdateConfig(0, 60))
	bot.ListenUpdates(sys.TelegramBot, sys.Cache, sys.DB, updates)
}
