package service

import (
	"my_shop/internal/models"

	"gorm.io/gorm"
)

var db *gorm.DB

// GetAllUsers trả về danh sách tất cả người dùng
func GetAllUsers() ([]models.Users, error) {
    var users []models.Users
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

// Other methods like CreateUser, GetUserByID, UpdateUser, DeleteUser can be defined here
