package model

import "telegram-bot/telegram-sample1/model/base"

type Uri struct {
	base.Id
	Path     string `json:"path" gorm:"DEFAULT:'';not null"`
	Interval int    `json:"interval" gorm:"DEFAULT:0;not null"`
	base.TimeStamps
	base.DeleteAt
}
