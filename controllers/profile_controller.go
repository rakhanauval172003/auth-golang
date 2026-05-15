package controllers

import (
	"net/http"

	"go-crud-api/config"
	"go-crud-api/models"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User tidak ditemukan"})

		return
	}

	var user models.User

	if err := config.DB.First(&user, userId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User Tidak ditemukan"})
		return
	}

	// jangan kirim password
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.NAME,
		"email": user.Email,
	})

}
