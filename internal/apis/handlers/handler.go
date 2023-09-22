package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapibackend/internal/application/services"
	"goapibackend/internal/domain/dto"
	"net/http"
	"strconv"
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
	_, err = hn.UserService.Signup(&userDto)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User Saved SuceessFully"),
	})
}

func (hn Handler) AddProject(c *gin.Context) {
	var userProjectDto dto.UserProjectDto
	err := c.BindJSON(&userProjectDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error Parsing the Data %t", err),
		})
		return
	}
	_, err = hn.UserService.AddProject(&userProjectDto)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Project Saved SuceessFully"),
	})
}

func (hn Handler) GetProject(c *gin.Context) {

	// We will get the query params
	page := c.Query("page")
	limit := c.Query("limit")

	pageInt, err2 := strconv.Atoi(page)
	if err2 != nil {
		pageInt = 1
	}
	limitInt, err2 := strconv.Atoi(limit)
	if err2 != nil {
		limitInt = 10
	}

	data, err := hn.UserService.GetAllProjects(pageInt, limitInt)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("SuceessFully Retrieved Projects"),
		"data":    data,
	})
}

func (hn Handler) SignIn(c *gin.Context) {
	var userDto dto.SignInDto
	err := c.BindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Error Parsing the Data %t", err),
		})
		return
	}
	_, err = hn.UserService.SignIn(&userDto)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("User LoggedIn SuceessFully"),
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
