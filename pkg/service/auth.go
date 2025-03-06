package service

import (
	"errors"
	"fmt"
	"time"

	"API/pkg/models"
	"API/pkg/repository"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	signKey = "fhglkj2384kjklasdfHaSS"
)

type TokenClaims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) CreateAdmin(admin models.Admin) (int, error) {
	existAdmin, err := s.repo.GetAdmin(admin.Email)
	if err != nil {
		fmt.Println("Error", err)
		return 0, err
	}

	if existAdmin != nil {
		return 0, errors.New("admin this email already exist")
	}

	hashPassword, err := s.generateHashPassword(admin.Password)
	if err != nil {
		return 0, err
	}

	admin.Password = hashPassword
	admin.CreatedAt = time.Now()

	return s.repo.CreateAdmin(admin)
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	hashPassword, err := s.generateHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = hashPassword
	user.CreatedAt = time.Now()
	return s.repo.InsertUser(user)
}

func (s *AuthService) generateHashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("generate hashed password: %w", err)
	}

	return string(hashPassword), nil
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	admin, err := s.repo.GetAdmin(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		admin.ID,
	})

	return token.SignedString([]byte(signKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign method")
		}
		return []byte(signKey), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.ID, nil
}

func (s *AuthService) GetAdmin(email string) (*models.Admin, error) {
	return s.repo.GetAdmin(email)
}
