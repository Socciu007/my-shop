package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// response function when interacting with a client via HTTP
func RespondForHTTP(c *gin.Context, status string, message string, err error, data interface{}) {
    switch status {
		case "OK":
			if data == nil {
				c.JSON(http.StatusOK, gin.H{"status": 1, "message": message})
			} else {
				c.JSON(http.StatusOK, gin.H{"status": 1, "message": message, "data": data})
			}
		case "BadRequest":
			c.JSON(http.StatusBadRequest, gin.H{"status": 0, "message": message, "error": err.Error()})
		case "Unauthorized":
			c.JSON(http.StatusUnauthorized, gin.H{"status": 0, "message": message, "error": err.Error()})
		case "NotFound":
			c.JSON(http.StatusNotFound, gin.H{"status": 0, "message": message, "error": err.Error()})
		case "Conflict":
			c.JSON(http.StatusConflict, gin.H{"status": 0, "message": message, "error": err.Error()})
		case "InternalServerError":
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": "Internal Server Error", "error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "message": "Internal Server Error"})
    }
}