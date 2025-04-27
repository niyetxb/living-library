package routes

import (
	"github.com/gin-gonic/gin"
	"living-library/internal/delivery"
	"living-library/internal/middleware"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *delivery.AuthHandler) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.PUT("/update-profile", middleware.AuthMiddleware(), authHandler.UpdateProfile)
	}
}
