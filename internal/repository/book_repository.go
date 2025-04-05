package repository

import "alive-library/internal/models"

var books = []models.Book{
	{ID: 1, Title: "1984", Author: "George Orwell", Year: 1949},
}

func GetBooks() []models.Book {
	return books
}

func GetBookByID(id uint) *models.Book {
	for _, book := range books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func CreateBook(book models.Book) models.Book {
	book.ID = uint(len(books) + 1)
	books = append(books, book)
	return book
}

func UpdateBook(id uint, updatedBook models.Book) *models.Book {
	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			books[i].ID = id
			return &books[i]
		}
	}
	return nil
}

func DeleteBook(id uint) bool {
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return true
		}
	}
	return false
}
