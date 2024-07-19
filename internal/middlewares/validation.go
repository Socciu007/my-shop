package middlewares

import (
	"my_shop/internal/models"
	"my_shop/internal/utils"
	"net/http"
	"regexp"

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

        // Check if ConfirmPassword is present in the request payload
        if _, exists := c.GetPostForm("ConfirmPassword"); !exists {
            utils.RespondStanders(c, http.StatusBadRequest, "ConfirmPassword is required", "ConfirmPassword is missing", nil)
            c.Abort()
            return
        }

        // Add condition check for confirmPassword with Password
        if user.Password != c.PostForm("ConfirmPassword") {
            utils.RespondStanders(c, http.StatusBadRequest, "Password and ConfirmPassword need to match", "Password and ConfirmPassword do not match", nil)
            c.Abort()
            return
        }

        // Validate user struct
        if err := validate.Struct(user); err != nil {
            utils.RespondStanders(c, http.StatusBadRequest, "Error validate user", err.Error(), nil)
            c.Abort()
            return
        }

        c.Set("validatedUser", user)
        c.Next()
    }
}

func ValidationCredentials() gin.HandlerFunc {
    return func(c *gin.Context) {
        var loginData struct {
            Email string `json:"email" binding:"required"`
            Password string `json:"password" binding:"required"`
        }
        
        if err := c.ShouldBindJSON(&loginData); err!= nil {
            utils.RespondStanders(c, http.StatusBadRequest, "The input is required", err.Error(), nil)
            c.Abort()
            return
        }

        // Validate email (alphanumeric, 3-30 characters)
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9]{3,30}$`)
		if !emailRegex.MatchString(loginData.Email) {
			utils.RespondStanders(c, http.StatusBadRequest, "Usermail must be email", "Invalid email format", nil)
            c.Abort()
			return
		}

		// Validate password (at least 8 characters)
		if len(loginData.Password) < 8 {
            utils.RespondStanders(c, http.StatusBadRequest, "Password must be at least 8 characters long", "Invalid password format", nil)
			c.Abort()
			return
		}

        c.Set("loginData", loginData)
        c.Next()
    }
}
