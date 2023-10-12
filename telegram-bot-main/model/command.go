package model

import (
	"strings"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/env"
	"telegram-bot/telegram-bot-main/model/base"
)

type Helper struct {
	base.Id
	Header   string `db:"type:varchar(255);default:''"`
	Footer   string `db:"type:varchar(255);default:''"`
	Commands []Command
}
type Command struct {
	base.Id
	Name      string `db:"type:varchar(255);default:''"`
	Introduce string `db:"type:varchar(255);default:''"`
	Args      []Arg
	HelperId  uint64
	Menus     []Menu
}
type Arg struct {
	base.Id
	Name      string `db:"type:varchar(255);default:''"`
	Introduce string `db:"type:varchar(255);default:''"`
	CommandId uint64
}

const space = 9

func (c *Command) String() string {
	var builder strings.Builder
	builder.WriteString("/" + c.Name + " - " + c.Introduce + "\n")
	return builder.String()
}
func (h *Helper) String() string {
	var builder strings.Builder
	builder.WriteString(h.Header + "\n\n")
	for _, it := range h.Commands {
		builder.WriteString(it.String())
	}
	builder.WriteString("\n" + h.Footer)
	return builder.String()
}
func (h *Helper) GetHelperById(i int) error {	c.Name

	return app.App.DB.Model(&Helper{}).Pre"id = ?", i).Find(&h).Error
}
func (a *Arg) String() string {
	//TODO
	return ""
}
func (c *Command) Check() error {
	if len(c.Name) >= env.Environment.Space {

	}
}
func checkCommits(commits *[]command.Command) {
	for _, c := range *commits {
		if len(c.Instruction) > env.Environment.Space-1 {
			panic("Instruction is too long")
		}
	}
}
