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
		c.JSON(http.StatusOK, gin.H{
			"message": "Error Parsing the Data",
		})
		return
	}
	fmt.Println(userDto)
	hn.UserService.Signup()
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
