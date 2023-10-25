package dao

import (
	"gorm.io/gorm"
)

type TodoDao struct {
	db   *gorm.DB
	name string
}

func NewTodoDao(db *gorm.DB) *TodoDao {
	return &TodoDao{db: db, name: "todoDao"}
}
func (t *TodoDao) GetAll() string {
	return
}
