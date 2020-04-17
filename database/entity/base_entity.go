package entity

import (
	"time"
)

// BaseModel .
type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// AutoIncrementEntity 基础属性.
type AutoIncrementEntity struct {
	CreatedAt time.Time  `gorm:"column:created_at;" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updated_at;" json:"-"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"-"`
	ID        uint64     `gorm:"primary_key;auto_increment;" json:"id" `
}
