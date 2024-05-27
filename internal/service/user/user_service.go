package user

import "API/internal/repository/user"

type UserService struct {
	UserRepo user.IUserRepo
}

func NewUserService(userRepo user.IUserRepo) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

type IUserService interface {
	Insert() error
}

func (u *UserService) Insert() error {
	return nil
}
