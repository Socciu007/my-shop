package routers

import (
	"my_shop/internal/controllers/ping"

	"github.com/gin-gonic/gin"
)

func PingRouter(r *gin.Engine) {
	pingRouter := r.Group("/ping")
	{
		pingRouter.GET("", ping.Ping)
	}
}