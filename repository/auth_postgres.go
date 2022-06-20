package repository

import (
	"books-api/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Create(user model.User) (int, error){
	var id int

	return id, nil
}