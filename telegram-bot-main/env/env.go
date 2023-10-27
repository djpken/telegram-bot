package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

type DB struct {
	Dsn    string
	Schema string
}
type Cache struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type Env struct {
	TelegramApiToken string
	Space            int
	DB               DB
	Cache            Cache
}

func NewEnv() *Env {
	mode := getMode()
	telegramApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	i, _ := strconv.Atoi(os.Getenv("SPACE"))
	log.Printf("[App] Environment %s initialized\n", mode)
	return &Env{
		TelegramApiToken: telegramApiToken,
		Space:            i,
		DB:               *getDB(),
		Cache:            *getCache(),
	}
}
func getMode() string {
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
	return mode
}
func getDB() *DB {
	var dsn strings.Builder
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")
	schema := os.Getenv("DATABASE_SCHEMA")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	properties := os.Getenv("DATABASE_PROPERTIES")
	dsn.WriteString("host=" + host + " ")
	dsn.WriteString("port=" + port + " ")
	dsn.WriteString("user=" + user + " ")
	dsn.WriteString("password=" + password + " ")
	dsn.WriteString("dbname=" + databaseName + " ")
	dsn.WriteString(properties)
	config := &DB{
		Dsn:    dsn.String(),
		Schema: schema,
	}
	return config
}
func getCache() *Cache {
	database, err := strconv.Atoi(os.Getenv("CACHE_DATABASE"))
	if err != nil {
		log.Fatal(err)
	}
	config := &Cache{
		Host:     os.Getenv("CACHE_HOST"),
		Port:     os.Getenv("CACHE_PORT"),
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       database,
	}
	return config
}
