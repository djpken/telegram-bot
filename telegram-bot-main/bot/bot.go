package bot

import (
	telegramBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"regexp"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/constant/command"
	"telegram-bot/telegram-bot-main/constant/commnadType"
	"telegram-bot/telegram-bot-main/env"
	serv "telegram-bot/telegram-bot-main/serv"
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
	commandService := serv.NewCommandService(cacher, db)
	todoService := serv.NewTodoService(cacher, db)
	replyMessage := telegramBotApi.NewMessage(update.Message.Chat.ID, "")
	words := camelCaseToWords(update.Message.Command())
	switch command.NewEnum(words[0]) {
	case command.Hello:
		handleHello(&replyMessage, &update)
	case command.Help:
		handleHelp(commandService, &replyMessage, &update, words)
	case command.Http:
		handleHttp(commandService, &replyMessage, &update, words)
	case command.Todo:
		handleTodo(todoService, &replyMessage, &update, words)
	default:
		replyMessage.Text = "No such command!!!"
	}
	return replyMessage
}

func camelCaseToWords(input string) []string {
	re := regexp.MustCompile("[A-Z][a-z]*|[a-z]+")
	return re.FindAllString(input, -1)
}
func ListenUpdates(bot *telegramBotApi.BotAPI, cacher cache.Cacher, db *gorm.DB, updates telegramBotApi.UpdatesChannel) {
	for update := range updates {
		go handleUpdate(bot, cacher, db, update)
	}
}
func handleHello(replyMessage *telegramBotApi.MessageConfig, update *telegramBotApi.Update) {
	replyMessage.Text = "Hello " + update.Message.From.FirstName
}
func handleHelp(service *serv.CommandService, replyMessage *telegramBotApi.MessageConfig, update *telegramBotApi.Update, strings []string) {
	str, err := service.GetCommandHelperByCommandType(commnadType.BASIC)
	if err != nil {
		replyMessage.Text = err.Error()
	} else {
		replyMessage.Text = str
	}
}
func handleHttp(service *serv.CommandService, replyMessage *telegramBotApi.MessageConfig, update *telegramBotApi.Update, strings []string) {
	replyMessage.Text = update.Message.CommandArguments()
}
func handleTodo(todoService *serv.TodoService, replyMessage *telegramBotApi.MessageConfig, update *telegramBotApi.Update, strings []string) {
	todo := command.Todo
	if len(strings) != 2 {
		replyMessage.Text = todo.GetFormat()
	}
	switch command.NewEnum(strings[1]) {
	case command.List:
		replyMessage.Text = todoService.GetAll()
	case command.Create:
	case command.Update:
	case command.Delete:
	}
}
