package main

import (
	"github.com/gin-gonic/gin"
	"living-library/internal/database"
	"living-library/internal/routes"
)

func main() {
	database.InitDB()

	r := gin.Default()

	routes.RegisterBookRoutes(r)

	r.Run(":8080")
}
