package controllers

import (
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