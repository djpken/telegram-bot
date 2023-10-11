package main

import (
	"context"
	"log"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/bot"
	_ "telegram-bot/telegram-bot-main/bot"
	_ "telegram-bot/telegram-bot-main/cache"
	_ "telegram-bot/telegram-bot-main/db"
	_ "telegram-bot/telegram-bot-main/env"
	"telegram-bot/telegram-bot-main/model"
)

var cache = app.App.Cache
var ctx = context.Background()

func main() {
	var helper model.Helper
	_ = helper.GetHelperById(1)
	err := cache.Set(ctx, "help", helper.String(), 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := app.App.TelegramBot
	updates := bot.GetUpdates(telegramBot, bot.GetUpdateConfig(0, 60))
	bot.ListenUpdates(updates)
	defer func() {
		app.App.DisConnect()
	}()
}
