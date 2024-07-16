package user

import (
	"fmt"
	"my_shop/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
    return UserService{db: db}
}

// logic to get a list of all users
func (us *UserService) GetAllUsers() ([]models.Users, error) {
    var users []models.Users
    if err := us.db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

// logic to create a new user
func (us *UserService) CreateUser(user *models.Users) (string, error) {
    var existUser models.Users

    //check exist user already email in database
    if err := us.db.Where("email =?", user.Email).First(&existUser).Error; err == nil {
        return "BadRequest", fmt.Errorf("user already exists")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) 
    if err != nil {
        return "BadRequest", err
    }
    user.Password = string(hashedPassword)
    user.ID = uuid.New().String()

    return "InternalServerError", us.db.Create(user).Error
}

// logic to update a user 
func (us *UserService) UpdateUser(id string, updateFields models.Users) (*models.Users, status, error) {
    var user models.Users

    // Find user in database by ID
    if err := us.db.Where("id = ?", id).First(&user).Error; err != nil {
        return &user, "BadRequest", err
    }

    // Update information user from map updateFields
    if err := us.db.Model(&user).Updates(updateFields).Error; err != nil {
        return &user, "InternalServerError", err
    }

    return &user, "OK", nil
}


// logic to delete a user
func (us *UserService) DeleteUser(id string) (string, error) {
    var user models.Users
    //check if user exists
    if err := us.db.Where("id =?", id).First(&user).Error; err!= nil {
        return "BadRequest", err
    }

    //Delete user from database
    if err := us.db.Where("id =?", id).Delete(&user).Error; err!= nil {
        return "InternalServerError", err
    }

    return "OK", nil
}

// Other methods like CreateUser, GetUserByID, UpdateUser, DeleteUser can be defined here