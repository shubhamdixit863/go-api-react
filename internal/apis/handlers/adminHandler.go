package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"

	"goapibackend/internal/application/services"
)

type AdminHandler struct {
	// We will be putting the services
	AdminService services.IAdminService
}

func (ad *AdminHandler) GetUsers(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	log.Println(page, limit)
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.Error(err)

	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.Error(err)

	}
	users, err := ad.AdminService.GetUser(pageInt, limitInt)
	if err != nil {
		c.Error(err)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    users,
	})

}
