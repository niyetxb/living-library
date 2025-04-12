package routes

import (
	"github.com/gin-gonic/gin"
	"living-library/internal/delivery"
)

func RegisterBookRoutes(r *gin.Engine, delivery *delivery.BookDelivery) {
	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("/", delivery.GetBooks)
		bookRoutes.GET("/:id", delivery.GetBookByID)
		bookRoutes.POST("/", delivery.CreateBook)
		bookRoutes.PUT("/:id", delivery.UpdateBook)
		bookRoutes.DELETE("/:id", delivery.DeleteBook)
	}
}
