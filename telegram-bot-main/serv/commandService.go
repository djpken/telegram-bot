package dao

import (
	"strings"
	"telegram-bot/telegram-bot-main/constant"
	"telegram-bot/telegram-bot-main/dao"
)

type commandServ struct{}

var commandDao = &dao.CommandDao{}

func (commandServ) GetCommandHelperByCommandType(commandType string) (string, error) {
	list, err := commandDao.GetCommandByCommandType(commandType)
	if err != nil {
		return "", err
	}
	stringWriter := strings.Builder{}
	for _, h := range list {
		stringWriter.WriteString(h.String() + "\n")

	}
	s := constant.MiddleInCommandType(commandType, stringWriter.String())
	return s, nil
}
