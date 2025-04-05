package repository

import (
	"living-library/internal/database"
	"living-library/internal/models"
)

func GetBooks() ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Find(&books)
	return books, result.Error
}

func GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func CreateBook(book models.Book) (models.Book, error) {
	result := database.DB.Create(&book)
	return book, result.Error
}

func UpdateBook(id uint, updated models.Book) (*models.Book, error) {
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	book.Title = updated.Title
	book.Author = updated.Author
	book.Year = updated.Year
	database.DB.Save(&book)
	return &book, nil
}

func DeleteBook(id uint) error {
	result := database.DB.Delete(&models.Book{}, id)
	return result.Error
}
