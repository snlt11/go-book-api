package main

import (
	"bookapi/config"
	"bookapi/models"
	"bookapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Book{}, &models.User{})

	routes.RegisterAuthRoutes(r)
	routes.RegisterBookRoutes(r)

	r.Run(":8080")
}
