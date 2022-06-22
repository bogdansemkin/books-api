package repository

import (
	"books-api/model"
	"github.com/jmoiron/sqlx"
)

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

type Repository struct {
	Authorization
	Book
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Book: 		   NewBookPostgres(db),
	}
}
