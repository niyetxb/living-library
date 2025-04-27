package routes

import (
	"github.com/gin-gonic/gin"
	"living-library/internal/delivery"
	"living-library/internal/middleware"
)

func RegisterBookRoutes(r *gin.Engine, delivery *delivery.BookDelivery) {
	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("/", delivery.GetBooks)
		bookRoutes.GET("/:id", delivery.GetBookByID)

		bookRoutes.POST("/", middleware.AuthMiddleware(), middleware.AdminOnly(), delivery.CreateBook)
		bookRoutes.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), delivery.UpdateBook)
		bookRoutes.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), delivery.DeleteBook)
	}
}
