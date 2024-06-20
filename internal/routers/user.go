package routers

import (
	"my_shop/internal/controller/user"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/:id", user.GetUser)
		userGroup.POST("/create", user.CreateUser)
	}
}