package app

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"telegram-bot/telegram-sample1/env"
)

type application struct {
	Environment env.Env
	TelegramBot telegramBotApi.BotAPI
}

var App = new(application)
