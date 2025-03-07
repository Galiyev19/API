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

func (s *UserService) GetUserByID(ID int) (*models.User, error) {
	return s.repo.GetUserByID(ID)
}

func (s *UserService) UpdateUser(user models.User, ID int) error {
	return s.repo.UpdateUser(user, ID)
}
