package service

import (
	"books-api/pkg/model"
	"books-api/pkg/repository"
)

type AuthService struct {
	repos repository.Authorization
}

func NewAuthService(repos repository.Authorization) *AuthService{
	return &AuthService{repos: repos}
}

func (s *AuthService) Create(user model.User) (int, error){
	return s.repos.Create(user)
}

func (s *AuthService) Get(name, password string) (model.User, error){
	return s.repos.Get(name, password)
}