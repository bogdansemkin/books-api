package service

import (
	"books-api/model"
	"books-api/repository"
)

type Authorization interface {
	Create(user model.User) (int, error)
}

type Service struct{
	Authorization
}

func NewService(repos *repository.Repository) *Service{
	return &Service{Authorization: NewAuthService(repos)}
}
