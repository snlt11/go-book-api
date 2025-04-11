package routes

import (
	"bookapi/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
}
