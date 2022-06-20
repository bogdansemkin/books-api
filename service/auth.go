package service

import (
	"books-api/model"
	"books-api/repository"
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