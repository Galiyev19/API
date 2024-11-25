package service

import (
	"API/pkg/models"
	"API/pkg/repository"
)

type Authorization interface {
	CreateAdmin(admin models.Admin) (int, error)
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int64, error)
}

type Users interface{}

type Products interface{}

type Service struct {
	Authorization
	Users
	Products
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
