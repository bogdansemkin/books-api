package service

import (

	"books-api/pkg/model"
	"books-api/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	Create(user model.User) (int, error)
	Get(name, password string) (model.User, error)
}

type Book interface {
	GetAllBooks() ([]model.Book, error)
	GetBookById(id int) (model.Book, error)
	CreateBook(book model.Book) (int, error)
	UpdateBook(book model.Book) error
	DeleteBook(id int) error
}

type Service struct{
	Authorization
	Book
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Book: NewBookService(repos.Book),
	}
}
