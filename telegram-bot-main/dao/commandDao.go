package dao

import (
	"gorm.io/gorm"
)

type CommandDao struct {
	db   *gorm.DB
	name string
}

func NewCommandDao(db *gorm.DB) *CommandDao {
	return &CommandDao{db: db, name: "commandDao"}
}
