package entity

import(
	"time"

)

// BaseModel .
type BaseModel struct{
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// AutoIncrementEntity 基础属性.
type AutoIncrementEntity struct{
	CreatedAt time.Time  `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"omitempty"`
}
