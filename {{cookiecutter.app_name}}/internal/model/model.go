package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"column:ID; primaryKey" json:"ID"`
	CreatedAt time.Time      `gorm:"column:CreatedAt" json:"CreatedAt"`
	UpdatedAt time.Time      `gorm:"column:UpdatedAt" json:"UpdatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:DeletedAt; index" json:"-"`
}
