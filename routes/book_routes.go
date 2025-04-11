package routes

import (
	"bookapi/controllers"
	"bookapi/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(r *gin.Engine) {
	book := r.Group("/books")
	book.Use(middleware.JWTMiddleware())
	{
		book.GET("", controllers.GetBooks)
		book.POST("", controllers.CreateBook)
	}
}
