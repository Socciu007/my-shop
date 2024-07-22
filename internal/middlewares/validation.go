package middlewares

import (
	"my_shop/internal/models"
	"my_shop/internal/utils"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validate *validator.Validate

// validates user data and logs it
func ValidationUser() gin.HandlerFunc {
    return func(c *gin.Context) {
        var user models.Users
        var confirmPassword string

        validate := validator.New()

        // Check if the request body is empty
        if c.Request.Body == nil {
            utils.RespondStanders(c, http.StatusBadRequest, "Request body is required", "Request body is empty", nil)
            c.Abort()
            return
        }

        // Bind and validate user and ConfirmPassword fields
        var requestBody struct {
            models.Users
            ConfirmPassword string `json:"confirmPassword"`
        }

        if err := c.ShouldBindJSON(&requestBody); err != nil {
            utils.RespondStanders(c, http.StatusBadRequest, "Invalid request payload", err.Error(), nil)
            c.Abort()
            return
        }

        user = requestBody.Users
        confirmPassword = requestBody.ConfirmPassword

        // Check if Password matches ConfirmPassword
        if user.Password != confirmPassword {
            utils.RespondStanders(c, http.StatusBadRequest, "Passwords do not match", "Password and ConfirmPassword do not match", nil)
            c.Abort()
            return
        }

        // Validate user struct
        if err := validate.Struct(&user); err != nil {
            utils.RespondStanders(c, http.StatusBadRequest, "Error validate user", err.Error(), nil)
            c.Abort()
            return
        }

        if user.Username == "" {
            user.Username = user.Email[:strings.Index(user.Email, "@")]
        }

        user.ID = uuid.New().String()

        c.Set("validatedUser", user)
        c.Next()
    }
}

type LoginDataType struct {
    Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}
func ValidationCredentials() gin.HandlerFunc {
    return func(c *gin.Context) {
        var loginData LoginDataType
        
        if err := c.ShouldBindJSON(&loginData); err!= nil {
            utils.RespondStanders(c, http.StatusBadRequest, "The input is required", err.Error(), nil)
            c.Abort()
            return
        }

        
        // Validate email (alphanumeric, 3-30 characters)
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
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
