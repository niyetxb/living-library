package routes

import (
	"github.com/gin-gonic/gin"
	"living-library/internal/delivery"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *delivery.AuthHandler) {
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.PUT("/update-profile", authHandler.UpdateProfile)
	}
}
