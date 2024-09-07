package service

import (
	"API/internal/repository"
	"API/internal/service/user"
)

type Service struct {
	User user.IUserService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		User: user.NewUserService(&r.User),
	}
}
