package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"my_shop/internal/controllers"
	"my_shop/internal/middlewares"
	"my_shop/internal/services"
)

func UserRouter(r *gin.Engine, db *gorm.DB) {

	userService := services.NewUserService(db)
	userController := controllers.NewUserController(&userService)

	userGroup := r.Group("/api")
	{
		userGroup.GET("/get-users", userController.GetAllUsers)
		userGroup.GET("/get-user/:id", userController.GetUserByID)
		userGroup.POST("/create-user", middlewares.ValidationUser(), userController.CreateUser)
		userGroup.POST("/delete-users", userController.DeleteManyUsers)
		userGroup.POST("/login", middlewares.ValidationCredentials(), userController.Login)
		userGroup.POST("/logout", userController.Logout)
		userGroup.POST("/refresh-token", userController.RefreshAccessToken)
		userGroup.PATCH("/update-user/:id", userController.UpdateUser)
		userGroup.DELETE("/delete-user/:id", userController.DeleteUser)
	}
}