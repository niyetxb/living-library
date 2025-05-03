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
	bookService := services.NewBookService(bookRepo)

	bookDelivery := delivery.NewBookDelivery(bookService)
	authHandler := delivery.NewAuthHandler()
	r := gin.Default()
	routes.SetupRoutes(r, bookDelivery, authHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Сервер не смог запуститься: ", err)
	}
}
