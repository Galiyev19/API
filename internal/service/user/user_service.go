package user

import (
	"time"

	"API/internal/api/helpers"
	"API/internal/model"
	"API/internal/repository/user"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
)

const (
	secretKey = "qweqje345345lkkadlaks9"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

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
	GetUserByEmail(email string) (model.User, error)
	GenerateToken(email, password string) (string, error)
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

// FIND USER BY EMAIL
func (u *UserService) GetUserByEmail(email string) (model.User, error) {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	return user, err
}

// GENERATE TOKEN
func (u *UserService) GenerateToken(email, password string) (string, error) {
	user, err := u.GetUserByEmail(email)
	if err != nil {
		return "NOT FOUND", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "WRONG PASSWORD", err
	}

	// Create the token claims
	claims := &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.UserId,
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate the signed token string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
