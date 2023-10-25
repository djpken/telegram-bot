package dao

import (
	"gorm.io/gorm"
	"telegram-bot/telegram-bot-main/cache"
	"telegram-bot/telegram-bot-main/dao"
)

type TodoService struct {
	cache   cache.Cacher
	todoDao *dao.TodoDao
	name    string
}

func NewTodoService(cache cache.Cacher, db *gorm.DB) *TodoService {
	return &TodoService{cache: cache, commandDao: dao.NewTodoDao(db), name: "todoService"}
}

func (cs *TodoService) GetAll() string {
	cs.todoDao.GetAll()

}
