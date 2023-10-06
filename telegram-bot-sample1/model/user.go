package model

import "telegram-bot/telegram-sample1/model/base"

type User struct {
	base.Id
	UserId string `json:"userId" gorm:"DEFAULT:'';not null"`
	Name   string `json:"name" gorm:"DEFAULT:'';not null"`
	base.TimeStamps
	base.DeleteAt
}
