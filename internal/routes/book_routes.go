package routes

import (
	"alive-library/internal/delivery"
	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(router *gin.Engine) {
	books := router.Group("/books")
	{
		books.GET("/", delivery.GetBooks)
		books.GET("/:id", delivery.GetBook)
		books.POST("/", delivery.CreateBook)
		books.PUT("/:id", delivery.UpdateBook)
		books.DELETE("/:id", delivery.DeleteBook)
	}
}
