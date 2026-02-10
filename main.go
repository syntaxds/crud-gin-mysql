package main

import (
	"crud-gin-mysql/config"
	"crud-gin-mysql/models"
	"crud-gin-mysql/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Product{})

	routes.SetupRoutes(r)

	r.Run(":8080")
}
