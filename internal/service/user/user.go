package user

import (
	"my_shop/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateUser thêm một User mới vào cơ sở dữ liệu
func CreateUser(c *gin.Context) {
	var user models.Users
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Thời gian hiện tại
		currentTime := time.Now()

		// Query thêm User vào cơ sở dữ liệu
		query := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
		_, err := service.DB().Exec(query, user.Username, user.Email, user.Password, currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}