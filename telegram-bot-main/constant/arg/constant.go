package arg

type Enum int

const (
	HELLO Enum = iota
	HELP
	HTTP
)

var list = []string{
	"hello",
	"help",
	"http",
}

type value struct {
	introduce string
	header    string
	middle    string
	footer    string
}

var listMap = map[Enum]value{
	HELLO: {
		introduce: "Hello",
		header:    "Hello",
		middle:    "Hello",
		footer:    "Hello",
	},
	HELP: {
		introduce: "Help",
		header:    "Help",
		middle:    "Help",
		footer:    "Help",
	},
	HTTP: {
		introduce: "Http",
		header:    "HTTP",
		middle:    "HTTP",
		footer:    "HTTP",
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
