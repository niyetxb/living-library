package main

import (
	"github.com/gin-gonic/gin"
	"living-library/internal/database"
	"living-library/internal/delivery"
	"living-library/internal/repository"
	"living-library/internal/routes"
	"living-library/internal/services"
	"log"
)

func main() {
	database.InitDB()
	bookRepo := repository.NewBookRepository()
	userRepo := repository.NewUserRepository()

	bookService := services.NewBookService(bookRepo)
	authService := services.NewAuthService(userRepo)

	bookDelivery := delivery.NewBookDelivery(bookService)
	authHandler := delivery.NewAuthHandler(authService)

	r := gin.Default()
	routes.RegisterBookRoutes(r, bookDelivery)
	routes.RegisterAuthRoutes(r, authHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Сервер не смог запуститься: ", err)
	}
}
