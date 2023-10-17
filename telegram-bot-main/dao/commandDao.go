package dao

import (
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/model"
)

type CommandDao struct {
}

var db = app.App.DB

func (c *CommandDao) GetCommandByCommandType(commandTypeName string) ([]model.Command, error) {
	var commands []model.Command
	err := db.Model(&commands).Preload("CommandType").Find("command_type.name = ?", commandTypeName).Error
	return commands, err
}
