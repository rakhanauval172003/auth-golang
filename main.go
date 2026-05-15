package main

import (
	"go-crud-api/config"
	"go-crud-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()
	routes.SetupRoutes(r)

	r.Run(":8000")
}
