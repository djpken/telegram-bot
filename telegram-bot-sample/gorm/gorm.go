package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"telegram-bot/telegram-sample/app"
	"telegram-bot/telegram-sample/env"
	"telegram-bot/telegram-sample/model"
)

func init() {
	postgresConfig := postgres.Config{
		DSN: env.Environment.Dsn,
	}
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   env.Environment.Schema + ".",
			SingularTable: true},
	}
	db, err := gorm.Open(postgres.New(postgresConfig), gormConfig)
	if err != nil {
		panic(err)
	}
	app.App.Gorm = db
	log.Println("[App] Database connected")
	err = db.AutoMigrate(&model.User{}, &model.Uri{})
	if err != nil {
		panic(err)
	}
	log.Println("[App] Database autoMigrated")
}
