//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"
	"gorm.io/gorm"
	"telegram-bot/telegram-bot-main/app"
	"telegram-bot/telegram-bot-main/bot"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/db"
	"telegram-bot/telegram-bot-main/env"
)

var envSet = wire.NewSet(
	providerEnv,
)
var botSet = wire.NewSet(
	providerBot,
)
var dbSet = wire.NewSet(
	providerDB,
)
var cacheSet = wire.NewSet(
	context.Background,
	providerCache,
)
var applicationSet = wire.NewSet(
	providerApplication,
)

func providerEnv() *env.Env {
	return env.NewEnv()
}

func providerBot(env *env.Env) *tgBotApi.BotAPI {
	return bot.NewBot(env)
}
func providerDB(env *env.Env) *gorm.DB {
	return db.NewDataBase(env)
}
func providerCache(env *env.Env, ctx context.Context) cache.Cacher {
	return cache.NewRedisCache(env, ctx)
}
func providerApplication(api *tgBotApi.BotAPI, db *gorm.DB, cache cache.Cacher) *app.Application {
	return app.NewApplication(api, db, cache)
}
func InitializeApp() (*app.Application, error) {
	panic(wire.Build(envSet, botSet, dbSet, cacheSet, applicationSet))
	return nil, nil
}
