package user

import (
	// "my_shop/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get user information
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	//service
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "name": "Tien"})
}

// CreateUser handles the creation of a new user
func CreateUser(c *gin.Context) {
    // statusCode, res := user.CreateUser(c)
    // c.JSON(statusCode, res)
}