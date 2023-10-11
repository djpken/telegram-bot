package model

import "telegram-bot/telegram-bot-main/model/base"

type Uri struct {
	base.Model
	Path     string `db:"type:varchar(255);default:''"`
	Interval int    `db:"DEFAULT:0;not null"`
}
