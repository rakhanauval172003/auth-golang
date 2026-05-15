package controllers

import (
	"net/http"

	"go-crud-api/config"
	"go-crud-api/models"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {

	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Where("token = ?", input.RefreshToken).Delete(&models.RefreshToken{})

	c.JSON(http.StatusOK, gin.H{"message": "Logout Berhasil"})

}
