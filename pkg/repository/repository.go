package repository

import (
	"database/sql"

	"API/pkg/models"
)

type Authorization interface {
	CreateAdmin(admin models.Admin) (int, error)
	GetAdmin(email string) (*models.Admin, error)
	InsertUser(user models.User) (int, error)
}

type Users interface {
	GetUserList() (*[]models.User, error)
	GetUserByID(ID int) (*models.User, error)
	UpdateUser(user models.User, ID int) error
}

type Repository struct {
	Authorization
	Users
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Users:         NewUsersPostgres(db),
	}
}
