package dao

import (
	"gorm.io/gorm"
	"telegram-bot/telegram-bot-main/model"
)

type CommandDao struct {
	db   *gorm.DB
	name string
}

func NewCommandDao(db *gorm.DB) *CommandDao {
	return &CommandDao{db: db, name: "commandDao"}
}

func (c *CommandDao) GetCommandByCommandType(commandTypeName string) ([]model.Command, error) {
	var commands []model.Command
	err := c.db.Preload("commandType").Find("command_type.name = ?", commandTypeName).Error
	return commands, err
}
