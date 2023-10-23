package dao

import (
	"gorm.io/gorm"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/constant/commnadType"
	"telegram-bot/telegram-bot-main/dao"
)

type CommandService struct {
	cache      cache.Cacher
	commandDao *dao.CommandDao
	name       string
}

func NewCommandService(cache cache.Cacher, db *gorm.DB) *CommandService {
	return &CommandService{cache: cache, commandDao: dao.NewCommandDao(db), name: "commandService"}
}

func (cs *CommandService) GetCommandHelperByCommandType(enum commnadType.Enum) (string, error) {
	return enum.GetFormat(), nil
}
