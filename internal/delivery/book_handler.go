package delivery

import (
	"alive-library/internal/models"
	"alive-library/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBooks(c *gin.Context) {
	books := repository.GetBooks()
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book := repository.GetBookByID(uint(id))
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdBook := repository.CreateBook(newBook)
	c.JSON(http.StatusCreated, createdBook)
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book := repository.UpdateBook(uint(id), updatedBook)
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	success := repository.DeleteBook(uint(id))
	if !success {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
