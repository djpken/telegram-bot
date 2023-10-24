package command

import (
	"strings"
)

type Enum int

const (
	Hello Enum = iota
	Help
	Http
	Todo
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
		header:    "Hello",
		footer:    "Hello",
	},
	Help: {
		introduce: "list basic command",
		header:    "Help",
		footer:    "Help",
	},
	Http: {
		introduce: "list all http command",
		header:    "Http",
		footer:    "Http",
	},
	Todo: {
		introduce: "list all todo command",
		header:    "todo",
		footer:    "todo",
		args: map[Enum]string{
			List:   "list all todo",
			Create: "create todo",
			Update: "update todo",
			Delete: "delete todo",
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

func (e Enum) GetRow() string {
	return "/" + e.String() + " - " + e.getValue().introduce
}
func (e Enum) GetFormat() string {
	c := e.getValue()

	var builder strings.Builder
	builder.WriteString(c.header + "\n\n")

	for key, value := range c.args {
		builder.WriteString(key.String() + " " + value + "\n")
	}
	if len(c.args) > 0 {
		builder.WriteString("\n")
	}
	builder.WriteString(c.footer)
	return builder.String()
}
