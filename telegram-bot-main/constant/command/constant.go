package command

import (
	"fmt"
	"strings"
)

type Enum int

const Nil Enum = -1
const (
	Hello Enum = iota
	Help
	Http
	Todo
	Open
	Close
	List
	Create
	Update
	Delete
)

var list = []string{
	"hello",
	"help",
	"http",
	"todo",
	"open",
	"close",
	"list",
	"create",
	"update",
	"delete",
}

type value struct {
	introduce string
	header    string
	footer    string
	args      map[Enum]string
}

var listMap = map[Enum]value{
	Hello: {
		introduce: "for testing only",
		header:    "Test my telegram API",
		footer:    "Test my telegram API",
	},
	Help: {
		introduce: "list basic command",
		header:    "This command is list command function",
		footer:    "This command is list command function",
	},
	Http: {
		introduce: "list all http command",
		header:    "List http command",
		footer:    "List http command",
		args: map[Enum]string{
			List:   "List all http",
			Create: "Create http",
			Update: "Update http",
			Delete: "Delete http",
		},
	},
	Todo: {
		introduce: "list all todo command",
		header:    "List todo command",
		footer:    "List todo command",
		args: map[Enum]string{
			List:   "List all todo",
			Create: "Create todo",
			Update: "Update todo",
			Delete: "Delete todo",
		},
	},
}

func NewEnum(str string) Enum {
	for i, v := range list {
		if v == str {
			return Enum(i)
		}
	}
	return Enum(-1)
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
func (e Enum) Arg() string {
	if e == -1 {
		return "Nil"
	}
	if int(e) >= len(list) {
		return fmt.Sprintf("Use %d out array length %d", int(e), len(list))
	}
	s := list[e]
	return strings.ToUpper(string(s[0])) + s[1:]

}
func (e Enum) getValue() value {
	return listMap[e]
}

func (e Enum) GetCommandRow() string {
	return "/" + e.Command() + " - " + e.getValue().introduce
}
func (e Enum) GetArgRow(prefix string) string {
	return "/" + prefix + e.Arg() + " - " + NewEnum(prefix).getValue().args[e]
}
func (e Enum) GetFormat() string {
	c := e.getValue()
	if c.args == nil {
		return ""
	}

	var builder strings.Builder
	builder.WriteString(c.header + "\n\n")

	for key := range c.args {
		builder.WriteString(key.GetArgRow(e.Command()) + "\n")
	}
	if len(c.args) > 0 {
		builder.WriteString("\n")
	}
	builder.WriteString(c.footer)
	return builder.String()
}
