package repository

import (
	"books-api/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres{
	return &BookPostgres{db: db}
}

func (r *BookPostgres) GetAllBooks() ([]model.Book, error){
	var book []model.Book

	query := fmt.Sprintf("SELECT * FROM %s", booksTable)
	_ = r.db.Select(&book, query)
	log.Printf("Book on getAllBooks method, %s", book)
	return book, nil
}

func (r *BookPostgres) GetBookById(id int) (model.Book, error){
	return model.Book{}, nil
}

func (r *BookPostgres) CreateBook(book model.Book)  (int, error){
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, author) VALUES($1, $2) RETURNING id ", booksTable)
	row := r.db.QueryRow(query, book.Name, book.Author)

	err := row.Scan(&id)
	if err != nil{
		log.Fatalf("error during scanning id on creating book, %s", err)
		return 0, err
	}

	return id, nil
}

func (r *BookPostgres) UpdateBook(book model.Book) error{
	return nil
}

func (r *BookPostgres) DeleteBook(id int) error{
	return nil
}
