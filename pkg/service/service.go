package service

import (
	"API/pkg/models"
	"API/pkg/repository"
)

type Authorization interface {
	CreateAdmin(admin models.Admin) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int64, error)
}

type Users interface{}

type Service struct {
	Authorization
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
