package service

import (
	"API/pkg/models"
	"API/pkg/repository"
)

type UserService struct {
	repo repository.Users
}

func NewUserService(repo repository.Users) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetListUser() (*[]models.User, error) {
	return s.repo.GetUserList()
}
