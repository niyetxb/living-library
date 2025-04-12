package repository

import (
	"living-library/internal/database"
	"living-library/internal/models"
)

type BookRepository interface {
	GetBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	CreateBook(book models.Book) (models.Book, error)
	UpdateBook(id uint, book models.Book) (*models.Book, error)
	DeleteBook(id uint) error
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) GetBooks() ([]models.Book, error) {
	var books []models.Book
	result := database.DB.Find(&books)
	return books, result.Error
}

func (r *bookRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	result := database.DB.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *bookRepository) CreateBook(book models.Book) (models.Book, error) {
	result := database.DB.Create(&book)
	return book, result.Error
}

func (r *bookRepository) UpdateBook(id uint, updatedBook models.Book) (*models.Book, error) {
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return nil, err
	}
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Year = updatedBook.Year
	database.DB.Save(&book)
	return &book, nil
}

func (r *bookRepository) DeleteBook(id uint) error {
	return database.DB.Delete(&models.Book{}, id).Error
}
