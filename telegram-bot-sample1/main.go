package main

import (
	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"varmeego-telegram-bot/varmeego-telegram-sample1/commit"
)

const (
	varmeegoBotToken = "6599882879:AAH_Sv9BnYEBN9SmqX-8EqUDBdaZk_8_ZmI"
	space            = 9
)

var commits = []commit.Command{
	{
		"help",
		"Help",
		nil,
	},
	{
		"start",
		"Hello",
		nil,
	},
}
var helpText = commit.HelpText{
	Header: "This bot commands are:",
	Body:   commits,
	Footer: `"/help <command>" get more information.`,
}

func main() {
	checkCommits(&commits)
	envToken, isFull := os.LookupEnv("TELEGRAM_API_TOKEN")
	if !isFull {
		envToken = varmeegoBotToken
	}
	bot := newBot(envToken)
	updateConfig := botApi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		go handleUpdate(bot, update)
	}
}
func handleUpdate(bot *botApi.BotAPI, update botApi.Update) {
	replyMessage := botApi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.IsCommand() {
		replyMessage = handleCommit(update)
	} else {
		replyMessage.Text = update.Message.Text
	}
	log.Printf("Message from %d: %s", update.Message.Chat.ID, replyMessage.Text)
	_, _ = bot.Send(replyMessage)
}
func newBot(token string) *botApi.BotAPI {
	bot, err := botApi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	return bot
}
func handleCommit(update botApi.Update) botApi.MessageConfig {
	replyMessage := botApi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case commits[0].Instruction:
		replyMessage.Text = commit.Help(&helpText, space)
	case commits[1].Instruction:
		replyMessage.Text = "Hello " + update.Message.From.FirstName
	default:
		replyMessage.Text = "No such command!!!"
	}
	return replyMessage
}
func checkCommits(commits *[]commit.Command) {
	for _, c := range *commits {
		if len(c.Instruction) > space-1 {
			panic("Introduce is too long")
		}
	}
}
