package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primaryKey" json:"ID"`
	CreatedAt time.Time      `json:"CreatedAT"`
	UpdatedAt time.Time      `json:"UpdatedAT"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
