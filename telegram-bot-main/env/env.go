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
	mode := os.Getenv("TELEGRAM_BOT_MODE")
	if mode == "test" {
		_ = godotenv.Load(".env.test")
	}
	if mode == "prod" {
		_ = godotenv.Load(".env.prod")
	}
	if mode != "test" && mode != "prod" {
		err := godotenv.Load(".env.dev")
		if err != nil {
			log.Fatal("Error loading .env.dev file")
		}
		mode = "dev"
	}
	var dsn strings.Builder
	telegramApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	i, _ := strconv.Atoi(os.Getenv("SPACE"))
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	database := os.Getenv("DATABASE_DATABASE")
	schema := os.Getenv("DATABASE_SCHEMA")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	properties := os.Getenv("DATABASE_PROPERTIES")
	dsn.WriteString("host=" + host + " ")
	dsn.WriteString("port=" + port + " ")
	dsn.WriteString("user=" + user + " ")
	dsn.WriteString("password=" + password + " ")
	dsn.WriteString("dbname=" + database + " ")
	dsn.WriteString(properties)
	Environment = &Env{
		TelegramApiToken: telegramApiToken,
		Space:            i,
		Dsn:              dsn.String(),
		Schema:           schema,
	}
	log.Printf("[App] Environment %s initialized\n", mode)
}
