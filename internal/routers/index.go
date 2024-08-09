package routers

import (
	"my_shop/global"
	"my_shop/internal/middlewares"
	"my_shop/internal/routers/v1"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// create new instance without default middleware
	router := gin.New()

	// Add necessary middleware
	router.Use(gin.Recovery()) // Helps applications recover after errors
	router.Use(middlewares.CORS()) // Apply the CORS middleware to all routes
	router.Use(gzip.Gzip(gzip.DefaultCompression)) // Apply GZIP middleware to compress HTTP responses
	router.Use(middlewares.LoggerMiddleware(global.Logger.Logger))
	router.Use(middlewares.ErrorHandler())
	
	// Register sub router
	PingRouter(router)
	UserRouter(router)
	v1.V1Router(router)

	return router
}