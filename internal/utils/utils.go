package utils

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
    Status  int      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

// response function when interacting with a client via HTTP
func RespondStanders(c *gin.Context, statusCode int, message string, err string, data interface{}) {
    switch statusCode {
		case 200:
			if data == nil {
				c.JSON(statusCode, APIResponse{
					Status: 1,
					Message: message,
				})
			} else {
				c.JSON(statusCode, APIResponse{
					Status: 1,
					Message: message,
                    Data:    data,
				})
			}
		case 400:
			c.JSON(statusCode, APIResponse{
				Status: 0,
                Message: message,
                Error:   err,
			})
		case 401:
			c.JSON(statusCode, APIResponse{
				Status: 0,
                Message: message,
                Error:   err,
			})
		case 404:
			c.JSON(statusCode, APIResponse{
				Status: 0,
                Message: message,
                Error:   err,
			})
		case 409:
			c.JSON(statusCode, APIResponse{
				Status: 0,
                Message: message,
                Error:   err,
			})
		case 500:
			c.JSON(statusCode, APIResponse{
				Status: 0,
                Message: "Internal Server Error",
                Error:   err,
			})
		default:
			c.JSON(statusCode, gin.H{"status": 0, "message": "Internal Server Error"})
    }
}