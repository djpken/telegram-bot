package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type Env struct {
	TelegramApiToken string
	Space            int
	Dsn              string
	Schema           string
}

var Environment = new(Env)

func init() {
	_ = godotenv.Load()
	i, _ := strconv.Atoi(os.Getenv("SPACE"))
	var builder strings.Builder
	properties := os.Getenv("DATABASE_PROPERTIES")
	builder.WriteString("host=" + os.Getenv("DATABASE_HOST") + " ")
	builder.WriteString("user=" + os.Getenv("DATABASE_USERNAME") + " ")
	builder.WriteString("password=" + os.Getenv("DATABASE_PASSWORD") + " ")
	builder.WriteString("port=" + os.Getenv("DATABASE_PORT") + " ")
	builder.WriteString("dbname=" + os.Getenv("DATABASE_NAME") + " ")
	builder.WriteString(properties)
	Environment = &Env{
		TelegramApiToken: os.Getenv("TELEGRAM_API_TOKEN"),
		Space:            i,
		Dsn:              builder.String(),
		Schema:           os.Getenv("DATABASE_SCHEMA"),
	}
	log.Println("[App] Environment initialized")
}
