package routers

import (
	"github.com/gin-gonic/gin"

	"my_shop/internal/controller/user"
)

func UserRouter(r *gin.Engine) {
	userGroup := r.Group("/user")
	{
		userGroup.GET("/get-users", user.GetAllUsers)
	}
}