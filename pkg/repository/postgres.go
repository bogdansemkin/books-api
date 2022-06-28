package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	usersTable = "users"
	booksTable = "books"
)

type Config struct {
	Host 	 string
	Port 	 string
	Username string
	Password string
	DBname 	 string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error){
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil{
		logrus.Errorf("Error during connection to database : %s", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil{
		logrus.Printf("Port is: %s, %s", cfg.Host, cfg.Port )
		logrus.Errorf("Error during binding to database : %s", err)
		return nil, err
	}
	return db, nil
}
