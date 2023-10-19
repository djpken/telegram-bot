package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"telegram-bot/telegram-bot-main/env"
	"telegram-bot/telegram-bot-main/model"
)

func NewDataBase() *gorm.DB {
	config := env.Environment.DB
	postgresConfig := postgres.Config{
		DSN: config.Dsn,
	}
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Schema + ".",
			SingularTable: true},
	}
	db, err := gorm.Open(postgres.New(postgresConfig), gormConfig)
	if err != nil {
		panic(err)
	}
	log.Println("[App] Database connected")
	err = db.AutoMigrate(
		&model.Command{},
		&model.CommandType{},
		&model.Arg{},
		&model.Uri{},
		&model.User{},
		&model.Role{},
		&model.Menu{},
		&model.Setting{},
	)
	if err != nil {
		panic(err)
	}
	log.Println("[App] Database autoMigrated")
	return db
}
