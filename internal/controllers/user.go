package controllers

import (
	"fmt"
	"my_shop/internal/models"
	"my_shop/internal/services"
	"my_shop/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
    return &UserController{userService: userService}
}

// handle request GET /users to get a list of all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.userService.GetAllUsers()
    if err != nil {
        utils.RespondStanders(c, http.StatusInternalServerError, "Internal Server Error", err.Error(), nil)
        return
    }

    utils.RespondStanders(c, http.StatusOK, "Get information of all users successfully", "", users)
}

//handle request POST /user to create a new user
func (uc *UserController) CreateUser(c *gin.Context) {
    validatedUser, exists := c.Get("validatedUser")
    if !exists {
        utils.RespondStanders(c, http.StatusNotFound, "User not found in context", "", nil)
        return
    }
    
    user := validatedUser.(models.Users)
    if status, err := uc.userService.CreateUser(&user); err != nil {
        utils.RespondStanders(c, status, "User creation failed", err.Error(), nil)
        return
    }

    utils.RespondStanders(c, http.StatusOK, "User created successfully", "", nil)
}

// handle request UPDATE /user to update a user
func (uc *UserController) UpdateUser(c *gin.Context) {
    var updatedUser models.Users
    
    // Get user ID from path parameter
    id := c.Param("id")
    if id == "" {
        utils.RespondStanders(c, http.StatusBadRequest, "ID is required", "ID is empty", nil)
        return
    }

    // Populate updatedUser with the data from the request body
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        utils.RespondStanders(c, http.StatusBadRequest, "Error format user information to edit", err.Error(), nil)
        return
    }

    // Call UserService to update the user
    updated, status, err := uc.userService.UpdateUser(id, updatedUser); 
    
    if err != nil {
        utils.RespondStanders(c, status, "User editing failed", err.Error(), nil)
        return
    }

    utils.RespondStanders(c, status, "User updated successfully", "", updated)
}

// handle request DELETE /user to delete a user
func (uc *UserController) DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        c.JSON(http.StatusBadRequest, gin.H{"status": 0, "error": "ID is required"})
        return
    }

    status, err := uc.userService.DeleteUser(id)
    if err != nil {
        utils.RespondStanders(c, status, "User deletion failed", err.Error(), nil)
        return
    }

    utils.RespondStanders(c, status, "User deleted successfully", "", nil)
}

// handle request GET /user/{id} to get a single user by ID
func (uc *UserController) GetUserByID(c *gin.Context) {
    id := c.Param("id")
    if id == "" {
        utils.RespondStanders(c, http.StatusBadRequest, "ID is required", "ID is empty", nil)
        return
    }

    user, err := uc.userService.GetUserByID(id)
    if err != nil {
        utils.RespondStanders(c, http.StatusInternalServerError, "Internal Server Error", err.Error(), nil)
        return
    }

    utils.RespondStanders(c, http.StatusOK, "Get user information successfully", "", user)
}

// handle request POST /delete-users to delete many users
func (uc *UserController) DeleteManyUsers(c *gin.Context) {
    var userIDs []string

    // Bind the request body to the userIDs slice
    if err := c.ShouldBindJSON(&userIDs); err != nil {
        utils.RespondStanders(c, http.StatusBadRequest, "Error binding user IDs", err.Error(), nil)
        return
    }

    // Check if at least one user ID is provided
    if len(userIDs) < 1 {
        utils.RespondStanders(c, http.StatusBadRequest, "At least one user ID is required", "UserIDs is empty", nil)
        return
    }

    // Call UserService to delete many users
    deletedCount, status, err := uc.userService.DeleteManyUsers(userIDs)
    if err!= nil {
        utils.RespondStanders(c, status, "Internal Server Error", err.Error(), nil)
        return
    }

    utils.RespondStanders(c, status, fmt.Sprintf("Deleted %d user(s) successfully", deletedCount), "", nil)
}

// handle request POST /login to login system
func (uc *UserController) Login(c *gin.Context) {
    loginData, exist := c.Get("loginData")
    if !exist {
        utils.RespondStanders(c, http.StatusNotFound, "Log in form is required", "Login form not found in context", nil)
        return
    }

    status, _, refreshToken, err := uc.userService.Login(loginData.(services.LoginDataType))
    if err!= nil {
        utils.RespondStanders(c, status, "Login failed", err.Error(), nil)
        return
    }

    // Set the refresh token cookie
    cookie := &http.Cookie{
        Name: "refresh_token",
        Value: refreshToken,
        Path: "/",
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
    }
    http.SetCookie(c.Writer, cookie)

    utils.RespondStanders(c, status, "Login successful", "", nil)
}

// handle request POST /logout to logout system
func (uc *UserController) Logout(c *gin.Context) {
    // Clear the refresh token cookie
    cookie := &http.Cookie{
        Name:   "refresh_token",
        Value:  "",
        Path:   "/",
        MaxAge: -1, // Set MaxAge to -1 to expire the cookie immediately
        HttpOnly: true,
        SameSite: http.SameSiteStrictMode,
    }
    http.SetCookie(c.Writer, cookie)

    // Respond with a success message
    utils.RespondStanders(c, http.StatusOK, "Logged out successfully", "", nil)
}

// handle request POST /refresh-token to refresh the access token
func (uc *UserController) RefreshAccessToken(c *gin.Context) {
    token, err := c.Request.Cookie("refresh_token")
    if err != nil || token == nil || token.Value == "" {
        utils.RespondStanders(c, http.StatusNotFound, "Refresh token is required", "Refresh token not found", nil)
        return
    }
}