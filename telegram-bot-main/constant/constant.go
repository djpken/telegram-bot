package constant

import "strings"

type CommandTypeEnum int

const (
	BASIC CommandTypeEnum = iota
)

var commandTypeEnum = [...]string{
	"basic",
}

func (c CommandTypeEnum) String() string {
	return commandTypeEnum[c]
}

type CommandTypeHelper struct {
	Header string
	Footer string
}

var commandTypeMap = map[string]CommandTypeHelper{
	BASIC.String(): {
		Header: "Basic command list",
		Footer: "/help <command> show help",
	},
}

func MiddleInCommandType(commandType string, middle string) string {
	return commandTypeMap[commandType].string(middle)
}

func (c CommandTypeHelper) string(middle string) string {
	var builder strings.Builder
	builder.WriteString(c.Header + "\n")
	builder.WriteString("\n")
	builder.WriteString(middle + "\n")
	builder.WriteString("\n")
	builder.WriteString(c.Footer + "\n")
	return builder.String()
}
