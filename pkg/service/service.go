package service

import (
	"API/pkg/models"
	"API/pkg/repository"
)

type Authorization interface {
	CreateAdmin(admin models.Admin) (int, error)
	CreateUser(user models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (string, error)
	GetAdmin(email string) (*models.Admin, error)
}

type Users interface {
	GetListUser() (*[]models.User, error)
}

type Products interface{}

type Service struct {
	Authorization
	Users
	Products
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Users:         NewUserService(repo.Users),
	}
}
