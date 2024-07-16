package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        string    `gorm:"primaryKey"`
	Username  string    `gorm:"unique;not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"default:user"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TableName sets the table name for the Users model.
func (Users) TableName() string {
	return "users"
}
