package model

import (
	"strings"
	"telegram-bot/telegram-bot-main/model/base"
)

type Command struct {
	base.Id
	Name        string `db:"type:varchar(255);default:''"`
	Introduce   string `db:"type:varchar(255);default:''"`
	Args        []Arg
	CommandType CommandType
	Helper      Helper
}

type CommandType struct {
	base.Id
	Name      string `db:"type:varchar(255);default:''"`
	CommandId uint64
}

type Helper struct {
	base.Id
	Header    string `db:"type:varchar(255);default:''"`
	Footer    string `db:"type:varchar(255);default:''"`
	CommandId uint64
}
type Arg struct {
	base.Id
	Name      string `db:"type:varchar(255);default:''"`
	Introduce string `db:"type:varchar(255);default:''"`
	CommandId uint64
}

func (c *Command) String() string {
	var builder strings.Builder
	builder.WriteString("/" + c.Name + " - " + c.Introduce + "\n")
	return builder.String()
}
func (c *Helper) String(middle string) string {
	var builder strings.Builder
	builder.WriteString(c.Header + "\n")
	builder.WriteString("\n")
	builder.WriteString(middle + "\n")
	builder.WriteString("\n")
	builder.WriteString(c.Footer + "\n")
	return builder.String()
}
func (a *Arg) String() string {
	//TODO WIP
	return ""
}
