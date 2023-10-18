package dao

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"strings"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/constant"
)

type CommandService struct {
	Cache app.Cacher
	Ctx   context.Context
}

func NewCommandService(cache app.Cacher, ctx context.Context) *CommandService {
	return &CommandService{Cache: cache, Ctx: ctx}
}

const serviceName = "commandService"

func (cs *CommandService) GetCommandHelperByCommandType(commandType string) (string, error) {
	stringWriter := strings.Builder{}
	key := serviceName + ":" + commandType

	result, err := cs.Cache.Get(cs.Ctx, key).Result()
	if err == redis.Nil { // cache miss
		list, err := commandDao.GetCommandByCommandType(commandType)
		if err != nil {
			return "", err
		}
		for _, h := range list {
			stringWriter.WriteString(h.String() + "\n")
		}
		cs.Cache.Set(cs.Ctx, key, stringWriter.String(), 0)
	} else if err != nil { // other cache errors
		log.Println(err)
		return "", err
	} else {
		stringWriter.WriteString(result)
	}

	s := constant.MiddleInCommandType(commandType, stringWriter.String())
	return s, nil
}
