package user

import (
	"API/internal/model"
	"API/internal/repository/user"
)

type UserService struct {
	UserRepo user.IUserRepo
}

func NewUserService(userRepo user.IUserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

type IUserService interface {
	Insert(user model.User) error
}

func (u *UserService) Insert(user model.User) error {
	userModel := model.User{
		UserId:   "1",
		Email:    user.Email,
		Password: user.Password,
	}

	err := u.UserRepo.Insert(userModel)
	if err != nil {
		return err
	}
	return nil
}
