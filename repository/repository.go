package repository

import (
	"books-api/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(user model.User) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
