package routes

import (
	"go-crud-api/controllers"
	"go-crud-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/profile", middleware.AuthMiddleware(), controllers.GetProfile)
		api.POST("/users", controllers.CreateUsers)
		api.GET("/users", middleware.AuthMiddleware(), controllers.GetUsers)
		api.GET("/users/:id", middleware.AuthMiddleware(), controllers.GetUserByID)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
		api.POST("/refresh", controllers.RefreshToken)
		api.POST("/logout", controllers.Logout)
	}
}
