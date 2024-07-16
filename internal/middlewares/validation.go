package middlewares

import (
	"my_shop/internal/models"
	"my_shop/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// ValidationUser validates user data and logs it
func ValidationUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.Users
        validate = validator.New()

        if err := c.ShouldBindJSON(&user); err != nil {
            utils.RespondForHTTP(c, "BadRequest", "Invalid request payload", err, nil)
            c.Abort()
            return
        }

        if err := validate.Struct(user); err != nil {
            utils.RespondForHTTP(c, "BadRequest", "Error validate user", err, nil)
            c.Abort()
            return
        }

        c.Set("validatedUser", user)
        c.Next()
    }
}
