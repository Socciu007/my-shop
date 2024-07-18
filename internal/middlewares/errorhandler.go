package middlewares

import (
	"my_shop/internal/utils"

	"github.com/gin-gonic/gin"
)

// Handles errors by capturing any errors that occur during request processing and responding with a standardized error message.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
        c.Next()

        // Only handle errors if there are any
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			statusCode := c.Writer.Status()
			message := "An error occurred"
			
			utils.RespondStanders(c, statusCode, message, err.Error(), nil)
		}
		// Abort to ensure no further handlers are called
		c.Abort()
    }
}