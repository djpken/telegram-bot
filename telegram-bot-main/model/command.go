package model

import (
	"telegram-bot/telegram-bot-main/model/base"
)

type Arg struct {
	base.Id
	Name      string `db:"type:varchar(255);default:''"`
	Introduce string `db:"type:varchar(255);default:''"`
	CommandId uint64
}

func (a Arg) String() string {
	//TODO WIP
	return ""
}
