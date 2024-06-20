package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// create new instance without default middleware
	router := gin.New()

	// Add necessary middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Register sub router
	PingRouter(router)
	UserRouter(router)
	V1Router(router)

	return router
}