package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"go-crud-api/config"
	"go-crud-api/helper"
	"go-crud-api/models"
)

func RefreshToken(c *gin.Context) {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek token di DB
	var rt models.RefreshToken

	if err := config.DB.Where("token = ?", input.RefreshToken).First(&rt).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token tidak valid"})

		return

	}

	token, err := jwt.Parse(input.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY_RAKHA"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh token expired"})

		return
	}

	claims := token.Claims.(jwt.MapClaims)

	if claims["type"] != "refresh" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token bukan refresh token"})
	}
	userID := uint(claims["user_id"].(float64))

	newAccessToken, _ := helper.GenerateAccessToken(userID)
	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
