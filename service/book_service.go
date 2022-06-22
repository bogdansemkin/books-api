package service

import (
	"books-api/model"
	"books-api/repository"
)

type BookService struct {
	repos repository.Book
}

func NewBookService(repos repository.Book) *BookService {
	return &BookService{repos: repos}
}

func (s *BookService) GetAllBooks() ([]model.Book, error){
	return s.repos.GetAllBooks()
}

func (s *BookService) GetBookById(id int) (model.Book, error){
	return s.repos.GetBookById(id)
}

func (s *BookService) CreateBook(book model.Book) (int, error){
	return s.repos.CreateBook(book)
}

func (s *BookService) UpdateBook(book model.Book) error{
	return s.repos.UpdateBook(book)
}

func (s *BookService) DeleteBook(id int) error{
	return s.repos.DeleteBook(id)
}
