package routes

import (
	"github.com/gin-gonic/gin"
	"living-library/internal/delivery"
	"living-library/internal/middleware"
)

func SetupRoutes(r *gin.Engine, delivery *delivery.BookDelivery, authHandler *delivery.AuthHandler) {

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	bookRoutes := r.Group("/books", middleware.AuthMiddleware())
	{
		bookRoutes.GET("/", delivery.GetBooks)
		bookRoutes.GET("/:id", delivery.GetBookByID)
		bookRoutes.POST("/", middleware.AdminOnly(), delivery.CreateBook)
		bookRoutes.PUT("/:id", middleware.AdminOnly(), delivery.UpdateBook)
		bookRoutes.DELETE("/:id", middleware.AdminOnly(), delivery.DeleteBook)
	}
}
