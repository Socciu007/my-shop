package user

import (
	"my_shop/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllUsersHandler xử lý request GET /api/users để lấy danh sách tất cả người dùng
func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": "err.Error()"})
    users, err := service.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": 1, "message": "Get information of all users successfully", "data": users})
}