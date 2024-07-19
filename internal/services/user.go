package services

import (
	"fmt"
	"my_shop/global"
	"my_shop/internal/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
func (us *UserService) CreateUser(user *models.Users) (int, error) {
    var existUser models.Users

    //check exist user already email in database
    if err := us.db.Where("email =?", user.Email).First(&existUser).Error; err == nil {
        return http.StatusBadRequest, fmt.Errorf("user already exists")
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost) 
    if err != nil {
        return http.StatusBadRequest, err
    }
    user.Password = string(hashedPassword)
    user.ID = uuid.New().String()

    return http.StatusInternalServerError, us.db.Create(user).Error
}

// logic to update a user 
func (us *UserService) UpdateUser(id string, updateFields models.Users) (models.Users, int, error) {
    var user models.Users

    // Find user in database by ID
    if err := us.db.Where("id = ?", id).First(&user).Error; err != nil {
        return user, http.StatusBadRequest, err
    }

    // Update information user from map updateFields
    if err := us.db.Model(&user).Updates(updateFields).Error; err != nil {
        return user, http.StatusInternalServerError, err
    }

    return user, http.StatusOK, nil
}


// logic to delete a user
func (us *UserService) DeleteUser(id string) (int, error) {
    var user models.Users
    //check if user exists
    if err := us.db.Where("id =?", id).First(&user).Error; err!= nil {
        return http.StatusBadRequest, err
    }

    //Delete user from database
    if err := us.db.Where("id =?", id).Delete(&user).Error; err!= nil {
        return http.StatusInternalServerError, err
    }

    return http.StatusOK, nil
}

// logic to get a user
func (us *UserService) GetUserByID(id string) (models.Users, error) {
    var user models.Users

    if err := us.db.Where("id =?", id).First(&user).Error; err!= nil {
        return user, err
    }

    return user, nil
}

// logic to delete many users
func (us *UserService) DeleteManyUsers(userIDs []string) (int64, int, error) {
    result := us.db.Where("id IN ?", userIDs).Delete(&models.Users{})
    
    if result.Error != nil {
        return 0, http.StatusInternalServerError, result.Error
    }

    return int64(result.RowsAffected), http.StatusOK, nil
}

// logic to login a user
type LoginDataType struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (us *UserService) Login(loginData LoginDataType) (int, string, string, error) {
    var user models.Users
    
    // Find user in db by email and check if it exists
    if err := us.db.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
        return http.StatusBadRequest, "", "", err
    }

    // Compare password from loginData with hashed password in the database
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err!= nil {
        return http.StatusBadRequest, "", "", err
    }

    // // If login is successful, create an access token and return it along with user ID and email
    accessToken, err := GenerateAccessToken(PayloadType{
        ID:   user.ID,
        Role: user.Role,
    }, global.Config.Security.AccessKey) // Use LoadConfig() to get the AccessKey
    if err != nil {
        return http.StatusInternalServerError, "", "", err
    }

    refreshToken, err := GenerateRefreshToken(PayloadType{
        ID:   user.ID,
        Role: user.Role,
    }, global.Config.Security.RefreshKey)
    if err != nil {
    // Handle the error, e.g., log it or return it
        return http.StatusInternalServerError, "", "", err
    }

    return http.StatusOK, accessToken, refreshToken, nil
}

// logic to refresh a user's access token
func (us *UserService) RefreshAccessToken(refreshToken string) (int, string, error) {
    // Define the JWT claims struct to hold the payload
    var payload PayloadType

    // Extract the payload from the refresh token
    token, err := jwt.ParseWithClaims(refreshToken, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(global.Config.Security.RefreshKey), nil
    })

    if err != nil {
        return http.StatusUnauthorized, "", err
    }

    // Validate the token and extract the claims
    if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
        // Extract the payload from the claims
        payload.ID = (*claims)["id"].(string)
        payload.Role = (*claims)["role"].(string)
        accessToken, err := GenerateAccessToken(payload, global.Config.Security.AccessKey)
        
        if err != nil {
            return http.StatusInternalServerError, "", err
        }

        return http.StatusOK, accessToken, nil
    } else {
        return http.StatusInternalServerError, "", nil
    }
}
