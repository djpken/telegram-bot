package bot

import (
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"regexp"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/constant/command"
	"telegram-bot/telegram-bot-main/constant/commnadType"
	serv "telegram-bot/telegram-bot-main/serv"
)

func handleMessage(bot *tgBotApi.BotAPI, cacher cache.Cacher, db *gorm.DB, update tgBotApi.Update) {
	if update.Message.IsCommand() {
		replyMessage := handleCommand(cacher, db, update)
		if _, err := bot.Send(replyMessage); err != nil {
			panic(err)
		}
	} else {
		replyMessage := handleReplyMsg(cacher, db, update)
		if _, err := bot.Send(replyMessage); err != nil {
			panic(err)
		}
	}
}
func handleCallBackQuery(bot *tgBotApi.BotAPI, cacher cache.Cacher, db *gorm.DB, update tgBotApi.Update) {
	callback := tgBotApi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	// And finally, send a message containing the data received.
	msg := tgBotApi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
func handleInlineQuery(bot *tgBotApi.BotAPI, cacher cache.Cacher, db *gorm.DB, update tgBotApi.Update) {

}

func handleCommand(cacher cache.Cacher, db *gorm.DB, update tgBotApi.Update) tgBotApi.MessageConfig {
	commandService := serv.NewCommandService(cacher, db)
	todoService := serv.NewTodoService(cacher, db)
	replyMessage := tgBotApi.NewMessage(update.Message.Chat.ID, "")
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
	case command.Open:
		handleOpen(todoService, &replyMessage, &update, words)
	case command.Close:
		handleClose(&replyMessage, &update, words)
	default:
		replyMessage.Text = "No such command!!!"
	}
	return replyMessage
}

func handleClose(m *tgBotApi.MessageConfig, t *tgBotApi.Update, words []string) {
	m.ReplyMarkup = tgBotApi.ReplyKeyboardRemove{
		RemoveKeyboard: true,
	}
	m.Text = "Close"
}

var inlineKeyboard = tgBotApi.NewInlineKeyboardMarkup(
	tgBotApi.NewInlineKeyboardRow(
		tgBotApi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
		tgBotApi.NewInlineKeyboardButtonData("2", "2"),
		tgBotApi.NewInlineKeyboardButtonSwitch("3", "3"),
		tgBotApi.InlineKeyboardButton{
			Text: "switch_inline_query_current_chat(1,4)",
		},
	),
	tgBotApi.NewInlineKeyboardRow(
		tgBotApi.NewInlineKeyboardButtonData("4", "4"),
		tgBotApi.NewInlineKeyboardButtonData("5", "5"),
		tgBotApi.NewInlineKeyboardButtonSwitch("6", "6"),
		tgBotApi.InlineKeyboardButton{
			Text: "switch_inline_query_current_chat(1,4)",
		},
	),
)
var keyboard = tgBotApi.NewReplyKeyboard(
	tgBotApi.NewKeyboardButtonRow(
		tgBotApi.NewKeyboardButton("1"),
		tgBotApi.NewKeyboardButtonLocation("2"),
		tgBotApi.NewKeyboardButtonContact("3"),
	),
	tgBotApi.NewKeyboardButtonRow(
		tgBotApi.NewKeyboardButton("4"),
		tgBotApi.NewKeyboardButtonLocation("5"),
		tgBotApi.NewKeyboardButtonContact("6"),
	),
)

func handleReplyMsg(cacher cache.Cacher, db *gorm.DB, update tgBotApi.Update) tgBotApi.MessageConfig {
	replyMessage := tgBotApi.NewMessage(update.Message.Chat.ID, "")
	log.Printf("Message from %d: %s", update.Message.Chat.ID, replyMessage.Text)
	switch update.Message.Text {
	case "open":
		replyMessage.ReplyMarkup = inlineKeyboard
		replyMessage.Text = "Open"
	default:
		replyMessage.Text = update.Message.Text
	}
	return replyMessage
}

func handleOpen(service *serv.TodoService, m *tgBotApi.MessageConfig, t *tgBotApi.Update, words []string) {
	m.ReplyMarkup = keyboard
	m.Text = "Test open key board"
	return
}

func camelCaseToWords(input string) []string {
	re := regexp.MustCompile("[A-Z][a-z]*|[a-z]+")
	return re.FindAllString(input, -1)
}
func handleHello(replyMessage *tgBotApi.MessageConfig, update *tgBotApi.Update) {
	replyMessage.Text = "Hello " + update.Message.From.FirstName
}
func handleHelp(service *serv.CommandService, replyMessage *tgBotApi.MessageConfig, update *tgBotApi.Update, strings []string) {
	str, err := service.GetCommandHelperByCommandType(commnadType.BASIC)
	if err != nil {
		replyMessage.Text = err.Error()
	} else {
		replyMessage.Text = str
	}
}
func handleHttp(service *serv.CommandService, replyMessage *tgBotApi.MessageConfig, update *tgBotApi.Update, strings []string) {
	if len(strings) == 1 {
		replyMessage.Text = command.Http.GetFormat()
		return
	}
}
func handleTodo(todoService *serv.TodoService, replyMessage *tgBotApi.MessageConfig, update *tgBotApi.Update, strings []string) {
	if len(strings) == 1 {
		replyMessage.Text = command.Todo.GetFormat()
		return
	}
	switch command.NewEnum(strings[1]) {
	case command.List:
		replyMessage.Text = todoService.GetAll()
	case command.Create:
	case command.Update:
	case command.Delete:
	}
}
