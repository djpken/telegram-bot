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
	var dsn strings.Builder
	telegramApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	i, _ := strconv.Atoi(os.Getenv("SPACE"))
	host := isEmptyElse("DOCKER_COMPOSE_POSTGRES_HOST", "DATABASE_HOST")
	port := isEmptyElse("DOCKER_COMPOSE_POSTGRES_PORT", "DATABASE_PORT")
	database := isEmptyElse("DOCKER_COMPOSE_POSTGRES_DATABASE", "DATABASE_DATABASE")
	schema := isEmptyElse("DOCKER_COMPOSE_POSTGRES_SCHEMA", "DATABASE_SCHEMA")
	user := isEmptyElse("DOCKER_COMPOSE_POSTGRES_USER", "DATABASE_USER")
	password := isEmptyElse("DOCKER_COMPOSE_POSTGRES_PASSWORD", "DATABASE_PASSWORD")
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
	log.Println("[App] Environment initialized")
}
func isEmptyElse(dockerComposeEnv string, localEnv string) string {
	env := os.Getenv(dockerComposeEnv)
	if env == "" {
		return os.Getenv(localEnv)
	}
	return env
}
