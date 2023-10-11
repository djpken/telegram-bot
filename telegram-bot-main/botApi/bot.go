package botApi

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/command"
	"telegram-bot/telegram-bot-main/env"
)

func init() {
	checkCommits(&commits)
	log.Println("[App] Commits checked")
	bot, err := telegramBotApi.NewBotAPI(env.Environment.TelegramApiToken)
	if err != nil {
		panic(err)
	}
	app.App.TelegramBot = bot
	log.Println("[App] Bot initialized")
}
func GetUpdateConfig(offset int, timeout int) telegramBotApi.UpdateConfig {
	updateConfig := telegramBotApi.NewUpdate(offset)
	updateConfig.Timeout = timeout
	return updateConfig
}
func GetUpdates(telegramBot *telegramBotApi.BotAPI, config telegramBotApi.UpdateConfig) telegramBotApi.UpdatesChannel {
	return telegramBot.GetUpdatesChan(config)
}

func handleUpdate(bot *telegramBotApi.BotAPI, update telegramBotApi.Update) {
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.IsCommand() {
		replyMessage = handleCommit(update)
	} else {
		replyMessage.Text = update.Message.Text
	}
	log.Printf("Message from %d: %s", update.Message.Chat.ID, replyMessage.Text)
	_, _ = bot.Send(replyMessage)
}
func handleCommit(update telegramBotApi.Update) telegramBotApi.MessageConfig {
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case commits[0].Instruction:
		replyMessage.Text = command.Help(&helpText, env.Environment.Space)
	case commits[1].Instruction:
		replyMessage.Text = "Hello " + update.Message.From.FirstName
	default:
		replyMessage.Text = "No such command!!!"
	}
	return replyMessage
}

func checkCommits(commits *[]command.Command) {
	for _, c := range *commits {
		if len(c.Instruction) > env.Environment.Space-1 {
			panic("Instruction is too long")
		}
	}
}
func ListenUpdates(updates telegramBotApi.UpdatesChannel) {
	for update := range updates {
		go handleUpdate(app.App.TelegramBot, update)
	}
}

var commits = []command.Command{
	{
		"help",
		"Help menu list all commands",
		nil,
	},
	{
		"start",
		"Hello world",
		nil,
	},
}
var helpText = command.HelpText{
	Header: "This bot commands are:",
	Body:   commits,
	Footer: `Use /help <command> get more information.`,
}
