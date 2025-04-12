package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"living-library/internal/models"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=living_library port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Book{}, &models.User{}) // добавь &models.User{}
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	DB = db
	fmt.Println("✅ Connected to database and migrated")
}
