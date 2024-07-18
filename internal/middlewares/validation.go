package middlewares

import (
	"my_shop/internal/models"
	"my_shop/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// validates user data and logs it
func ValidationUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.Users
        validate = validator.New()

        if err := c.ShouldBindJSON(&user); err != nil {
            utils.RespondStanders(c, http.StatusBadRequest, "Invalid request payload", err.Error(), nil)
            c.Abort()
            return
        }

        if err := validate.Struct(user); err != nil {
            utils.RespondStanders(c, http.StatusBadRequest, "Error validate user", err.Error(), nil)
            c.Abort()
            return
        }

        c.Set("validatedUser", user)
        c.Next()
    }
}
