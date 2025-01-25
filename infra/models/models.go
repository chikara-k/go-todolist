package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	*gorm.Model
	// gorm.Model provides fields
	// ID         uint `gorm:"primaryKey"`
	// CreatedAt  time.Time
	// UpdatedAt  time.Time
	// DeletedAt  gorm.DeletedAt `gorm:"index"`
	Title      string
	expireDate *time.Time
	Content    string `json:"content"`
}
