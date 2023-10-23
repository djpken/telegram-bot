package bot

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/constant/command"
	command2 "telegram-bot/telegram-bot-main/constant/commnadType"
	"telegram-bot/telegram-bot-main/env"
	dao "telegram-bot/telegram-bot-main/serv"
)

func NewBot() *telegramBotApi.BotAPI {
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
func GetUpdates(bot *telegramBotApi.BotAPI, config telegramBotApi.UpdateConfig) telegramBotApi.UpdatesChannel {
	return bot.GetUpdatesChan(config)
}

func handleUpdate(bot *telegramBotApi.BotAPI, cacher cache.Cacher, db *gorm.DB, update telegramBotApi.Update) {
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.IsCommand() {
		replyMessage = handleCommit(cacher, db, update)
	} else {
		replyMessage.Text = update.Message.Text
	}
	log.Printf("Message from %d: %s", update.Message.Chat.ID, replyMessage.Text)
	_, _ = bot.Send(replyMessage)
}

func handleCommit(cacher cache.Cacher, db *gorm.DB, update telegramBotApi.Update) telegramBotApi.MessageConfig {
	service := dao.NewCommandService(cacher, db)
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, "")
	switch command.NewEnum(update.Message.Command()) {
	case command.HELLO:
		handleHello(service, &replyMessage, &update)
	case command.HELP:
		handleHelp(service, &replyMessage, &update)
	case command.HTTP:
		handleHttp(service, &replyMessage, &update)
	default:
		replyMessage.Text = "No such command!!!"
	}
	return replyMessage
}

func ListenUpdates(bot *telegramBotApi.BotAPI, cacher cache.Cacher, db *gorm.DB, updates telegramBotApi.UpdatesChannel) {
	for update := range updates {
		go handleUpdate(bot, cacher, db, update)
	}
}
func handleHello(service *dao.CommandService, replyMessage *telegramBotApi.MessageConfig, update *telegramBotApi.Update) {
	replyMessage.Text = "Hello " + update.Message.From.FirstName
}
func handleHelp(service *dao.CommandService, replyMessage *telegramBotApi.MessageConfig, update *telegramBotApi.Update) {
	str, err := service.GetCommandHelperByCommandType(command2.BASIC)
	if err != nil {
		replyMessage.Text = err.Error()
	} else {
		replyMessage.Text = str
	}
}
func handleHttp(service *dao.CommandService, replyMessage *telegramBotApi.MessageConfig, update *telegramBotApi.Update) {
}
