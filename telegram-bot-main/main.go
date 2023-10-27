package main

import (
	"telegram-bot/telegram-bot-main/bot"
	_ "telegram-bot/telegram-bot-main/env"
)

func main() {
	sys, err := InitializeApp()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = sys.Disconnect()
	}()
	updates := bot.GetUpdates(sys.TelegramBot, bot.GetUpdateConfig(0, 60))
	bot.ListenUpdates(sys.TelegramBot, sys.Cache, sys.DB, updates)
}
