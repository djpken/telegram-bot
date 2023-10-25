package model

import "telegram-bot/telegram-bot-main/model/base"

type Todo struct {
	base.Model
	Text     string `db:"type:varchar(255);default:''"`
	Interval int    `db:"DEFAULT:0;not null"`
}

type Enum int

const (
	FiveMinutes Enum = iota
	FifteenMinutes
	ThirtyMinutes
	OneHour
	FourHours
	OneDay
)

var list = []int{
	300,
	900,
	1800,
	3600,
	14400,
	86400,
}

func NewEnum(str int) Enum {
	for i, v := range list {
		if v == str {
			return Enum(i)
		}
	}
	return Enum(-1) // Return an invalid enum value to indicate not found.
}
func (e Enum) Int() int {
	if e < 0 || int(e) >= len(list) {
		return -1
	}
	return list[e]
}
