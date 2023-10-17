package dao

import (
	"strings"
	"telegram-bot/telegram-bot-main/dao"
)

type commandServ struct {
}

var commandDao = &dao.CommandDao{}

func (commandServ) GetCommandHelperByCommandType(commandTypeName string) (string, error) {
	list, err := commandDao.GetCommandByCommandType(commandTypeName)
	if err != nil {
		return "", err
	}
	stringWriter := strings.Builder{}
	for _, h := range list {
		stringWriter.WriteString(h.String() + "\n")
	}
	return , nil
}
