package repository

import (
	"database/sql"

	"API/internal/repository/user"
)

type Repository struct {
	User user.UserRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User: *user.NewUserRepo(db),
	}
}
