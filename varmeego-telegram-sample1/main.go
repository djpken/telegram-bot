package main

import (
	telegramBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

const varmeegoBotToken = "6599882879:AAH_Sv9BnYEBN9SmqX-8EqUDBdaZk_8_ZmI"

func main() {
	envToken, isFull := os.LookupEnv("TELEGRAM_API_TOKEN")
	if !isFull {
		envToken = varmeegoBotToken
	}
	bot, err := telegramBot.NewBotAPI(envToken)
	if err != nil {
		panic(err)
	}
	updateConfig := telegramBot.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		text := update.Message.Text
		id := update.Message.Chat.ID
		newMessage := telegramBot.NewMessage(id, text)
		_, _ = bot.Send(newMessage)
	}
}
