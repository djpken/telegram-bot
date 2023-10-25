package commnadType

import (
	"strings"
	"telegram-bot/telegram-bot-main/constant/command"
)

type Enum int

const (
	BASIC Enum = iota
)

var list = []string{
	"basic",
}

type value struct {
	header   string
	footer   string
	commands []command.Enum
}

var listMap = map[Enum]value{
	BASIC: {
		header: "Basic command list",
		footer: "/help <command> show help",
		commands: []command.Enum{
			command.Hello,
			command.Help,
			command.Http,
			command.Todo,
		},
	},
}

func NewEnum(str string) Enum {
	for i, v := range list {
		if v == str {
			return Enum(i)
		}
	}
	return Enum(-1) // Return an invalid enum value to indicate not found.
}
func (e Enum) String() string {
	if e < 0 || int(e) >= len(list) {
		return "UNKNOWN"
	}
	return list[e]
}
func (e Enum) getValue() value {
	return listMap[e]
}

func (e Enum) GetFormat() string {
	commands := e.getValue().commands
	stringWriter := strings.Builder{}
	for _, h := range commands {
		stringWriter.WriteString(h.GetRow() + "\n")
	}
	return e.getValue().format(stringWriter.String())
}
func (v value) format(middle string) string {
	var builder strings.Builder
	builder.WriteString(v.header + "\n\n")
	builder.WriteString(middle + "\n")
	builder.WriteString(v.footer)
	return builder.String()
}
