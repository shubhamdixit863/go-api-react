package handlers

import (
	"github.com/gin-gonic/gin"
	"goapibackend/internal/application/services"
	"net/http"
)

// This will be for handlers

type Handler struct {
	// This will hold the service dependency
	UserService services.IUserService
}

func (hn Handler) SignUp(c *gin.Context) {
	hn.UserService.Signup()
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
