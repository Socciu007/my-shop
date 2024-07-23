package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// logs details of each incoming request and response
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
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

		// Log details about the request
		logger.Info("Logging request",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.Path),
			zap.String("client_ip", clientIP),
			zap.Int("status_code", statusCode),
            zap.Duration("duration", latency),
		)
    }
}