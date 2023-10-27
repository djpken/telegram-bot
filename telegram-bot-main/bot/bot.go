package bot

import (
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/env"
)

func NewBot(env *env.Env) *tgBotApi.BotAPI {
	bot, err := tgBotApi.NewBotAPI(env.TelegramApiToken)
	if err != nil {
		panic(err)
	}
	log.Println("[App] Bot initialized")
	return bot
}
func GetUpdateConfig(offset int, timeout int) tgBotApi.UpdateConfig {
	updateConfig := tgBotApi.NewUpdate(offset)
	updateConfig.Timeout = timeout
	return updateConfig
}
func GetUpdates(bot *tgBotApi.BotAPI, config tgBotApi.UpdateConfig) tgBotApi.UpdatesChannel {
	return bot.GetUpdatesChan(config)
}
func ListenUpdates(bot *tgBotApi.BotAPI, cacher cache.Cacher, db *gorm.DB, updates tgBotApi.UpdatesChannel) {
	for update := range updates {
		go handleUpdate(bot, cacher, db, update)
	}
}

func handleUpdate(bot *tgBotApi.BotAPI, cacher cache.Cacher, db *gorm.DB, update tgBotApi.Update) {
	if update.Message != nil {
		handleMessage(bot, cacher, db, update)
	}
	if update.CallbackQuery != nil {
		handleCallBackQuery(bot, cacher, db, update)
	}
	if update.InlineQuery != nil {
		handleInlineQuery(bot, cacher, db, update)
	}
}
