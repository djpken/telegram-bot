package bot

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/env"
	dao "telegram-bot/telegram-bot-main/serv"
)

const (
	HELLO = "hello"
	HELP  = "help"
)

func NewBot() *telegramBotApi.BotAPI {
	//checkCommits(&commits)
	log.Println("[App] Commits checked")
	bot, err := telegramBotApi.NewBotAPI(env.Environment.TelegramApiToken)
	if err != nil {
		panic(err)
	}
	log.Println("[App] Bot initialized")
	return bot
}
func GetUpdateConfig(offset int, timeout int) telegramBotApi.UpdateConfig {
	updateConfig := telegramBotApi.NewUpdate(offset)
	updateConfig.Timeout = timeout
	return updateConfig
}
func GetUpdates(app *app.Application, config telegramBotApi.UpdateConfig) telegramBotApi.UpdatesChannel {
	return app.TelegramBot.GetUpdatesChan(config)
}

func handleUpdate(sys *app.Application, update telegramBotApi.Update) {
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.IsCommand() {
		replyMessage = handleCommit(sys, update)
	} else {
		replyMessage.Text = update.Message.Text
	}
	log.Printf("Message from %d: %s", update.Message.Chat.ID, replyMessage.Text)
	_, _ = sys.TelegramBot.Send(replyMessage)
}

func handleCommit(sys *app.Application, update telegramBotApi.Update) telegramBotApi.MessageConfig {
	service := dao.NewCommandService(sys.Cache, sys.DB)
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case HELLO:
		replyMessage.Text = "Hello " + update.Message.From.FirstName
	case HELP:
		str, err := service.GetCommandHelperByCommandType(HELP)
		if err != nil {
			replyMessage.Text = err.Error()
		}
		replyMessage.Text = str
	default:
		replyMessage.Text = "No such command!!!"
	}
	return replyMessage
}

func ListenUpdates(sys *app.Application, updates telegramBotApi.UpdatesChannel) {
	for update := range updates {
		go handleUpdate(sys, update)
	}
}
