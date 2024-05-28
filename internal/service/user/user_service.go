package user

import (
	"API/internal/api/helpers"
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

// Create new user
func (u *UserService) Insert(user model.User) error {
	hashPass, err := helpers.HashPassword(user.Password) // hashed password
	if err != nil {
		return err
	}
	id := helpers.GenerateId() // generate unique id

	userModel := model.User{
		UserId:   id.String(), // id
		Email:    user.Email,  // email
		Password: hashPass,    // password
	}

	err = u.UserRepo.Insert(userModel) // insert in db
	if err != nil {
		return err
	}
	return nil
}
