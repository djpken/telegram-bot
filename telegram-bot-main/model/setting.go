package model

import "telegram-bot/telegram-bot-main/model/base"

type Setting struct {
	base.Id
	Name  string `db:"type:varchar(255);default:''"`
	Value string `db:"type:varchar(255);default:''"`
}
