package env

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Env struct {
	TelegramApiToken string
	Space            int
}

var Environment Env

func init() {
	_ = godotenv.Load()
	i, _ := strconv.Atoi(os.Getenv("SPACE"))

	Environment = Env{
		TelegramApiToken: os.Getenv("TELEGRAM_API_TOKEN"),
		Space:            i,
	}
}
