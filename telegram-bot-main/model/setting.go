package model

import "telegram-bot/telegram-bot-main/model/base"

type Setting struct {
	base.Id
	name  string `db:"type:varchar(255);default:''"`
	value string `db:"type:varchar(255);default:''"`
}
