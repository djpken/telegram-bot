package commnadType

import (
	"fmt"
	"strings"
	"telegram-bot/telegram-bot-main/constant/command"
)

type Enum int

const Nil Enum = -1
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
		footer: "Basic command list",
		commands: []command.Enum{
			command.Hello,
			command.Help,
			command.Http,
			command.Todo,
			command.Open,
			command.InlineOpen,
			command.Close,
			command.InlineClose,
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
func (e Enum) Command() string {
	if e == -1 {
		return "Nil"
	}
	if int(e) >= len(list) {
		return fmt.Sprintf("Use %d out array length %d", int(e), len(list))
	}
	return list[e]
}
func (e Enum) getValue() value {
	return listMap[e]
}

func (e Enum) GetFormat() string {
	v := e.getValue()
	commands := v.commands
	stringWriter := strings.Builder{}

	for _, h := range commands {
		stringWriter.WriteString(h.GetCommandRow() + "\n")
	}

	middle := stringWriter.String()

	var builder strings.Builder
	builder.WriteString(v.header + "\n\n")
	builder.WriteString(middle)
	if len(middle) > 0 {
		builder.WriteString("\n")
	}
	builder.WriteString(v.footer)

	return builder.String()
}
