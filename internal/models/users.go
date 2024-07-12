package models

type Users struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:user"`
}

// TableName sets the table name for the Users model.
func (Users) TableName() string {
	return "users"
}
