package base

import (
	"gorm.io/gorm"
	"time"
)

type Id struct {
	Id uint64 `json:"id" db:"primary_key"`
}
type TimeStamps struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type DeleteAt struct {
	DeletedAt gorm.DeletedAt `json:"deletedAt" db:"index"`
}
type Model struct {
	Id
	TimeStamps
	DeleteAt
}
