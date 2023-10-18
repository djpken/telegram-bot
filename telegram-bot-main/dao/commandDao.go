package dao

import (
	"gorm.io/gorm"
	"telegram-bot/telegram-bot-main/model"
)

type CommandDao struct {
	DB *gorm.DB
}

func NewCommandDao(db *gorm.DB) *CommandDao {
	return &CommandDao{DB: db}
}

func (c *CommandDao) GetCommandByCommandType(commandTypeName string) ([]model.Command, error) {
	var commands []model.Command
	err := c.DB.Preload("CommandType").Find("command_type.name = ?", commandTypeName).Error
	return commands, err
}
