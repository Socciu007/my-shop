package routers

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Open (or create) the debug.log file for logging
	f, err := os.OpenFile("internal/log/debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	// Create new logger use debug.log file
	logger := log.New(f, "", log.LstdFlags)

	// create new instance without default middleware
	router := gin.New()

	// Add necessary middleware
	router.Use(gin.LoggerWithWriter(f)) // Records information about HTTP requests: http method, URL path, response status code, proccessing time
	router.Use(gin.RecoveryWithWriter(f)) // Helps applications recover after errors
	router.Use(cors.Default()) // Allow or block requests from different sources based on configuration. Default: all origins, all http methods, all headers
	router.Use(gzip.Gzip(gzip.DefaultCompression)) // Compress HTTP responses
	router.Use(func(c *gin.Context) {// Middleware to set the X-Frame-Options header
		c.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")
		c.Next()
	})
	router.Use(func(c *gin.Context) {// Middleware to set the X-XSS-Protection header
		c.Writer.Header().Set("X-XSS-Protection", "1; mode=block")
		c.Next()
	})
	router.Use(func(c *gin.Context) {// Custom Middleware for more detailed error logging
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Printf("| %3d | %13v | %15s | %-7s  %#v\n",
			status,
			latency,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
		)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				logger.Printf("ERROR: %s\n", e)
			}
		}
	})

	// Register sub router
	PingRouter(router)
	UserRouter(router, db)
	V1Router(router)

	return router
}