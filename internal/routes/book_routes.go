package routes

import (
	"github.com/gin-gonic/gin"
	"living-library/internal/delivery"
)

func RegisterBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("/", handlers.GetBooks)
		bookRoutes.GET("/:id", handlers.GetBookByID)
		bookRoutes.POST("/", handlers.CreateBook)
		bookRoutes.PUT("/:id", handlers.UpdateBook)
		bookRoutes.DELETE("/:id", handlers.DeleteBook)
	}
}
