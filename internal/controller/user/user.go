package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get user information
func GetUser(c *gin.Context) {
	userID := c.Param("id")
	//service
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "name": "Tien"})
}

// create user
func CreateUser(c *gin.Context) {
	var user struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&user); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	//service
	c.JSON(http.StatusOK, gin.H{"message": "User created", "user": user})
}