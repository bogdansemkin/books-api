package repository

import (
	"books-api/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres{
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Create(user model.User) (int, error){
	var id int
	query := fmt.Sprintf("INSERT INTO %s(name,username,password) VALUES($1,$2,$3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	err := row.Scan(&id)
	if err != nil {
		log.Fatalf("error during scanning id, %s", err)
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) Get(name, password string) (model.User, error){
	user := model.User{}
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1 AND password=$2", usersTable)

	err := r.db.Get(&user, query, name, password)
	if err != nil {
		log.Fatalf("Error during getting data, %s", err)
	}
	return user, nil
}