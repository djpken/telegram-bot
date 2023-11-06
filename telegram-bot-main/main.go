package main

import (
	"telegram-bot/telegram-bot-main/bot"
	"telegram-bot/telegram-bot-main/wire"
)

func main() {
	sys, err := wire.InitializeApp()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = sys.Disconnect()
	}()
	updates := bot.GetUpdates(sys.TelegramBot, bot.GetUpdateConfig(0, 60))
	bot.ListenUpdates(sys.TelegramBot, sys.Cache, sys.DB, updates)
}
