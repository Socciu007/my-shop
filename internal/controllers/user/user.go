package user

import (
	"my_shop/internal/models"
	"my_shop/internal/services/user"
	"my_shop/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    userService *user.UserService
}

func NewUserController(userService *user.UserService) *UserController {
    return &UserController{userService: userService}
}

// handle request GET /users to get a list of all users
func (uc *UserController) GetAllUsers(c *gin.Context) {
    users, err := uc.userService.GetAllUsers()
    if err != nil {
        utils.RespondForHTTP(c, "InternalServerError", "Internal Server Error", err, nil)
        return
    }

    utils.RespondForHTTP(c, "OK", "Get information of all users successfully", nil, users)
}

//handle request POST /user to create a new user
func (uc *UserController) CreateUser(c *gin.Context) {
    validatedUser, exists := c.Get("validatedUser")
    if !exists {
        utils.RespondForHTTP(c, "NotFound", "User not found in context", nil, nil)
        return
    }
    
    user := validatedUser.(models.Users)
    if status, err := uc.userService.CreateUser(&user); err != nil {
        utils.RespondForHTTP(c, status, "", err, nil)
        return
    }

    utils.RespondForHTTP(c, "OK", "User created successfully", nil, nil)
}

// handle request UPDATE /user to update a user
func (uc *UserController) UpdateUser(c *gin.Context) {
    var updatedUser models.Users
    
    // Get user ID from path parameter
    id := c.Param("id")
    if id == "" {
        utils.RespondForHTTP(c, "BadRequest", "ID is required", nil, nil)
        return
    }

    // Populate updatedUser with the data from the request body
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        utils.RespondForHTTP(c, "BadRequest", "Bad request body", err, nil)
        return
    }

    // Call UserService to update the user
    updated, status, err := uc.userService.UpdateUser(id, updatedUser); 
    
    if err != nil {
        utils.RespondForHTTP(c, status, "", nil, nil)
        c.JSON(http.StatusInternalServerError, gin.H{"status": 0, "error": err.Error()})
        return
    }

    if !updated {
        c.JSON(http.StatusNotFound, gin.H{"status": 0, "error": "User not found"})
        return
    }

    utils.RespondForHTTP(c, status, "User updated successfully", nil, nil)
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
        utils.RespondForHTTP(c, status, "User is not exist", err, nil)
        return
    }

    utils.RespondForHTTP(c, status, "User deleted successfully", nil, nil)
}