package services

import (
	"living-library/internal/models"
	"living-library/internal/repository"
)

type BookService struct {
	BookRepo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{BookRepo: repo}
}

func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.BookRepo.GetBooks()
}

func (s *BookService) GetBookByID(id uint) (*models.Book, error) {
	return s.BookRepo.GetBookByID(id)
}

func (s *BookService) CreateBook(book models.Book) (models.Book, error) {
	return s.BookRepo.CreateBook(book)
}

func (s *BookService) UpdateBook(id uint, book models.Book) (*models.Book, error) {
	return s.BookRepo.UpdateBook(id, book)
}

func (s *BookService) DeleteBook(id uint) error {
	return s.BookRepo.DeleteBook(id)
}
