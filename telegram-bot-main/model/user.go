package model

import "telegram-bot/telegram-bot-main/model/base"

type User struct {
	base.Model
	Name  string `db:"type:varchar(255);default:''"`
	Roles []Role
}
type Role struct {
	base.Model
	Name   string `db:"type:varchar(255);default:''"`
	Menus  []Menu
	UserId uint64
}
type Menu struct {
	base.Model
	Name      string `db:"type:varchar(255);default:''"`
	CommandId uint64
	RoleId    uint64
}
