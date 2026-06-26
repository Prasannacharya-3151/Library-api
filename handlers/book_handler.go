package handlers

import (
	"library-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onsi/ginkgo/example/books"
)

func CreateBook(c *gin.Context) {
	var input models.CreatedBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	book, err := services.CreateBookService(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"book":book})
}

func GetAllBooks(c *gin.Context) {
	book, err := services.GetAllBooksServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"books":books})
}

