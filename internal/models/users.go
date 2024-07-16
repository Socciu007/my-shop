package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
    ID        string    `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null" json:"username" validate:"required,min=5,max=50"`
    Email     string    `gorm:"unique;not null" json:"email" validate:"required,email"`
    Password  string    `gorm:"not null" json:"password" validate:"required,min=8"`
    Role      string    `gorm:"default:user"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TableName sets the table name for the Users model.
func (Users) TableName() string {
	return "users"
}
