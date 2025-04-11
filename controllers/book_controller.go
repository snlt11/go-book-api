package controllers

import (
	"bookapi/config"
	"bookapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	query := config.DB
	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}
	query.Offset((page - 1) * limit).Limit(limit).Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}
