package commit

import (
	"strings"
)

type Command struct {
	Instruction string
	Introduce   string
	Args        []string
}
type HelpText struct {
	Header string
	Body   []Command
	Footer string
}

func Help(help *HelpText, space int) string {
	var builder strings.Builder
	builder.WriteString(help.Header + "\n\n")
	for _, it := range help.Body {
		builder.WriteString("/" + it.Instruction + strings.Repeat(" ", space-len(it.Introduce)) + it.Introduce + "\n")
	}
	builder.WriteString("\n" + help.Footer)
	return builder.String()
}
