package dao

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"strings"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/constant"
	"telegram-bot/telegram-bot-main/dao"
)

type CommandService struct {
	cache      cache.Cacher
	commandDao *dao.CommandDao
	name       string
}

func NewCommandService(cache cache.Cacher, db *gorm.DB) *CommandService {
	commandDao := dao.NewCommandDao(db)
	return &CommandService{cache: cache, commandDao: commandDao, name: "commandService"}
}

func (cs *CommandService) GetCommandHelperByCommandType(commandType string) (string, error) {
	key := cs.name + ":" + commandType

	result, err := cs.cache.Get(key)
	if err == nil {
		s := constant.MiddleInCommandType(commandType, result)
		return s, nil
	}
	if err != redis.Nil {
		log.Println(err)
		return "", err
	}

	//Handle cache miss
	stringWriter := strings.Builder{}
	list, err := cs.commandDao.GetCommandByCommandType(commandType)
	if err != nil {
		return "", err
	}
	for _, h := range list {
		stringWriter.WriteString(h.String() + "\n")
	}
	err = cs.cache.Set(key, stringWriter.String(), 0)
	stringWriter.WriteString(result)
	s := constant.MiddleInCommandType(commandType, stringWriter.String())
	return s, nil
}
