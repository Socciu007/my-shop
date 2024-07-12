package user

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"my_shop/internal/models"
)

func CreateUser(c *gin.Context) (gin.H) {
    var user models.Users
    if err := c.ShouldBindJSON(&user); err != nil {
        return gin.H{"error": err.Error()}
    }

    // Get current time
    currentTime := time.Now()

    // Insert user into the database
    db := config.db()
	stmt, err := db.Prepare("INSERT INTO users(username, email, password, createdAt, updatedAt) VALUES(?, ?, ?, ?, ?)")
    if err != nil {
        log.Fatal(err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database prepare statement error"})
        return
    }

    
}