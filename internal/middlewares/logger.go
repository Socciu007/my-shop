package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// logs details of each incoming request and response
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
        // Start time
		startTime := time.Now()
		
		// Process request
        c.Next()

		// End time
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Get status code
		statusCode := c.Writer.Status()

		// Get client IP address
		clientIP := c.ClientIP()

		// Log format
		log.Printf("%s - [%s] \"%s %s %s\" %d %s \"%s\"",
			clientIP,
			startTime.Format(time.RFC1123),
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			statusCode,
			latency,
			c.Request.UserAgent(),
		)
    }
}