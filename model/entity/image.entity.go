package entity

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Image      string         `json:"image"`
	CategoryID uint           `json:"category_id"`
	CratedAt   time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index,column:deleted_at"`
}
