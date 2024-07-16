package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	uc "my_shop/internal/controllers/user"
	"my_shop/internal/middlewares"
	us "my_shop/internal/services/user"
)

func UserRouter(r *gin.Engine, db *gorm.DB) {

	userService := us.NewUserService(db)
	userController := uc.NewUserController(&userService)

	userGroup := r.Group("/api")
	{
		userGroup.GET("/get-users", userController.GetAllUsers)
		userGroup.POST("/create-user", middlewares.ValidationUser(), userController.CreateUser)
	}
}