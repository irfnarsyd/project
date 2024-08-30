package controllers

import (
	"net/http"
	"project/app/infra/database"
	"project/app/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (ctrl *UserController) Register(c *gin.Context) {
	var requestRegister models.User

	if err := c.ShouldBindJSON(&requestRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ga ada payload"})
		return
	}

	saveUser := database.DB.Create(&requestRegister)

	if saveUser.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"meesage": &requestRegister})
}
