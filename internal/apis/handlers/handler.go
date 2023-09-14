package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapibackend/internal/application/services"
	"goapibackend/internal/domain/dto"
	"net/http"
)

// This will be for handlers

type Handler struct {
	// This will hold the service dependency
	UserService services.IUserService
}

func (hn Handler) SignUp(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error Parsing the Data %t", err),
		})
		return
	}
	userID, err := hn.UserService.Signup(&userDto)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User Saved SuceessFully With Id -%d", userID),
	})
}

func (hn Handler) GetAllUsers(c *gin.Context) {
	users, err := hn.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully retrieved the list of user",
		"data":    users,
	})
}
